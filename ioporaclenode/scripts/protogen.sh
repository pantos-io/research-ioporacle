#!/bin/sh

cd "$(dirname "$0")" || exit 1

if ! which protoc >/dev/null; then
  echo "error: protoc not installed" >&2
  exit 1
fi

protoc -I ../api/proto/iop --go_out=../pkg/iop --go-grpc_out=../pkg/iop ../api/proto/iop/oraclenode.proto
