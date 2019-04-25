# web-api

====

nginx reverse proxy

## Requirement

- kubectl(client:v1.11.9, server:v1.10.11
- helm(client:v2.12.3, server:v2.12.3)

## Install

deploy to docker-for-desktop

```shell
make
```

deploy check

```shell
kubectl get service web-api
> NAME      TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
> web-api   ClusterIP   10.96.39.186   <none>        80/TCP    3d

kubectl get deploy web-api
> NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
> web-api   2         2         2            2           3d

kubectl get po -l app=web-api
> NAME                       READY     STATUS    RESTARTS   AGE
> web-api-7f69498ffd-4dwjf   1/1       Running   0          6m
> web-api-7f69498ffd-9kvrq   1/1       Running   0          6m```
```

## Usage

forward a local port to a port on the cluster ip

```shell
make forward
curl  localhost:8081/todos
```
