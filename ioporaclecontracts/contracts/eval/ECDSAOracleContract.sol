// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../RegistryContract.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract ECDSAOracleContract {
    using ECDSA for bytes32;

    struct ValidationResult {
        uint256 id;
        bool result;
    }

    uint256 public constant FEE = 0.001 ether;

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

    constructor(address _registryContract) {
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

    function submitValidationResult(
        uint256 _id,
        bool _result,
        bytes[] memory _signatures
    ) public {
        require(!validationResultExists(_id), "already exists");
        require(
            registryContract.isAggregator(msg.sender),
            "not the aggregator"
        );

        ValidationResult storage validationResult = validationResults[_id];
        validationResult.id = _id;
        validationResult.result = _result;

        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 result = keccak256(encodeResult(_id, _result));
        bytes32 hash = keccak256(abi.encodePacked(prefix, result));

        address[] memory witnesses = new address[](_signatures.length);
        for (uint256 i = 0; i < _signatures.length; i++) {
            address witness = hash.recover(_signatures[i]);
            require(
                registryContract.oracleNodeIsRegistered(witness),
                "witness not registered"
            );
            require(
                !containsWitness(witnesses, witness),
                "witness already included"
            );
            witnesses[i] = witness;
            payable(address(uint160(witness))).transfer(FEE);
        }

        uint256 majority = registryContract.countOracleNodes() / 2 + 1;
        require(witnesses.length >= majority, "no majority");

        emit SubmitValidationResultLog(msg.sender, _id, _result);
    }

    function containsWitness(address[] memory _witnesses, address _witness)
        private
        pure
        returns (bool)
    {
        for (uint256 i = 0; i < _witnesses.length; i++) {
            if (_witnesses[i] == _witness) {
                return true;
            }
        }
        return false;
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

    function encodeResult(uint256 _id, bool _result)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_id, _result);
    }
}
