// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

contract Migrations {
    address public owner = msg.sender;
    // solhint-disable-next-line var-name-mixedcase
    uint256 public last_completed_migration;

    modifier restricted() {
        require(msg.sender == owner, "restricted to contract's owner");
        _;
    }

    function setCompleted(uint256 completed) public restricted {
        last_completed_migration = completed;
    }
}
