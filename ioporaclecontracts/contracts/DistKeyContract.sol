// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";

contract DistKeyContract {
    uint256 public constant KEY_GEN_INTERVAL = 3;

    event DistKeyGenerationLog(uint256 threshold);

    uint256[4] private publicKey;

    RegistryContract private registryContract;

    function generate() external {
        require(msg.sender == address(registryContract), "invalid sender");
        emit DistKeyGenerationLog(threshold());
    }

    function threshold() private view returns (uint256) {
        return (registryContract.countOracleNodes() + 1) / 2;
    }

    function getPublicKey() external view returns (uint256[4] memory) {
        return publicKey;
    }

    function setPublicKey(uint256[4] calldata _publicKey) external {
        publicKey = _publicKey;
    }

    function setRegistryContract(address _registryContract) external {
        require(
            address(registryContract) == address(0),
            "registry contract already set"
        );
        registryContract = RegistryContract(_registryContract);
    }
}
