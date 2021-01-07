// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "cdlbot-solidity/contracts/crypto/MerkleRoot.sol";

contract MerkleOracleContract {
    using MerkleRoot for bytes32[];

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
        bytes32 merkleRoot;
        uint256 finalizedAt;
        address sender;
        address rewardRecipient;
        bool rewardTransfered;
    }

    uint256 private constant PRICE = 0.001 ether;
    uint256 private constant BLOCKS_UNTIL_FINALIZED = 6;

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
        bool result,
        address[] witnesses,
        bytes32 merkleRoot,
        address rewardRecipient
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
        address[] memory _witnesses
    ) public {
        require(registryContract.isLeader(msg.sender), "not the leader");

        bytes32[] memory leaves = new bytes32[](_witnesses.length);
        for (uint256 i = 0; i < leaves.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(_witnesses[i]));
        }
        bytes32 merkleRoot = leaves.calculate();

        uint256 id = ++resultCounter;

        VerificationResult storage verificationResult = verificationResults[id];
        verificationResult.id = id;
        verificationResult.request = _id;
        verificationResult.result = _result;
        verificationResult.merkleRoot = merkleRoot;
        verificationResult.finalizedAt = block.number + BLOCKS_UNTIL_FINALIZED;
        verificationResult.sender = msg.sender;

        uint256 rewardedWitnessIndex =
            uint256(keccak256(abi.encode(block.timestamp, block.difficulty))) %
                _witnesses.length;
        verificationResult.rewardRecipient = _witnesses[rewardedWitnessIndex];

        emit SubmitVerificationLog(
            msg.sender,
            id,
            _id,
            _result,
            _witnesses,
            verificationResult.merkleRoot,
            verificationResult.rewardRecipient
        );
    }

    function transferRewards(uint256 _id) public {
        VerificationResult memory verificationResult = findVerification(_id);
        require(!verificationResult.rewardTransfered, "already transfered");

        verificationResult.rewardTransfered = true;
        verificationResults[verificationResult.id] = verificationResult;

        address(uint160(verificationResult.sender)).transfer(PRICE);
        address(uint160(verificationResult.rewardRecipient)).transfer(PRICE);
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
}
