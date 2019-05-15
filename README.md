# go-slalom

<https://github.com/slalom/go-slalom>

[![Build Status](https://dev.azure.com/tredfield/go-slalom/_apis/build/status/slalom.go-slalom?branchName=master)](https://dev.azure.com/tredfield/go-slalom/_build/latest?definitionId=2&branchName=master)

go-slalom is a tiny service made with [Go](https://golang.org) that showcases how to easily develop and deploy
Go services in [Kubernetes](https://kubernetes.io).

This repository and guides will help you

- Setup [Docker Desktop](https://www.docker.com/products/docker-desktop) with [Kubernetes](https://kubernetes.io)
- Setup [Homebrew](https://brew.sh) and install [go](https://golang.org)
- Use [corba](https://github.com/spf13/cobra) to create a CLI for your service
- Use [skaffold](https://skaffold.dev) to build and deploy your service
- Use [goreleaser](https://goreleaser.com) to release your cli/service

## Guides

- [Pre-requisite Instructions](docs/prereqs.md)
- [Why Go](docs/why-go.md)
- [Why Kubernetes](docs/why-kubernetes.md)
- [Hello World in Go](docs/go-hello-world.md)
- [Run go-slalom](docs/build-go-slalom.md)
- [Deploy go-slalom with skaffold](docs/skaffold-go-slalom.md)
- [Release go-slalom](docs/go-releaser.md)
- [What's Next?](docs/whats-next.md)

### Quick 101 on go and kubernetes

The above guides are an introduction to go and kubernetes but are not intended to teach you either. However, below are quick 101s
that should only take a few minutes with references to more help

- [go 101](docs/go-101.md)
- [kubernetes 101](docs/k8s-101.md)
  
![gopher](docs/images/go-n-docker.png)
