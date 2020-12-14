const fs = require("fs");
const oracleContract = JSON.parse(
  fs.readFileSync("../build/contracts/OracleContract.json", "utf8")
);

const registryContract = JSON.parse(
  fs.readFileSync("../build/contracts/RegistryContract.json", "utf8")
);

fs.mkdir("../build/contracts/abi", { recursive: true }, err => {
  if (err) throw err;
});

fs.mkdir("../build/contracts/bin", { recursive: true }, err => {
  if (err) throw err;
});

fs.writeFile(
  "../build/contracts/abi/OracleContract.abi",
  JSON.stringify(oracleContract.abi),
  function(err) {
    if (err) {
      return console.error(err);
    }
    console.log("OracleContract ABI written successfully!");
  }
);

fs.writeFile(
  "../build/contracts/bin/OracleContract.bin",
  JSON.stringify(oracleContract.bytecode),
  function(err) {
    if (err) {
      return console.error(err);
    }
    console.log("OracleContract BIN written successfully!");
  }
);

fs.writeFile(
  "../build/contracts/abi/RegistryContract.abi",
  JSON.stringify(registryContract.abi),
  function(err) {
    if (err) {
      return console.error(err);
    }
    console.log("RegistryContract ABI written successfully!");
  }
);

fs.writeFile(
  "../build/contracts/bin/RegistryContract.bin",
  JSON.stringify(registryContract.bytecode),
  function(err) {
    if (err) {
      return console.error(err);
    }
    console.log("RegistryContract BIN written successfully!");
  }
);
