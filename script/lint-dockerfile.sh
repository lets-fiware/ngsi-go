#!/bin/sh
docker pull hadolint/hadolint
docker run --rm -i hadolint/hadolint < docker/Dockerfile
docker run --rm -i hadolint/hadolint < e2e/ngsi-test/Dockerfile
docker run --rm -i hadolint/hadolint < e2e/server/accumulator/Dockerfile
docker run --rm -i hadolint/hadolint < e2e/server/atcontext/Dockerfile
docker run --rm -i hadolint/hadolint < e2e/server/csource/Dockerfile
docker run --rm -i hadolint/hadolint < e2e/server/oauth/Dockerfile
