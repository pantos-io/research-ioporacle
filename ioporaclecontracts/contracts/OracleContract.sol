// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./RegistryContract.sol";
import "./DistKeyContract.sol";
import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";

contract OracleContract {

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

    mapping(bytes32 => bool) private blockValidationResults;
    mapping(bytes32 => bool) private txValidationResults;

    enum ValidationType { UNKNOWN, BLOCK, TRANSACTION }

    event ValidationRequest(ValidationType typ, address indexed from, bytes32 hash);

    event ValidationResponse(
        ValidationType typ,
        address indexed aggregator,
        bytes32 hash,
        bool valid,
        uint256 fee
    );

    RegistryContract private registryContract;
    DistKeyContract private distKeyContract;

    constructor(address _registryContract, address _distKeyContract) {
        registryContract = RegistryContract(_registryContract);
        distKeyContract = DistKeyContract(_distKeyContract);
    }

    modifier minFee(uint _min) {
        require(msg.value >= _min, "too few fee amount");
        _;
    }

    function validateBlock(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _hash);
    }

    function validateTransaction(bytes32 _hash) external payable minFee(TOTAL_FEE) {
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _hash);
    }

    function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] calldata _signature) external {
        submitValidationResult(ValidationType.BLOCK, _hash, _result, _signature);
    }

    function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] calldata _signature) external {
        submitValidationResult(ValidationType.TRANSACTION, _hash, _result, _signature);
    }

    function submitValidationResult(
        ValidationType _typ,
        bytes32 _hash,
        bool _result,
        uint256[2] calldata _signature
    ) private {

        require(_typ != ValidationType.UNKNOWN, "unknown validation type");

        RegistryContract.OracleNode memory aggregator =
            registryContract.getAggregator();
        require(aggregator.addr == msg.sender, "not the aggregator");

        uint256[2] memory hash =
            BN256G1.hashToPointSha256(abi.encode(_hash, _result, _typ));
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

        if (_typ == ValidationType.BLOCK) {
            blockValidationResults[_hash] = _result;
        } else if (_typ == ValidationType.TRANSACTION) {
            txValidationResults[_hash] = _result;
        }

        emit ValidationResponse(_typ, msg.sender, _hash, _result, fee);
    }

    function calculateFee(uint256 _stake, uint256 _seed) private returns (uint256) {
        if (isValidationFeeReceiver(_stake, _seed)) {
            uint256 totalFee =
                BASE_FEE + VALIDATOR_FEE * requestsSinceLastPayout;
            requestsSinceLastPayout = 0;
            return totalFee;
        }
        requestsSinceLastPayout++;
        return BASE_FEE;
    }

    function isValidationFeeReceiver(uint256 _stake, uint256 _seed) public pure returns (bool) {
        uint256 scalingFactor = (_stake / 1 ether)**2; //the exp should be smaller something like 10/9
        return (_seed % (1000 / scalingFactor)) == 1;
    }

    function findBlockValidationResult(bytes32 _hash) public view returns (bool) {
        return blockValidationResults[_hash];
    }

    function findTransactionValidationResult(bytes32 _hash) public view returns (bool) {
        return txValidationResults[_hash];
    }
}
