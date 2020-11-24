#!/bin/sh

cd "$(dirname "$0")" || exit 1

if ! which abigen >/dev/null; then
  echo "error: abigen not installed" >&2
  exit 1
fi

abigen --sol ../../ioporaclecontracts/contracts/OracleContract.sol --pkg iop --out ../pkg/iop/oraclecontract.go