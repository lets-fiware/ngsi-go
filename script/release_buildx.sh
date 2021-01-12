#!/bin/sh
NGSIGO_TARGET=$1
if [ -z "$NGSIGO_TARGET" ]; then
  NGSIGO_TARGET=all
fi
export NGSIGO_TARGET
cd docker || exit 1
make build || exit 1
make run || exit 1
