# Kubernetes Dashboard

The Dashboard UI is not deployed by default. To deploy it, run the following command:

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
```

You can access Dashboard using the kubectl command-line tool by running the following command:

```bash
kubectl proxy
```

run `open` or browse to url

```bash
open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
```

## Creating a User

To access the dashboard you will need to create a user and retrieve a token 

Use the yaml below to create a file `dashboard/dashboard-adminuser.yaml`

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
```

Then run the following command:

```bash
kubectl apply -f dashboard-adminuser.yaml
```

## Retrieving a token

Use the command below to retrieve a token for the user that you created above. Copy the token to use for logging into the dashboard:

```bash
kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}" 
```

## Reference

See <https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/>