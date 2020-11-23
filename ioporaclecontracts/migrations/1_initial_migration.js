const Migrations = artifacts.require("Migrations");
const IopOracleContract = artifacts.require("IopOracleContract");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(IopOracleContract);
};
