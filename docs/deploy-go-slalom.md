# Deploy go-slalom

go-slalom demonstrates how to create a basic service to run in kubernetes. The service includes the following:

- handles interrupts expected from kubernetes
- basic liveness and readiness end-points
- endpoints for enabling/disabling readiness and liveness endpoints
- endpoint for basic service information

### go-slalom cli

go-slalom includes a basic cli built using [cobra](https://github.com/spf13/cobra). It is common in go to build a cli
for starting a service and/or administering the service.

```bash
./go-slalom help
```