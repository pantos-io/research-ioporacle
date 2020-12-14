const OracleContract = artifacts.require("OracleContract");
const RegistryContract = artifacts.require("RegistryContract");

module.exports = function(deployer) {
  deployer.deploy(RegistryContract).then(function() {
    return deployer.deploy(OracleContract, RegistryContract.address);
  });
};
