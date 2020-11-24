const Migrations = artifacts.require("Migrations");
const OracleContract = artifacts.require("OracleContract");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(OracleContract);
};
