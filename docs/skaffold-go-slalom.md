# deploy go-slalom to kubernetes

We will use [skaffold](https://skaffold.dev) to deploy go-slalom

Skaffold is a tool that will build a docker image and deploy it to kubernetes. It can run in `dev` mode where it listens to changes and will automatically rebuild and update the deployed version in kubernetes. Or you can use `run` to do one time deployment. It is a great tool for development but can also be applied in your ci/cd pipelines.

## install skaffold

```bash
brew install skaffold
```

## check kubernetes is running

First kill go-slalom if it is still running using `ctrl-c`. It should log showing it has handled the interrupt it received.

```bash
^C{"level":"info","msg":"received interrupt","pkg":"signals","time":"2019-05-09T13:27:32-07:00"}
{"level":"warning","msg":"stopping server","service":"api","time":"2019-05-09T13:27:32-07:00"}
```

Make sure `kubernetes` is running in docker desktop.

```bash
kubectl config current-context
```

You should see output `docker-for-desktop`

Now run

```bash
kubectl cluster-info
```

You should see the following output

```bash
Kubernetes master is running at https://localhost:6443
KubeDNS is running at https://localhost:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

If all above is good then kubernetes is running!

## run skaffold

Ok, now let's deploy go-slalom. Run

```bash
skaffold dev --port-forward=pods
```

skaffold will do the following

- generates a tag
- builds the docker image
- deploy to kubernetes (using [deploy/skaffold/deployment.yml](../deploy/skaffold/deployment.yaml))
- port forward to go-slalom container
- wait and watch for changes

### verify go-slalom is running

Check kubernetes

```bash
kubectl get pods
```

You should see output similar to

```bash
NAME                         READY   STATUS    RESTARTS   AGE
go-slalom-5dbc56689c-jfsd8   1/1     Running   0          4m
```

curl go-slalom

```bash
curl localhost:8080
```

You should see similar output below and logging in console where skaffold is running.

```bash
{
  "hostname": "TODO",
  "version": "0.0.1",
  "revision": "unknown",
  "color": "",
  "message": "",
  "goos": "darwin",
  "goarch": "amd64",
  "runtime": "go1.12.4",
  "num_goroutine": "6",
  "num_cpu": "8",
  "magic_value": "TODO"
}
```

### deploy on change

Note in the output above the `hostname` is `TODO`. Lets fix that.

Add the following change at line 14 to `pkg/api/info.go`

```go
host, err := os.Hostname()
if err != nil {
    logger.Errorf("can not retrieve host")
}
```

Find `"TODO"` and replace with `host` (no quotes). Save the file.

You should see skaffold rebuild the image and update the deployment in kubernetes. Generating the docker image should be
much faster this time since dependencies have been downloaded and cached in a docker layer.

Skaffold will again port forward the container. Note the port and use to test

```bash
# use port from skaffold
curl localhost:8080
```

You should now see the hostname matches the pod it is running in.

#### How is skaffold managing the deployment?

When skaffold detects the change it is doing the same steps listed above. The thing to note is that the image configured in
[deploy/skaffold/deployment.yml](../deploy/skaffold/deployment.yaml) does not designate the tag. Skaffold manages the tag
using the tag from the last image it generated. There are different strategies that can be used with skaffold for managing
the tag.

## Next

[configmaps](configmaps.md)

![slalom-gopher](images/gopherski.png)