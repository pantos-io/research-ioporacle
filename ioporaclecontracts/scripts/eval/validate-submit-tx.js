const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function (callback) {
  const csvWriter = createCsvWriter({
    path: "./data/validate-submit.csv",
    header: [
      { id: "id", title: "id" },
      { id: "gas", title: "gas" },
    ],
  });
  let records = [];

  let topic, eventInterface;
  let eventName = "SubmitValidationResultLog";
  for (const [key, value] of Object.entries(OracleContract.events)) {
    if (value.name === eventName) {
      topic = key;
      eventInterface = value;
    }
  }

  let counter = 0;
  let oracleContract = await OracleContract.deployed();
  let tx = "0xa67220981e1760824947fb294f65adf47c505c3bfbe5960341d64c7f7512be8a";

  await web3.eth.subscribe(
    "logs",
    {
      address: OracleContract.address,
      topics: [topic],
    },
    async function (error, result) {
      if (!error) {
        let receipt = await web3.eth.getTransactionReceipt(
          result.transactionHash
        );
        records.push({ id: counter, gas: receipt.cumulativeGasUsed });

        if (counter === 100) {
          await csvWriter.writeRecords(records).then(() => {
            console.log("...Done");
          });
          return;
        }

        await oracleContract.validateTransaction(tx, 3, {
          value: web3.utils.toWei("0.0011", "ether"),
        });
        counter++;
      }
    }
  );

  await oracleContract.validateTransaction(tx, 3, {
    value: web3.utils.toWei("0.0011", "ether"),
  });
};
