const RegistryContract = artifacts.require("RegistryContract");
const OracleContract = artifacts.require("OracleContract");
const DistKeyContract = artifacts.require("DistKeyContract");

const ECDSAOracleContract = artifacts.require("ECDSAOracleContract");

module.exports = function (deployer) {
  let distKeyContract;
  deployer
    .deploy(DistKeyContract)
    .then(function () {
      return deployer.deploy(RegistryContract, DistKeyContract.address);
    })
    .then(function () {
      return DistKeyContract.deployed().then(function (instance) {
        distKeyContract = instance;
      });
    })
    .then(function () {
      return deployer.deploy(RegistryContract, DistKeyContract.address);
    })
    .then(function () {
      distKeyContract.setRegistryContract(RegistryContract.address);
      return deployer.deploy(
        OracleContract,
        RegistryContract.address,
        DistKeyContract.address
      );
    })
    .then(function () {
      return deployer.deploy(ECDSAOracleContract, RegistryContract.address);
    });
};
