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

## Basic Redis values you can change
1. In line 20 `global.redis.password`: you can changed access redis with password. For empty string was mean no password to access.
2. In line 29 `nameOverride`: you can rename to use in kubernetes resouce with this value
3. In line 32 `fullnameOverride`: you can change HELM release name with this value
4. In line 343 `master.persistence.enabled`: If you do not need to keep data in volume, you can change this value to `false`.
5. In line 423 `replica.replicaCount`: you can incease number to scale readable node(also nodd for Read only)
6. In line 651 `replica.persistence.enabled`: If you do not need to keep data in volume, you can change this value to `false`.


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

