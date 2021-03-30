const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const RegistryContract = artifacts.require("RegistryContract");
const DistKeyContract = artifacts.require("DistKeyContract");
const ECDSAOracleContract = artifacts.require("ECDSAOracleContract");

module.exports = async function (callback) {
  const csvWriter = createCsvWriter({
    path: "./data/validate-submit-tx-ecdsa.csv",
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
    let ecdsaOracleContract = await ECDSAOracleContract.new(
      registryContract.address
    );
    await distKeyContract.setRegistryContract(registryContract.address);

    let accounts = await web3.eth.getAccounts();
    for (let i = 0; i < no_nodes; i++) {
      await registryContract.registerOracleNode("127.0.0.1", "0x0", {
        from: accounts[i],
      });
    }

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

    let aggregator = await registryContract.getAggregator();
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
