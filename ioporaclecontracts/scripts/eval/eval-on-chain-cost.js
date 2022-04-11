const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const RegistryContract = artifacts.require("RegistryContract");
const DistKeyContract = artifacts.require("DistKeyContract");
const OnChainOracleContract = artifacts.require("OnChainOracleContract");

module.exports = async function () {
  let dir = "./data";
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir);
  }
  const csvWriter = createCsvWriter({
    path: "./data/on-chain-cost.csv",
    header: [
      { id: "no_nodes", title: "no_nodes" },
      { id: "gas", title: "gas" },
    ],
  });
  let records = [];

  let no_nodes = 3;

  while (no_nodes < 64) {
    let distKeyContract = await DistKeyContract.new();
    let registryContract = await RegistryContract.new(distKeyContract.address);
    let onChainOracleContract = await OnChainOracleContract.new(
      registryContract.address
    );
    await distKeyContract.setRegistryContract(registryContract.address);

    let accounts = await web3.eth.getAccounts();
    let stake = await registryContract.MIN_STAKE();
    for (let i = 0; i < no_nodes; i++) {
      await registryContract.registerOracleNode("127.0.0.1", "0x0", {
        from: accounts[i],
        value: stake,
      });
    }

    let fee = await onChainOracleContract.FEE();
    let fees = no_nodes * fee;
    await onChainOracleContract.validateTransaction(
      "0xa67220981e1760824947fb294f65adf47c505c3bfbe5960341d64c7f7512be8a",
      3,
      {
        value: fees,
      }
    );

    let majority = Math.trunc(no_nodes / 2) + 1;

    let gas = 0;
    for (let i = 0; i < majority; i++) {
      let tx = await onChainOracleContract.submitValidationResult(1, true, {
        from: accounts[i],
      });
      gas += tx.receipt.cumulativeGasUsed;
    }

    records.push({ no_nodes: no_nodes, gas: gas });
    no_nodes++;
  }

  await csvWriter.writeRecords(records).then(() => {
    console.log("...Done");
  });
};
