# Why Kubernetes

This guide is not a how-to on Kubernetes as that is a much longer discussion. But quickly I want to cover the reasons
why Slalom consultants should consider Kubernetes if they are not already using it. That said, here is a quick Kubernetes
101:

[Kubernetes 101](k8s-101.md)

### Cloud Agnostic
Every major cloud provider now provides a managed Kubernetes offering. Thus you can easily create a Kubernetes cluster and deploy your
application to multiple cloud providers. 

**That's cool, but...** When do I need to deploy to multiple cloud providers? Probably never. But if a client is using 
(or open to using) Kubernetes and you know Kubernetes, then you know how to deploy your application regardless of cloud provider.

#### Velocity
Deploying a new version of an application or updating an application running in Kubernetes is fast and easy.

#### Easily Assemble a Complete Stack
Use Helm Charts (and other tools) to deploy databases, middleware, and other components for your stack quick and easily. 

#### Scaling
Efficiently scale horizontally using metrics (cpu, memory, etc). Soon, vertical auto-scaling.

#### Repeatable
Because everything is configured in yaml, it is easy to repeat. E.g. create multiple environments

![kubernetes](images/kubernetes.png)





