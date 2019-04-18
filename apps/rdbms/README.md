
# rdbms

====

mysql

## Requirement

- kubectl(client:v1.11.9, server:v1.10.11
- [kubectx](https://github.com/ahmetb/kubectx)
- helm(client:v2.12.3, server:v2.12.3)

## Install

deploy to docker-for-desktop

```shell
make
```

deploy check

```shell
kubectl get service rdbms
> NAME      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
> rdbms     ClusterIP   10.105.186.42   <none>        3306/TCP   16s

kubectl get deploy rdbms
> NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
> rdbms     1         1         1            1           40s

kubectl get po
> NAME                           READY     STATUS    RESTARTS   AGE
> rdbms-698d97d666-6vkkr         1/1       Running   0          56s
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
