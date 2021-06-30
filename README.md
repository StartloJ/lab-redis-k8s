# Demo lab PoC redis

## Provision local kubernetes cluster
Prerequisite
- Go cmd in version 1.16.x
- Helm version 3.x

Download kind package
```bash
$ go get sigs.k8s.io/kind
```

Create kind cluster with configuration
```bash
$ kind create cluster --config=kind.yaml
```

Fetch repository bitnami
```bash
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm repo update
```

Deploy redis cluster with HELM
```bash
$ kubectl create namespace cache
$ helm upgrade --namespace cache --install redev -f values/redis.yaml bitnami/redis --version 14.6.2
```

## Use demo application for proof redis
1. Use for container, you can pull images on 
```bash
$ docker pull dukecyber/demo-redev:0.1.0
```

2. Use app in kubernetes
```bash
$ kubectl create namespace demo
$ kubectl apply -f values/deployment.yaml --namespace demo
```
