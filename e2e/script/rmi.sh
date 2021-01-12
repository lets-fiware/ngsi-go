#!/bin/sh
docker image prune --filter label=stage=ngsi-builder -f
docker image prune --filter label=stage=ngsi-runner -f
