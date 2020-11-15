#!/bin/sh
NGSIGO_TARGET=$1
if [ -z "$NGSIGO_TARGET" ]; then
  NGSIGO_TARGET=all
fi
export NGSIGO_TARGET
cd docker
make build
make run
