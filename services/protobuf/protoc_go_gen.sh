#!/bin/bash
GEN_FOLDER=_go
if [[ ! -d $GEN_FOLDER ]]; then
  mkdir $GEN_FOLDER
fi
protoc --go_out=$GEN_FOLDER \
  --go_opt=paths=source_relative \
  --go-grpc_out=$GEN_FOLDER \
  --go-grpc_opt=paths=source_relative ./**/*.proto
