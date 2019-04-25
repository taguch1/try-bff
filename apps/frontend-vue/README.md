# frontend-vue

====

vue app

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
kubectl get service frontend-vue
> NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
> frontend-vue   ClusterIP   10.110.47.187   <none>        80/TCP    57m

kubectl get deploy frontend-vue
> NAME           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
> frontend-vue   2         2         2            2           57m

kubectl get po -l app=frontend-vue
> NAME                            READY     STATUS    RESTARTS   AGE
> frontend-vue-68d8f9b96b-27cjf   1/1       Running   0          52m
> frontend-vue-68d8f9b96b-hrvkd   1/1       Running   0          52m```

## Usage

forward a local port to a port on the cluster ip

```shell
make forward
open localhost:8080
```
