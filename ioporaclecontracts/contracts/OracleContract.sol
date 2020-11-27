// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

contract OracleContract {
    struct OracleNode {
        address addr;
        string ipAddr;
        uint256 index;
    }

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
        address[] witnesses;
        uint256 finalized;
        bool rewardTransfered;
    }

    uint256 private constant BLOCK_RANGE = 6;
    uint256 private constant FINAL_RANGE = 6;
    uint256 private constant PRICE = 0.001 ether;

    mapping(address => OracleNode) private oracleNodes;
    address[] private oracleNodeIndices;

    mapping(uint256 => VerificationRequest) private verificationRequests;
    uint256[] private verificationRequestIndices;
    uint256 private requestCounter;

    mapping(uint256 => VerificationResult) private verificationResults;
    uint256 private resultCounter;

    event RegisterOracleNodeLog(address indexed sender);
    event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations);
    event SubmitVerificationLog(
        address indexed sender,
        uint256 id,
        uint256 request,
        bool result,
        address[] witnesses,
        uint256 finalized
    );
    event TransferRewardLog(uint256 indexed request);

    function registerOracleNode(string calldata _ipAddr) external payable {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.index = oracleNodeIndices.length;
        oracleNodeIndices.push(iopNode.addr);
        emit RegisterOracleNodeLog(msg.sender);
    }

    function oracleNodeIsRegistered(address _addr) public view returns (bool) {
        if (oracleNodeIndices.length == 0) return false;
        return (oracleNodeIndices[oracleNodes[_addr].index] == _addr);
    }

    function findOracleNodeByAddress(address _addr)
        public
        view
        returns (OracleNode memory)
    {
        require(oracleNodeIsRegistered(_addr), "not found");
        OracleNode memory iopNode = oracleNodes[_addr];
        return iopNode;
    }

    function findOracleNodeByIndex(uint256 _index)
        external
        view
        returns (OracleNode memory)
    {
        require(_index >= 0 && _index < oracleNodeIndices.length, "not found");
        OracleNode memory iopNode = oracleNodes[oracleNodeIndices[_index]];
        return iopNode;
    }

    function countOracleNodes() external view returns (uint256) {
        return oracleNodeIndices.length;
    }

    function verifyTransaction(string calldata _hash, uint256 _confirmations)
        external
        payable
    {
        require(
            msg.value >= oracleNodeIndices.length * PRICE,
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
        uint256 id = ++resultCounter;
        VerificationResult storage verificationResult = verificationResults[id];
        verificationResult.id = id;
        verificationResult.request = _id;
        verificationResult.result = _result;
        verificationResult.witnesses = _witnesses;
        verificationResult.finalized = block.number + FINAL_RANGE;
        emit SubmitVerificationLog(
            msg.sender,
            id,
            _id,
            _result,
            _witnesses,
            verificationResult.finalized
        );
    }

    function transferReward(uint256 _id) external {
        VerificationResult memory result = findVerification(_id);
        require(result.finalized <= block.number, "not finalized");
        require(!result.rewardTransfered, "reward already transfered");
        result.rewardTransfered = true;
        verificationResults[result.id] = result;
        for (uint256 i = 0; i < result.witnesses.length; i++) {
            address(uint160(result.witnesses[i])).transfer(PRICE);
        }
        emit TransferRewardLog(_id);
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

    function isLeader(address addr) public view returns (bool) {
        OracleNode memory iopNode = findOracleNodeByAddress(addr);
        uint256 chosen =
            block.number % (oracleNodeIndices.length * BLOCK_RANGE);
        return
            chosen >= iopNode.index * BLOCK_RANGE &&
            chosen < (iopNode.index + 1) * BLOCK_RANGE;
    }
}
