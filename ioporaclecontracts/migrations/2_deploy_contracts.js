const RegistryContract = artifacts.require("RegistryContract");
const OracleContract = artifacts.require("OracleContract");

module.exports = function (deployer) {
  deployer.deploy(RegistryContract).then(function () {
    return deployer.deploy(OracleContract, RegistryContract.address);
  });
};
