#!/bin/sh

if ! which abigen >/dev/null; then
  echo "error: abigen not installed" >&2
  exit 1
fi

abigen --sol ../../ioporaclecontracts/contracts/IopOracle.sol --pkg iop --out ../pkg/iop/ioporaclecontract.go