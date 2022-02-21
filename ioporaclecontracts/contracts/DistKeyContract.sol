// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RegistryContract.sol";

contract DistKeyContract {
    uint256 public constant KEY_GEN_INTERVAL = 3;
    uint256 public constant KEY_FINAL_TIME = 0;

    event DistKeyGenerationLog(uint256 threshold);

    uint256 private signatureThreshold;
    uint256 private validatorThreshold;
    bool private generating;

    uint256[4] private publicKey;
    uint256 private numberOfValidators;

    address[] private disputes;
    uint256 private publicKeyFinal;

    RegistryContract private registryContract;

    function generate() public {
        require(msg.sender == address(registryContract), "invalid sender");
        generateKey();
    }

    function generateKey() private {
        uint256 noNodes = registryContract.countOracleNodes();
        signatureThreshold = (noNodes + 1) / 2;
        validatorThreshold =
            signatureThreshold +
            (noNodes - signatureThreshold) /
            2;

        generating = true;
        emit DistKeyGenerationLog(signatureThreshold);
    }

    function getNumberOfValidators() external view returns (uint256) {
        return numberOfValidators;
    }

    function getPublicKey() external view returns (uint256[4] memory) {
        require(block.number >= publicKeyFinal, "pubKey not final");
        return publicKey;
    }

    function setPublicKey(
        uint256[4] calldata _publicKey,
        uint256 _numberOfValidators
    ) external {
        require(generating, "not generating");
        require(
            registryContract.oracleNodeIsRegistered(msg.sender),
            "not registered"
        );
        require(
            _numberOfValidators >= validatorThreshold,
            "too few validators"
        );

        publicKey = _publicKey;
        numberOfValidators = _numberOfValidators;

        publicKeyFinal = block.number + KEY_FINAL_TIME;
        generating = false;
        delete disputes;
    }

    function disputePublicKey() external {
        require(
            registryContract.oracleNodeIsRegistered(msg.sender),
            "not registered"
        );
        require(block.number < publicKeyFinal, "key already final");
        for (uint256 i = 0; i < disputes.length; i++) {
            require(disputes[i] != msg.sender, "dispute already issued");
        }
        disputes.push(msg.sender);
        if (disputes.length >= signatureThreshold) {
            delete publicKeyFinal;
            delete disputes;
            generateKey();
        }
    }

    function setRegistryContract(address _registryContract) external {
        require(
            address(registryContract) == address(0),
            "registry contract already set"
        );
        registryContract = RegistryContract(_registryContract);
    }
}
