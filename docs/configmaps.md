# Configmaps

Configmaps are a simple, but vital object in kubernetes. Most applications will use configmaps for configuration

## go-slalom configmap

The go-slalom configmap is configured in [deploy/skaffold/configmap.yaml](../deploy/skaffold/configmap.yaml)

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: go-slalom
data:
  # property-like keys; each key maps to a simple value
  magic_value: "TODO"
```

We can also see it using kubectl commands

```bash
# list configmaps
kubectl get configmaps

# show go-slalom configmap
kubectl get configmap go-slalom -o yaml
```

The "magic_value" is mapped as environment variable "MAGIC_VALUE" in [deploy/skaffold/deployment.yaml](../deploy/skaffold/deployment.yaml)

```yaml
env:  
  - name: MAGIC_VALUE
    valueFrom:
      configMapKeyRef:
        name: go-slalom
        key: magic_value
```

The environment variable value is used in the output of the service

## Editing Configmaps

You can see there is only 1 field, magic_value, that is set to "TODO". Let's change it

```bash
kubectl edit configmap go-slalom
```

Now curl go-slalom again...

What! It is still showing "TODO" for "magic_value". That is expected... the pod needs restarted so it can receive the new value.

We could restart skaffold, or another way is to disable the healthcheck and let kubernetes restart the pod:

```bash
curl -X POST localhost:8080/health/disable
```

After a moment you will see kubernetes sends a signal to tell the pod to shutdown. The pod handles the signal and exits. Kubernetes starts a new pod.

Now if we curl the service again we will see the updated configmap value in the result

## Next

[Release go-slalom](go-releaser.md)

![slalom-gopher](images/gopher-mic-drop.png)