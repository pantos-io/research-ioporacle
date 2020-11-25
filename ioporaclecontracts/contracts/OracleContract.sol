pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

contract OracleContract {
    struct IopNode {
        address addr;
        string ipAddr;
        uint256 index;
    }

    uint256 private constant BLOCK_RANGE = 6;

    mapping(address => IopNode) private iopNodes;
    address[] private iopNodeIndices;

    bool public result;
    event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations);
    event Test(uint256 a);

    function registerIopNode(string calldata _ipAddr) external payable {
        require(!iopNodeIsRegistered(msg.sender), "already registered");
        IopNode storage iopNode = iopNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.index = iopNodeIndices.length;
        iopNodeIndices.push(iopNode.addr);
    }

    function iopNodeIsRegistered(address _addr) public view returns (bool) {
        if (iopNodeIndices.length == 0) return false;
        return (iopNodeIndices[iopNodes[_addr].index] == _addr);
    }

    function findIopNodeByAddress(address _addr)
        public
        view
        returns (IopNode memory)
    {
        require(iopNodeIsRegistered(_addr), "not found");
        IopNode memory iopNode = iopNodes[_addr];
        return iopNode;
    }

    function findIopNodeByIndex(uint256 _index)
        external
        view
        returns (IopNode memory)
    {
        require(_index >= 0 && _index < iopNodeIndices.length, "not found");
        IopNode memory iopNode = iopNodes[iopNodeIndices[_index]];
        return iopNode;
    }

    function countIopNodes() external view returns (uint256) {
        return iopNodeIndices.length;
    }

    function verifyTransaction(string calldata _hash, uint256 _confirmations)
        external
    {
        emit VerifyTransactionLog(1, _hash, _confirmations);
    }

    function verifyTransactionResult(bool _result) external {
        result = _result;
    }

    function isLeader(address addr) public view returns (bool) {
        IopNode memory iopNode = findIopNodeByAddress(addr);
        uint256 chosen = block.number % (iopNodeIndices.length * BLOCK_RANGE);
        return
            chosen >= iopNode.index * BLOCK_RANGE &&
            chosen < (iopNode.index + 1) * BLOCK_RANGE;
    }
}
