
# bff-grpc

====

backends for frontends (with gRPC)

## Requirement

- [Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
- kubectl(client:v1.11.9, server:v1.10.11
- [kubectx](https://github.com/ahmetb/kubectx)
- helm(client:v2.12.3, server:v2.12.3)
- go v1.12.0

## Install

deploy to docker-for-desktop

```shell
make docker-build
make helm-apply
```

deploy check

```shell
kubectl get service bff-grpc

kubectl get deploy bff-grpc

kubectl get po | grep bff-grpc
```

## Usage

forward a local port to a port on the cluster ip

```shell
kubectl port-forward service/bff-grpc 1323:1323
```

```shell
# save
# get
# list
# udpate
# delete
```
