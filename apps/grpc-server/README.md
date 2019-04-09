
# grpc-server

====

gRPC mock server

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
kubectl get service grpc-server
> NAME          TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)     AGE
> grpc-server   ClusterIP   10.106.35.28   <none>        50051/TCP   17h

kubectl get deploy
> NAME          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
> grpc-server   2         2         2            2           17h

kubectl get po
> NAME                          READY     STATUS    RESTARTS   AGE
> grpc-server-f98c78d44-cxqch   1/1       Running   0          16h
> grpc-server-f98c78d44-gz626   1/1       Running   0          16h
```

## Usage

forward a local port to a port on the cluster ip

```shell
kubectl port-forward service/grpc-server 50051:50051
```

check ([grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe))

```shell
grpc-health-probe -addr=localhost:50051
> status: SERVING
```

crud sample ([grpcurl](https://github.com/fullstorydev/grpcurl))

```shell
# save
grpcurl -plaintext -d @ localhost:50051 proto.Todo.Save <<EOM
{
  "title": "TitleA"
}
EOM
# get
grpcurl -plaintext -d @ localhost:50051 proto.Todo.Get <<EOM
{
  "id": "xxx"
}
EOM
# list
grpcurl -plaintext localhost:50051 proto.Todo.List
# udpate
grpcurl -plaintext -d @ localhost:50051 proto.Todo.Update <<EOM
{
  "id": "xxx",
  "title": "title X-2"
}
EOM
# delete
grpcurl -plaintext -d @ localhost:50051 proto.Todo.Delete <<EOM
{
  "id": "xxx"
}
EOM
```
