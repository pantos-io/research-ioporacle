// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

abstract contract TxInclusionVerifier {
    function isBlockConfirmed(uint feeInWei, bytes32 blockHash, uint requiredConfirmations) payable public virtual returns (bool);

    function verifyTransaction(uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedTx,
        bytes memory path, bytes memory rlpEncodedNodes) payable public virtual returns (uint8);

    function verifyReceipt(uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedReceipt,
        bytes memory path, bytes memory rlpEncodedNodes) payable public virtual returns (uint8);

    function verifyState(uint feeInWei, bytes memory rlpHeader, uint8 noOfConfirmations, bytes memory rlpEncodedState,
        bytes memory path, bytes memory rlpEncodedNodes) payable public virtual returns (uint8);
}
