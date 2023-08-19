#!/bin/sh
set -ue
echo "{"serviceMappings":[]}" > ./config/cygnus-name_mappings.conf
docker compose up -d
