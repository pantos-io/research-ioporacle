// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "../RegistryContract.sol";

contract OnChainOracleContract {
    struct ValidationResult {
        uint256 id;
        bool result;
    }

    uint256 public constant FEE = 0.001 ether;

    uint256 private requestCounter;

    mapping(uint256 => mapping(bool => address[])) private votes;

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

    constructor(address _registryContract) public {
        registryContract = RegistryContract(_registryContract);
    }

    function validateTransaction(bytes32 _hash, uint256 _confirmations)
        external
        payable
    {
        require(
            msg.value >= FEE * registryContract.countOracleNodes(),
            "msg value below total fee"
        );
        emit ValidateTransactionLog(
            msg.sender,
            ++requestCounter,
            _hash,
            _confirmations
        );
    }

    function submitValidationResult(uint256 _id, bool _result) public {
        require(
            registryContract.oracleNodeIsRegistered(msg.sender),
            "not registered"
        );

        votes[_id][_result].push(msg.sender);

        uint256 majority = registryContract.countOracleNodes() / 2 + 1;
        if (votes[_id][_result].length >= majority) {
            ValidationResult storage validationResult = validationResults[_id];
            validationResult.id = _id;
            validationResult.result = _result;

            for (uint256 i = 0; i < votes[_id][_result].length; i++) {
                address(uint160(votes[_id][_result][i])).transfer(FEE);
            }

            emit SubmitValidationResultLog(msg.sender, _id, _result);
        }
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
