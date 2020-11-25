pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

contract OracleContract {
    struct OracleNode {
        address addr;
        string ipAddr;
        uint256 index;
    }

    uint256 private constant BLOCK_RANGE = 6;

    mapping(address => OracleNode) private oracleNodes;
    address[] private oracleNodeIndices;

    bool public result;
    event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations);

    function registerOracleNode(string calldata _ipAddr) external payable {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.index = oracleNodeIndices.length;
        oracleNodeIndices.push(iopNode.addr);
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
    {
        emit VerifyTransactionLog(1, _hash, _confirmations);
    }

    function submitVerification(bool _result) external {
        result = _result;
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
