// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DistKeyContract.sol";

contract RegistryContract {
    struct OracleNode {
        address addr;
        string ipAddr;
        bytes pubKey;
        uint256 stake;
        uint256 index;
    }

    uint256 private constant BLOCK_RANGE = 10000000000000000;
    uint256 public constant MIN_STAKE = 0;

    mapping(address => OracleNode) private oracleNodes;
    address[] private oracleNodeIndices;

    event RegisterOracleNode(address indexed sender);

    DistKeyContract private distKeyContract;

    constructor(address _distKeyContract) {
        distKeyContract = DistKeyContract(_distKeyContract);
    }

    function registerOracleNode(string calldata _ipAddr, bytes calldata _pubKey)
        external
        payable
    {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        require(msg.value >= MIN_STAKE, "min stake too low");

        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.pubKey = _pubKey;
        iopNode.stake = msg.value;
        iopNode.index = oracleNodeIndices.length;
        oracleNodeIndices.push(iopNode.addr);

        emit RegisterOracleNode(msg.sender);

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

    function isAggregator(address _addr) public view returns (bool) {
        /**
         * As this method is a call, block.number refers to an already finalized block.
         * But aggregators submit transactions that should be included in a new block,
         * therefore the information of the aggregator of a finalized block is outdated.
         */
        return isAggregatorByBlock(_addr, block.number + 1);
    }

    function isAggregatorByBlock(address _addr, uint256 _block)
        public
        view
        returns (bool)
    {
        OracleNode memory iopNode = findOracleNodeByAddress(_addr);
        uint256 chosen = blockNumberMod(_block);
        return
            chosen >= iopNode.index * BLOCK_RANGE &&
            chosen < (iopNode.index + 1) * BLOCK_RANGE;
    }

    function getAggregator() public view returns (OracleNode memory) {
        /**
         * As this method is a call, block.number refers to an already finalized block.
         * But aggregators submit transactions that should be included in a new block,
         * therefore the information of the aggregator of a finalized block is outdated.
         */
        return getAggregatorByBlock(block.number + 1);
    }

    function getAggregatorByBlock(uint256 _block)
        public
        view
        returns (OracleNode memory)
    {
        uint256 i = blockNumberMod(_block) / BLOCK_RANGE;
        return findOracleNodeByIndex(i);
    }

    function blockNumberMod(uint256 _block) internal view returns (uint256) {
        return _block % (oracleNodeIndices.length * BLOCK_RANGE);
    }
}
