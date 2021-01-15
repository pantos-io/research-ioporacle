// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

contract RegistryContract {
    struct OracleNode {
        address addr;
        string ipAddr;
        bytes pubKey;
        uint256 index;
    }

    uint256 private constant BLOCK_RANGE = 6;

    mapping(address => OracleNode) private oracleNodes;
    address[] private oracleNodeIndices;

    event RegisterOracleNodeLog(address indexed sender);

    function registerOracleNode(
        string calldata _ipAddr,
        bytes calldata _pubKey
    ) external payable {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.pubKey = _pubKey;
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
        public
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

    function isLeader(address addr) public view returns (bool) {
        OracleNode memory iopNode = findOracleNodeByAddress(addr);
        uint256 chosen =
            block.number % (oracleNodeIndices.length * BLOCK_RANGE);
        return
            chosen >= iopNode.index * BLOCK_RANGE &&
            chosen < (iopNode.index + 1) * BLOCK_RANGE;
    }

    function getLeader() public view returns (OracleNode memory) {
        for (uint256 i = 0; i < oracleNodeIndices.length; i++) {
            OracleNode memory iopNode = oracleNodes[oracleNodeIndices[i]];
            if (isLeader(iopNode.addr)) {
                return iopNode;
            }
        }
        revert("no leader");
    }
}
