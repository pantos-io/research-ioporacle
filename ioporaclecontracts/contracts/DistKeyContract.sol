// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";

contract DistKeyContract {
    uint256 public constant KEY_GEN_INTERVAL = 3;

    event DistKeyGenerationLog(uint256 threshold);

    uint256[4] private publicKey;
    uint256 private numberOfValidators;

    RegistryContract private registryContract;

    function generate() external {
        require(msg.sender == address(registryContract), "invalid sender");
        emit DistKeyGenerationLog(signatureThreshold());
    }

    function signatureThreshold() private view returns (uint256) {
        return (registryContract.countOracleNodes() + 1) / 2;
    }

    function validatorThreshold() private view returns (uint256) {
        uint256 threshold = signatureThreshold();
        return threshold + ((registryContract.countOracleNodes() - threshold) / 2);
    }

    function getPublicKey() external view returns (uint256[4] memory) {
        return publicKey;
    }

    function getNumberOfValidators() external view returns (uint256) {
        return numberOfValidators;
    }

    function setPublicKey(uint256[4] calldata _publicKey, uint256 _numberOfValidators) external {
        require(_numberOfValidators >= validatorThreshold(), "too few validators");
        publicKey = _publicKey;
        numberOfValidators = _numberOfValidators;
    }

    function setRegistryContract(address _registryContract) external {
        require(
            address(registryContract) == address(0),
            "registry contract already set"
        );
        registryContract = RegistryContract(_registryContract);
    }
}
