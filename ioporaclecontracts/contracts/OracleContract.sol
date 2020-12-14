// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "@openzeppelin/contracts/cryptography/ECDSA.sol";

contract OracleContract {
    using ECDSA for bytes32;

    struct VerificationRequest {
        uint256 id;
        string hash;
        uint256 confirmations;
        uint256 index;
    }

    struct VerificationResult {
        uint256 id;
        uint256 request;
        bool result;
    }

    uint256 private constant PRICE = 0.001 ether;

    mapping(uint256 => VerificationRequest) private verificationRequests;
    uint256[] private verificationRequestIndices;
    uint256 private requestCounter;

    mapping(uint256 => VerificationResult) private verificationResults;
    uint256 private resultCounter;

    event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations);
    event SubmitVerificationLog(
        address indexed sender,
        uint256 id,
        uint256 request,
        bool result
    );

    RegistryContract private registryContract;

    constructor(address _registryContract) public {
        registryContract = RegistryContract(_registryContract);
    }

    function verifyTransaction(string calldata _hash, uint256 _confirmations)
        external
        payable
    {
        require(
            msg.value >= registryContract.countOracleNodes() * PRICE,
            "msg value too low"
        );
        uint256 id = ++requestCounter;
        VerificationRequest storage verificationRequest =
            verificationRequests[id];
        verificationRequest.id = id;
        verificationRequest.hash = _hash;
        verificationRequest.confirmations = _confirmations;
        verificationRequest.index = verificationRequestIndices.length;
        verificationRequestIndices.push(id);
        emit VerifyTransactionLog(id, _hash, _confirmations);
    }

    function submitVerification(
        uint256 _id,
        bool _result,
        bytes[] memory _signatures
    ) public {
        require(registryContract.isLeader(msg.sender), "not the leader");

        uint256 id = ++resultCounter;

        VerificationResult storage verificationResult = verificationResults[id];
        verificationResult.id = id;
        verificationResult.request = _id;
        verificationResult.result = _result;

        bytes32 hash = keccak256(encodeResult(_id, _result));

        uint256 witnessCount = 0;
        for (uint256 i = 0; i < _signatures.length; i++) {
            address witness = hash.recover(_signatures[i]);
            if (registryContract.oracleNodeIsRegistered(witness)) {
                witnessCount += 1;
                address(uint160(witness)).transfer(PRICE);
            }
        }

        uint256 majority = (registryContract.countOracleNodes() + 1) / 2;
        require(witnessCount > majority, "no majority");

        emit SubmitVerificationLog(msg.sender, id, _id, _result);
    }

    function findVerification(uint256 _id)
        public
        view
        returns (VerificationResult memory)
    {
        require(verificationExists(_id), "not found");
        VerificationResult memory verificationResult = verificationResults[_id];
        return verificationResult;
    }

    function verificationExists(uint256 _id) public view returns (bool) {
        return verificationResults[_id].id != 0;
    }

    function encodeResult(uint256 _id, bool _result)
        public
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_id, _result);
    }
}
