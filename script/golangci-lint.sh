#!/bin/sh
for i in internal/ngsicmd internal/ngsilib e2e/ngsi-test e2e/server/accumulator e2e/server/atcontext e2e/server/csource e2e/server/oauth
do
  cd "$i" || exit 1
  golangci-lint run || exit 1
  cd - || exit 1
done
