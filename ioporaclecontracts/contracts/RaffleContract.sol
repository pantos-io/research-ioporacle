// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";

contract RaffleContract {
    event ReceiveLog(address indexed sender, uint256 amount);

    RegistryContract private registryContract;

    constructor(address payable _registryContract) public {
        registryContract = RegistryContract(_registryContract);
    }

    receive() external payable {
        emit ReceiveLog(msg.sender, msg.value);
    }
}
