const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const RegistryContract = artifacts.require("RegistryContract");
const DistKeyContract = artifacts.require("DistKeyContract");
const ECDSAOracleContract = artifacts.require("ECDSAOracleContract");

module.exports = async function () {
  let dir = "./data";
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir);
  }
  const csvWriter = createCsvWriter({
    path: "./data/ecdsa-cost.csv",
    header: [
      { id: "no_nodes", title: "no_nodes" },
      { id: "gas", title: "gas" },
    ],
  });
  let records = [];

  let no_nodes = 3;

  while (no_nodes < 64) {
    console.log(no_nodes);
    let distKeyContract = await DistKeyContract.new();
    let registryContract = await RegistryContract.new(distKeyContract.address);
    let ecdsaOracleContract = await ECDSAOracleContract.new(
      registryContract.address
    );
    await distKeyContract.setRegistryContract(registryContract.address);

    let stake = await registryContract.MIN_STAKE();
    let accounts = await web3.eth.getAccounts();
    for (let i = 0; i < no_nodes; i++) {
      await registryContract.registerOracleNode("127.0.0.1", "0x0", {
        from: accounts[i],
        value: stake,
      });
    }

    let fee = await ecdsaOracleContract.FEE();
    let fees = no_nodes * fee;
    await ecdsaOracleContract.validateTransaction(
      "0xa67220981e1760824947fb294f65adf47c505c3bfbe5960341d64c7f7512be8a",
      3,
      {
        value: fees,
      }
    );

    let result = await ecdsaOracleContract.encodeResult(0, true);
    let hash = web3.utils.keccak256(result);

    let signatures = [];
    let majority = Math.trunc(no_nodes / 2) + 1;

    for (let i = 0; i < majority; i++) {
      let sig = await web3.eth.sign(hash, accounts[i]);
      sig = sig.split("x")[1];
      let v = parseInt(sig.substring(128, 130)) + 27;
      signatures.push(
        "0x" + sig.substring(0, 128) + web3.utils.numberToHex(v).substring(2)
      );
    }

    let block = (await web3.eth.getBlockNumber()) + 1;
    let aggregator = await registryContract.getAggregatorByBlock(block);
    let tx = await ecdsaOracleContract.submitValidationResult(
      0,
      true,
      signatures,
      { from: aggregator.addr }
    );

    records.push({ no_nodes: no_nodes, gas: tx.receipt.cumulativeGasUsed });
    no_nodes++;
  }

  await csvWriter.writeRecords(records).then(() => {
    console.log("...Done");
  });
};
