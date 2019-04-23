# bff-server

====

backends for frontends

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
kubectl get service bff-server
> NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
> bff-server   ClusterIP   10.103.180.57   <none>        1323/TCP   3d

kubectl get deploy bff-server
> NAME         DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
> bff-server   2         2         2            2           3d

kubectl get po -l app=bff-server
> NAME                          READY     STATUS    RESTARTS   AGE
> bff-server-7cd858d4cf-4qrf7   1/1       Running   0          11m
> bff-server-7cd858d4cf-xrkb2   1/1       Running   0          12m
```

## Usage

forward a local port to a port on the cluster ip

```shell
kubectl port-forward service/bff-server 1323:1323
```

````shell
# save
```shell
ID=$(curl -v -XPOST -H 'Content-Type: application/json' -d '{"title": "TitleA"}' http://localhost:1323/todos | jq -r '.id')
````

# get

```shell
curl -v "http://localhost:1323/todos/${ID}"
```

# list

```shell
curl -v http://localhost:1323/todos
```

# update

```shell
curl -XPATCH -v -H 'Content-Type: application/json'  -d '{"title": "TitleA-update"}'  "http://localhost:1323/todos/${ID}"
```

# delete

```shell
curl -v -XDELETE "http://localhost:1323/todos/${ID}"
```
