#!/bin/sh
echo "*** build test ***"
make build
echo "*** unit test ***"
make test
echo "*** e2e test ***"
make e2e_test
echo "*** lint dockerfile ***"
make lint-dockerfile
