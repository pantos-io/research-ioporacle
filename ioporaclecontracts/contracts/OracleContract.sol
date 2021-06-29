// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "./DistKeyContract.sol";
import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";

contract OracleContract {
    struct ValidationResult {
        uint256 id;
        bool result;
    }

    uint256 public constant BASE_FEE = 0.001 ether;
    uint256 public constant VALIDATOR_FEE = 0.0001 ether;
    uint256 public constant TOTAL_FEE = BASE_FEE + VALIDATOR_FEE;

    uint256 private constant G2_NEG_X_RE =
        0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 private constant G2_NEG_X_IM =
        0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 private constant G2_NEG_Y_RE =
        0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 private constant G2_NEG_Y_IM =
        0x1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d;

    uint256 private requestCounter;
    uint256 private requestsSinceLastPayout;

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
        bool result,
        uint256 fee
    );

    RegistryContract private registryContract;
    DistKeyContract private distKeyContract;

    constructor(address _registryContract, address _distKeyContract) public {
        registryContract = RegistryContract(_registryContract);
        distKeyContract = DistKeyContract(_distKeyContract);
    }

    function validateTransaction(bytes32 _hash, uint256 _confirmations)
        external
        payable
    {
        require(msg.value >= TOTAL_FEE, "msg value below total fee");
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

        RegistryContract.OracleNode memory aggregator =
            registryContract.getAggregator();
        require(aggregator.addr == msg.sender, "not the aggregator");

        uint256[2] memory hash =
            BN256G1.hashToPointSha256(abi.encode(_id, _result));
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
                G2_NEG_X_RE,
                G2_NEG_X_IM,
                G2_NEG_Y_RE,
                G2_NEG_Y_IM
            ];
        require(BN256G1.bn256CheckPairing(input), "invalid signature");

        uint256 fee = calculateFee(aggregator.stake, _signature[0]);
        payable(msg.sender).transfer(fee);

        ValidationResult storage validationResult = validationResults[_id];
        validationResult.id = _id;
        validationResult.result = _result;

        emit SubmitValidationResultLog(msg.sender, _id, _result, fee);
    }

    function calculateFee(uint256 stake, uint256 seed)
        private
        returns (uint256)
    {
        if (isValidationFeeReceiver(stake, seed)) {
            uint256 totalFee =
                BASE_FEE + VALIDATOR_FEE * requestsSinceLastPayout;
            requestsSinceLastPayout = 0;
            return totalFee;
        }
        requestsSinceLastPayout++;
        return BASE_FEE;
    }

    function isValidationFeeReceiver(uint256 stake, uint256 seed)
        public
        pure
        returns (bool)
    {
        uint256 scalingFactor = (stake / 1 ether)**2; //the exp should be smaller something like 10/9
        return (seed % (1000 / scalingFactor)) == 1;
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
