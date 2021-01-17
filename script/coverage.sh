#!/bin/sh
cd internal || exit 1
go test -coverprofile=../coverage/coverage.out -covermode=atomic ./... || exit 1
cd .. || exit 1
go tool cover -html=coverage/coverage.out -o coverage/coverage.html
