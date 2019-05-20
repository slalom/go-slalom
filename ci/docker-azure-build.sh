#!/usr/bin/env bash

DOCKER_IMAGE_VERSION=${1:-latest}

# go to repo root
repo_root="$(cd $(dirname ${BASH_SOURCE[0]})/.. && pwd)"
cd "${repo_root}"

# Step 1: pull the latest builder image
docker pull goslalom.azurecr.io/go-slalom:builder || true

# Step 2: build the 'builder' stage, using pulled image as cache. If there are no changes to dependencies
# this step will all come from cache.
docker build \
  --target builder \
  --cache-from goslalom.azurecr.io/go-slalom:builder \
  -t goslalom.azurecr.io/go-slalom:builder \
  "."

# Step 3: build the final image using builder cache. This step will create compile layer using cached
# dependencies from builder
docker build \
  --cache-from goslalom.azurecr.io/go-slalom:builder \
  -t goslalom.azurecr.io/go-slalom:$DOCKER_IMAGE_VERSION \
  "."

# Step 4: push runtime images to remote repository
docker push goslalom.azurecr.io/go-slalom:$DOCKER_IMAGE_VERSION

# Step 5: push the builder image for next build
docker push goslalom.azurecr.io/go-slalom:builder