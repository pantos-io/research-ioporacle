const RegistryContract = artifacts.require("RegistryContract");
const OracleContract = artifacts.require("OracleContract");
const RaffleContract = artifacts.require("RaffleContract");

module.exports = function (deployer) {
  deployer
    .deploy(RegistryContract)
    .then(function () {
      return deployer.deploy(RaffleContract, RegistryContract.address);
    })
    .then(function () {
      return deployer.deploy(
        OracleContract,
        RegistryContract.address,
        RaffleContract.address
      );
    });
};
