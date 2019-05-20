#!/usr/bin/env bash

DOCKER_IMAGE_VERSION=${1:-latest}

# go to repo root
repo_root="$(cd $(dirname ${BASH_SOURCE[0]})/.. && pwd)"
cd "${repo_root}"

# pull the latest builder image
docker pull goslalom.azurecr.io/go-slalom:builder || true

# only build the 'builder' stage, using pulled image as cache
docker build \
  --target builder \
  --cache-from goslalom.azurecr.io/go-slalom:builder \
  -t goslalom.azurecr.io/go-slalom:builder \
  "."

# build the final image using builder cache
docker build \
  --cache-from goslalom.azurecr.io/go-slalom:builder \
  -t goslalom.azurecr.io/go-slalom:$DOCKER_IMAGE_VERSION \
  "."

# push runtime images to remote repository
docker push goslalom.azurecr.io/go-slalom:$DOCKER_IMAGE_VERSION

# push the builder image for next build
docker push goslalom.azurecr.io/go-slalom:builder