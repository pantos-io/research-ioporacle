pragma solidity >=0.4.22 <0.8.0;

contract IopOracle {
    struct IopNode {
        address addr;
        string ipAddr;
        uint256 index;
    }

    mapping(address => IopNode) private iopNodes;
    address[] private iopNodeIndices;

    function registerIopNode(string calldata _ipAddr) external payable {
        require(!iopNodeIsRegistered(msg.sender), "already registered");
        IopNode storage iopNode = iopNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.index = iopNodeIndices.push(iopNode.addr) - 1;
    }

    function iopNodeIsRegistered(address _addr) public view returns (bool) {
        if (iopNodeIndices.length == 0) return false;
        return (iopNodeIndices[iopNodes[_addr].index] == _addr);
    }

    function findIopNode(address _addr)
        external
        view
        returns (address, string memory)
    {
        require(iopNodeIsRegistered(_addr), "not found");
        IopNode memory iopNode = iopNodes[_addr];
        return (iopNode.addr, iopNode.ipAddr);
    }

    function countIopNodes() external view returns (uint256) {
        return iopNodeIndices.length;
    }
}
