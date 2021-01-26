#!/bin/sh

cd "$(dirname "$0")" || exit 1

if ! which abigen >/dev/null; then
  echo "error: abigen not installed" >&2
  exit 1
fi

abigen --abi ../../ioporaclecontracts/build/contracts/abi/RegistryContract.abi --pkg iop --type RegistryContract --out ../pkg/iop/registrycontract.go
abigen --abi ../../ioporaclecontracts/build/contracts/abi/OracleContract.abi --pkg iop --type OracleContract --out ../pkg/iop/oraclecontract.go
abigen --abi ../../ioporaclecontracts/build/contracts/abi/RaffleContract.abi --pkg iop --type RaffleContract --out ../pkg/iop/rafflecontract.go
abigen --abi ../../ioporaclecontracts/build/contracts/abi/DistKeyContract.abi --pkg iop --type DistKeyContract --out ../pkg/iop/distkeycontract.go