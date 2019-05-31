# Base build image
FROM golang:1.12 AS builder

WORKDIR /go/src/go-slalom

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./

# This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the `go mod` download
# command will be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download


# This image builds the server
FROM builder AS compile
# Here we copy the rest of the source code
COPY . .

# And compile the project
RUN GIT_TAG=$(git describe --tags --always) GIT_COMMIT=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux go install -ldflags="-X 'go-slalom/pkg/version.Version=$GIT_TAG' -X 'go-slalom/pkg/version.Revision=$GIT_COMMIT'"

# In this last stage, we start from 'scratch' to reduce the image size and not ship the Go compiler in our production artifacts.
FROM scratch AS go_slalom

# Finally we copy the statically compiled Go binary.
COPY --from=compile /go/src/go-slalom/templates ./templates
COPY --from=compile /go/src/go-slalom/static ./static
COPY --from=compile /go/bin/go-slalom /bin/go-slalom
ENTRYPOINT ["/bin/go-slalom"]