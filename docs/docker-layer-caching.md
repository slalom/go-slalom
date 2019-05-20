# Docker Layer Caching for Faster Builds

The build script [docker-azure-build.sh](../ci/docker-azure-build.sh) uses docker `--cache-from` and `--target` arguments with multi-stage builds. This process speeds up builds by caching dependencies in a separate image and using it as a build cache. This is in `go` but the same could be applied for any other language. Below is how it works.

## Mulit-Stage Docker Build

There are three stages in the [Dockerfile](../Dockerfile)

- builder
  - uses `go mod download` to download all required dependencies
- compile
  - compiles source and generates binary using `builder` layer for dependencies
- go_slalom
  - creates the final image from `scatch` and copies the binary from `compile` layer

## Using --target and --cache-from to cache dependencies

The `--target` argument allows building a specific stage in a multi-stage Dockerfile. This is beneficial
because it allows us to build the `builder` stage and push it as a separate image so that it can later
be used as a cache.

The `--cache-from` tells Docker to use a specific image as part of it's build cache. If the provided image and your current build have layers in common, you get the same speed up as if the image had been built on the same machine.

## Putting it all together

The script [ci/docker-azure-build.sh](../ci/docker-azure-build.sh) performs the following steps:

1. Pull the `builder` image to use as a build cache
2. Build the `builder` image using pulled image from step 1 as a build cache. If the layers (i.e. dependencies) have not changed in the new build, then image built in this step will come entirely from cache. Otherwise depenencies will be updated in the image.
3. Build the final image using image from step 2 as build cache.
4. Push the runtime image (go-slalom) to registry
5. Push the `builder` image from step 2 so that it can be used in future builds

If dependencies have not changed then `builder` image is pulled from registry and the final image uses it as a build cache. If dependencies have changed then the `builder` image is updated and pushed at the end to be used as a cache for future builds.
