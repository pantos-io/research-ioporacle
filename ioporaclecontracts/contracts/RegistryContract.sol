// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./DistKeyContract.sol";

contract RegistryContract {
    struct OracleNode {
        address addr;
        string ipAddr;
        bytes pubKey;
        uint256 index;
    }

    uint256 private constant BLOCK_RANGE = 6;
    uint256 private constant KEY_GEN_INTERVAL = 3;

    mapping(address => OracleNode) private oracleNodes;
    address[] private oracleNodeIndices;

    event RegisterOracleNodeLog(address indexed sender);

    DistKeyContract private distKeyContract;

    constructor(address _distKeyContract) public {
        distKeyContract = DistKeyContract(_distKeyContract);
    }

    function registerOracleNode(string calldata _ipAddr, bytes calldata _pubKey)
        external
        payable
    {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.pubKey = _pubKey;
        iopNode.index = oracleNodeIndices.length;
        oracleNodeIndices.push(iopNode.addr);

        emit RegisterOracleNodeLog(msg.sender);

        if (
            oracleNodeIndices.length % distKeyContract.KEY_GEN_INTERVAL() == 0
        ) {
            distKeyContract.generate();
        }
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
        return oracleNodes[_addr];
    }

    function findOracleNodeByIndex(uint256 _index)
        public
        view
        returns (OracleNode memory)
    {
        require(_index >= 0 && _index < oracleNodeIndices.length, "not found");
        return oracleNodes[oracleNodeIndices[_index]];
    }

    function countOracleNodes() external view returns (uint256) {
        return oracleNodeIndices.length;
    }

    function isAggregator(address addr) public view returns (bool) {
        OracleNode memory iopNode = findOracleNodeByAddress(addr);
        uint256 chosen = blockNumberMod();
        return
            chosen >= iopNode.index * BLOCK_RANGE &&
            chosen < (iopNode.index + 1) * BLOCK_RANGE;
    }

    function getAggregator() public view returns (OracleNode memory) {
        uint256 i = blockNumberMod() / BLOCK_RANGE;
        return findOracleNodeByIndex(i);
    }

    function blockNumberMod() internal view returns (uint256) {
        return block.number % (oracleNodeIndices.length * BLOCK_RANGE);
    }
}
