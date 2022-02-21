// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RegistryContract.sol";
import "./DistKeyContract.sol";
import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./TxInclusionVerifier.sol";

contract OracleContract is TxInclusionVerifier {
    struct ValidationResult {
        bytes32 _hash;
        bool result;
    }

    uint256 public constant BASE_FEE = 0.001 ether;
    uint256 public constant VALIDATOR_FEE = 0.0001 ether;
    uint256 public constant TOTAL_FEE = BASE_FEE + VALIDATOR_FEE;

    uint256 private constant G2_NEG_X_RE =
        0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 private constant G2_NEG_X_IM =
        0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 private constant G2_NEG_Y_RE =
        0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 private constant G2_NEG_Y_IM =
        0x1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d;

    uint256 private requestCounter;
    uint256 private requestsSinceLastPayout;

    mapping(bytes32 => ValidationResult) private txValidationResults;
    mapping(bytes32 => ValidationResult) private blockValidationResults;

    event ValidateTransactionLog(address indexed sender, bytes32 _hash);
    event ValidateBlockLog(address indexed sender, bytes32 _hash);

    event SubmitValidationResultLog(
        address indexed sender,
        bytes32 _hash,
        bool result,
        uint256 fee
    );

    RegistryContract private registryContract;
    DistKeyContract private distKeyContract;

    constructor(address _registryContract, address _distKeyContract) {
        registryContract = RegistryContract(_registryContract);
        distKeyContract = DistKeyContract(_distKeyContract);
    }

    modifier feeCheck {
        require(msg.value >= TOTAL_FEE, "msg value below total fee");
        _;
    }

    function validateTransaction(bytes32 _hash) external payable feeCheck {
        emit ValidateTransactionLog(msg.sender, _hash);
    }

    function validateBlock(bytes32 _hash) external payable feeCheck {
        emit ValidateBlockLog(msg.sender, _hash);
    }

    function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] calldata _signature) external {
        require(!txValidationResultExists(_hash), "already exists");
        submitValidationResult(_hash, _result, _signature, txValidationResults);
    }

    function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] calldata _signature) external {
        require(!blockValidationResultExists(_hash), "already exists");
        submitValidationResult(_hash, _result, _signature, blockValidationResults);
    }

    function submitValidationResult(
        bytes32 _hash,
        bool _result,
        uint256[2] calldata _signature,
        mapping(bytes32 => ValidationResult) storage validationResults
    ) private {
        RegistryContract.OracleNode memory aggregator =
            registryContract.getAggregator();
        require(aggregator.addr == msg.sender, "not the aggregator");

        uint256[2] memory hash =
            BN256G1.hashToPointSha256(abi.encode(_hash, _result));
        uint256[4] memory publicKey = distKeyContract.getPublicKey();
        uint256[12] memory input =
            [
                hash[0],
                hash[1],
                publicKey[0],
                publicKey[1],
                publicKey[2],
                publicKey[3],
                _signature[0],
                _signature[1],
                G2_NEG_X_RE,
                G2_NEG_X_IM,
                G2_NEG_Y_RE,
                G2_NEG_Y_IM
            ];
        require(BN256G1.bn256CheckPairing(input), "invalid signature");

        uint256 fee = calculateFee(aggregator.stake, _signature[0]);
        payable(msg.sender).transfer(fee);

        ValidationResult storage validationResult = validationResults[_hash];
        validationResult._hash = _hash;
        validationResult.result = _result;

        emit SubmitValidationResultLog(msg.sender, _hash, _result, fee);
    }

    function calculateFee(uint256 stake, uint256 seed) private returns (uint256) {
        if (isValidationFeeReceiver(stake, seed)) {
            uint256 totalFee =
                BASE_FEE + VALIDATOR_FEE * requestsSinceLastPayout;
            requestsSinceLastPayout = 0;
            return totalFee;
        }
        requestsSinceLastPayout++;
        return BASE_FEE;
    }

    function isValidationFeeReceiver(uint256 stake, uint256 seed) public pure returns (bool) {
        uint256 scalingFactor = (stake / 1 ether)**2; //the exp should be smaller something like 10/9
        return (seed % (1000 / scalingFactor)) == 1;
    }

    function findTxValidationResult(bytes32 _hash) public view returns (ValidationResult memory) {
        require(txValidationResultExists(_hash), "not found");
        return txValidationResults[_hash];
    }

    function findBlockValidationResult(bytes32 _hash) public view returns (ValidationResult memory) {
        require(blockValidationResultExists(_hash), "not found");
        return blockValidationResults[_hash];
    }

    function txValidationResultExists(bytes32 _hash) public view returns (bool) {
        return txValidationResults[_hash]._hash != 0;
    }

    function blockValidationResultExists(bytes32 _hash) public view returns (bool) {
        return blockValidationResults[_hash]._hash != 0;
    }

    function isBlockConfirmed(uint /*feeInWei*/, bytes32 blockHash, uint /*requiredConfirmations*/) payable public override returns (bool) {
        return findBlockValidationResult(blockHash).result;
    }

    function verifyTransaction(uint /*feeInWei*/, bytes memory /*rlpHeader*/, uint8 /*noOfConfirmations*/, bytes memory /*rlpEncodedTx*/,
        bytes memory /*path*/, bytes memory /*rlpEncodedNodes*/) payable public override returns (uint8) {
        return 0;
    }

    function verifyReceipt(uint /*feeInWei*/, bytes memory /*rlpHeader*/, uint8 /*noOfConfirmations*/, bytes memory /*rlpEncodedReceipt*/,
        bytes memory /*path*/, bytes memory /*rlpEncodedNodes*/) payable public override returns (uint8) {
        return 0;
    }

    function verifyState(uint /*feeInWei*/, bytes memory /*rlpHeader*/, uint8 /*noOfConfirmations*/, bytes memory /*rlpEncodedState*/,
        bytes memory /*path*/, bytes memory /*rlpEncodedNodes*/) payable public override returns (uint8) {
        return 0;
    }
}
