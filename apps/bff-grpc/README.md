
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
make
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
```shell
curl -v -XPOST -d @ http://localhost:1323/todos <<EOM
{
  "id": "1",
  "title": "TitleA"
}
EOM
```

# get
```shell
curl -v http://localhost:1323/todos/id
```

# list

```shell
curl -v http://localhost:1323/todos
```

# update
```shell
curl -v -XPATCH -d @ http://localhost:1323/todos/1 <<EOM
{
  "title": "Titlea"
}
EOM
```

# delete
```shell
curl -v -XDELETE -d @ http://localhost:1323/todos/1
```

