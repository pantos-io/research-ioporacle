const fs = require("fs");
const path = require("path");

const registryContractPath = path.resolve(
  __dirname,
  "../build/contracts/RegistryContract.json"
);
const oracleContractPath = path.resolve(
  __dirname,
  "../build/contracts/OracleContract.json"
);
const distKeyContractPath = path.resolve(
  __dirname,
  "../build/contracts/DistKeyContract.json"
);

const abiPath = path.resolve(__dirname, "../build/contracts/abi");
const binPath = path.resolve(__dirname, "../build/contracts/bin");

const registryContractAbiPath = path.resolve(abiPath, "RegistryContract.abi");
const oracleContractAbiPath = path.resolve(abiPath, "OracleContract.abi");
const distKeyContractAbiPath = path.resolve(abiPath, "DistKeyContract.abi");

const registryContractBinPath = path.resolve(binPath, "RegistryContract.bin");
const oracleContractBinPath = path.resolve(binPath, "OracleContract.bin");
const distKeyContractBinPath = path.resolve(binPath, "DistKeyContract.bin");

const registryContract = JSON.parse(
  fs.readFileSync(registryContractPath, "utf8")
);

const oracleContract = JSON.parse(fs.readFileSync(oracleContractPath, "utf8"));

const distKeyContract = JSON.parse(
  fs.readFileSync(distKeyContractPath, "utf8")
);

fs.mkdirSync(abiPath, { recursive: true });
fs.mkdirSync(binPath, { recursive: true });

fs.writeFile(
  registryContractAbiPath,
  JSON.stringify(registryContract.abi),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("RegistryContract ABI written successfully!");
  }
);

fs.writeFile(
  registryContractBinPath,
  JSON.stringify(registryContract.bytecode),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("RegistryContract BIN written successfully!");
  }
);

fs.writeFile(
  oracleContractAbiPath,
  JSON.stringify(oracleContract.abi),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("OracleContract ABI written successfully!");
  }
);

fs.writeFile(
  oracleContractBinPath,
  JSON.stringify(oracleContract.bytecode),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("OracleContract BIN written successfully!");
  }
);

fs.writeFile(
  distKeyContractAbiPath,
  JSON.stringify(distKeyContract.abi),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("DistKeyContract ABI written successfully!");
  }
);

fs.writeFile(
  distKeyContractBinPath,
  JSON.stringify(distKeyContract.bytecode),
  function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("DistKeyContract BIN written successfully!");
  }
);
