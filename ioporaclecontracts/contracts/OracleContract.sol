// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "cdlbot-solidity/contracts/crypto/BN256G1.sol";
import "cdlbot-solidity/contracts/crypto/BN256G2.sol";

contract OracleContract {
    struct VerificationResult {
        uint256 id;
        bool result;
    }

    uint256 private constant LEADER_REWARD = 0.001 ether;
    uint256 private constant RANDOM_NODE_REWARD = 0.001 ether;

    uint256 public constant PUB_KEY_X_RE =
        8886964294152041318536467674815330882393317464853688366853649667782743822558;
    uint256 public constant PUB_KEY_X_IM =
        6307524910493884348606458859953541626125901056736784276770927119641109937230;
    uint256 public constant PUB_KEY_Y_RE =
        12945817349354013390550542415775465846478641491902578740035047875702529653965;
    uint256 public constant PUB_KEY_Y_IM =
        5515339196438690171267452832958689363782700262658561041524067975725456384728;

    uint256 private requestCounter;

    mapping(uint256 => VerificationResult) private verificationResults;

    event VerifyTransactionLog(
        address indexed sender,
        uint256 indexed id,
        bytes32 hash,
        uint256 confirmations
    );
    event SubmitVerificationLog(
        address indexed sender,
        uint256 id,
        bool result
    );

    RegistryContract private registryContract;

    constructor(address _registryContract) public {
        registryContract = RegistryContract(_registryContract);
    }

    function verifyTransaction(bytes32 _hash, uint256 _confirmations)
        external
        payable
    {
        require(
            msg.value >= LEADER_REWARD + RANDOM_NODE_REWARD,
            "msg value too low"
        );
        emit VerifyTransactionLog(
            msg.sender,
            ++requestCounter,
            _hash,
            _confirmations
        );
    }

    function submitVerification(
        uint256 _id,
        bool _result,
        uint256[2] calldata _signature
    ) external {
        require(!verificationExists(_id), "already submitted");

        uint256[2] memory hash =
            BN256G1.hashToPointSha256(abi.encode(_id, _result));
        uint256[4] memory negG2 = BN256G2.getNegG2();
        uint256[12] memory input =
            [
                hash[0],
                hash[1],
                PUB_KEY_X_RE,
                PUB_KEY_X_IM,
                PUB_KEY_Y_RE,
                PUB_KEY_Y_IM,
                _signature[0],
                _signature[1],
                negG2[0],
                negG2[1],
                negG2[2],
                negG2[3]
            ];
        require(BN256G1.bn256CheckPairing(input), "invalid signature");

        payable(msg.sender).transfer(LEADER_REWARD);

        VerificationResult storage verificationResult =
            verificationResults[_id];
        verificationResult.id = _id;
        verificationResult.result = _result;

        emit SubmitVerificationLog(msg.sender, _id, _result);
    }

    function findVerification(uint256 _id)
        public
        view
        returns (VerificationResult memory)
    {
        require(verificationExists(_id), "not found");
        VerificationResult memory verificationResult = verificationResults[_id];
        return verificationResult;
    }

    function verificationExists(uint256 _id) public view returns (bool) {
        return verificationResults[_id].id != 0;
    }
}
