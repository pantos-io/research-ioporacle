#!/bin/sh

if ! which protoc >/dev/null; then
  echo "error: protoc not installed" >&2
  exit 1
fi

protoc -I ../api/proto --go_out=../pkg/iop --go-grpc_out=../pkg/iop ../api/proto/ioporaclenode.proto