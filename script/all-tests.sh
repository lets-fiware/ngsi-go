#!/bin/sh
echo "*** build test ***"
make build || exit 1
echo "*** unit test ***"
make test || exit 1
echo "*** e2e test ***"
make e2e_test || exit 1
echo "*** lint dockerfile ***"
make lint-dockerfile || exit 1
