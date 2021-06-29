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
$ helm upgrade --install redev -f values/redis.yaml bitnami/redis --version 14.6.2
```

