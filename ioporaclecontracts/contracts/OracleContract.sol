// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "./DistKeyContract.sol";
import "./RaffleContract.sol";
import "cdlbot-solidity/contracts/crypto/BN256G1.sol";
import "cdlbot-solidity/contracts/crypto/BN256G2.sol";

contract OracleContract {
    struct ValidationResult {
        uint256 id;
        bool result;
    }

    uint256 private constant AGGREGATOR_REWARD = 0.001 ether;
    uint256 private constant VALIDATOR_REWARD = 0.0001 ether;
    uint256 private constant TOTAL_REWARD =
        AGGREGATOR_REWARD + VALIDATOR_REWARD;

    uint256 private requestCounter;

    mapping(uint256 => ValidationResult) private validationResults;

    event ValidateTransactionLog(
        address indexed sender,
        uint256 indexed id,
        bytes32 hash,
        uint256 confirmations
    );
    event SubmitValidationResultLog(
        address indexed sender,
        uint256 id,
        bool result
    );

    RegistryContract private registryContract;
    DistKeyContract private distKeyContract;
    address payable private raffleContract;

    constructor(
        address _registryContract,
        address _distKeyContract,
        address payable _raffleContract
    ) public {
        registryContract = RegistryContract(_registryContract);
        distKeyContract = DistKeyContract(_distKeyContract);
        raffleContract = _raffleContract;
    }

    function validateTransaction(bytes32 _hash, uint256 _confirmations)
        external
        payable
    {
        require(msg.value >= TOTAL_REWARD, "msg value below total reward");
        emit ValidateTransactionLog(
            msg.sender,
            ++requestCounter,
            _hash,
            _confirmations
        );
    }

    function submitValidationResult(
        uint256 _id,
        bool _result,
        uint256[2] calldata _signature
    ) external {
        require(!validationResultExists(_id), "already exists");
        require(
            registryContract.getAggregator().addr == msg.sender,
            "not the aggregator"
        );

        uint256[2] memory hash =
            BN256G1.hashToPointSha256(abi.encode(_id, _result));
        uint256[4] memory negG2 = BN256G2.getNegG2();
        uint256[4] memory publicKey = distKeyContract.getPublicKey();
        uint256[12] memory input =
            [
                hash[0],
                hash[1],
                publicKey[0],
                publicKey[1],
                publicKey[2],
                publicKey[3],
                _signature[0],
                _signature[1],
                negG2[0],
                negG2[1],
                negG2[2],
                negG2[3]
            ];
        require(BN256G1.bn256CheckPairing(input), "invalid signature");

        payable(msg.sender).transfer(AGGREGATOR_REWARD);
        raffleContract.transfer(VALIDATOR_REWARD);

        ValidationResult storage validationResult = validationResults[_id];
        validationResult.id = _id;
        validationResult.result = _result;

        emit SubmitValidationResultLog(msg.sender, _id, _result);
    }

    function findValidationResult(uint256 _id)
        public
        view
        returns (ValidationResult memory)
    {
        require(validationResultExists(_id), "not found");
        return validationResults[_id];
    }

    function validationResultExists(uint256 _id) public view returns (bool) {
        return validationResults[_id].id != 0;
    }
}
