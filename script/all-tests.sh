#!/bin/sh
echo "*** build test ***"
make build || exit 1
echo "*** unit test ***"
make unit-test || exit 1
echo "*** e2e test ***"
make e2e-test || exit 1
echo "*** golangci-lint"
make golangci-lint || exit 1
echo "*** lint dockerfile ***"
make lint-dockerfile || exit 1
