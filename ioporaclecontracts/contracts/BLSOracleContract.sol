// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;
pragma experimental ABIEncoderV2;

import "./RegistryContract.sol";
import "cdlbot-solidity/contracts/crypto/BN256G1.sol";
import "cdlbot-solidity/contracts/crypto/BN256G2.sol";

contract BLSOracleContract {
    struct VerificationRequest {
        uint256 id;
        string hash;
        uint256 confirmations;
        uint256 index;
    }

    struct VerificationResult {
        uint256 id;
        uint256 request;
        bool result;
    }

    uint256 private constant PRICE = 0.001 ether;

    uint256 public constant PUB_KEY_X_RE =
        18530246154675985083946582251796722642614829756181350107086775621752958860856;
    uint256 public constant PUB_KEY_X_IM =
        13000374846473950965678136334447345075094988074199869611261273825039642148401;
    uint256 public constant PUB_KEY_Y_RE =
        3884615071854075031801167849596201425084295738405003455544554117442171322358;
    uint256 public constant PUB_KEY_Y_IM =
        11060966608439035944330745442774880538039124395544732940162528622795873633998;

    mapping(uint256 => VerificationRequest) private verificationRequests;
    uint256[] private verificationRequestIndices;
    uint256 private requestCounter;

    mapping(uint256 => VerificationResult) private verificationResults;
    uint256 private resultCounter;

    event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations);
    event SubmitVerificationLog(
        address indexed sender,
        uint256 id,
        uint256 request,
        bool result
    );

    RegistryContract private registryContract;

    constructor(address _registryContract) public {
        registryContract = RegistryContract(_registryContract);
    }

    function verifyTransaction(string calldata _hash, uint256 _confirmations)
        external
        payable
    {
        require(msg.value >= PRICE, "msg value too low");

        uint256 id = ++requestCounter;

        VerificationRequest storage verificationRequest =
            verificationRequests[id];
        verificationRequest.id = id;
        verificationRequest.hash = _hash;
        verificationRequest.confirmations = _confirmations;
        verificationRequest.index = verificationRequestIndices.length;
        verificationRequestIndices.push(id);

        emit VerifyTransactionLog(id, _hash, _confirmations);
    }

    function submitVerification(
        uint256 _id,
        bool _result,
        uint256[2] memory signature
    ) public {
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
                signature[0],
                signature[1],
                negG2[0],
                negG2[1],
                negG2[2],
                negG2[3]
            ];
        require(BN256G1.bn256CheckPairing(input), "invalid signature");

        uint256 id = ++resultCounter;

        VerificationResult storage verificationResult = verificationResults[id];
        verificationResult.id = id;
        verificationResult.request = _id;
        verificationResult.result = _result;

        emit SubmitVerificationLog(msg.sender, id, _id, _result);
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
