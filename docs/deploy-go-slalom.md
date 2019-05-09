# Deploy go-slalom

go-slalom demonstrates how to create a basic service to run in kubernetes. The service includes the following:

- handles interrupts expected from kubernetes
- basic liveness and readiness end-points
- endpoints for enabling/disabling readiness and liveness endpoints
- endpoint for basic service information

### run go-slalom

```bash
# tell go to use modules (still opt-in feature)
export GO111MODULE=on 

# install go-slalom
go get github.com/tredfield/go-slalom
```

### build go-slalom locally

```bash
git clone git@github.com:tredfield/go-slalom.git
```

### go-slalom cli

go-slalom includes a basic cli built using [cobra](https://github.com/spf13/cobra). It is common in go to build a cli
for starting a service and/or administering the service.

```bash
# run to see cli help message
./go-slalom help
```

