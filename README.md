# k8sでBFF 環境の構築

## 目的

今までの総まとめとして BFF 環境の Mock を作成する

- Backends For Frontends が HttpClient で実装できる場合

![bff-http](https://raw.githubusercontent.com/taguch1/try-bff/images/readme/bff-http.png?token=AEy1_j9jIPsrv6bzAq6NewdgcdcE3Yqwks5cqr1RwA%3D%3D)

```mermaid
sequenceDiagram
    frondends(js)->>LB(nginx):https
    LB(nginx)->>BFF:http
    BFF->>APIServer:http(REST)
    APIServer->>AttchedResourece:tcp(mysql)
```


- Backends For Frontends が GRPClient で実装できる場合

![bff-grpc](https://raw.githubusercontent.com/taguch1/try-bff/images/readme/bff-grpc.png?token=AEy1_nlXpcXPVQ3rGXdq5iF2W14FCzteks5cqrz3wA%3D%3D)

```mermaid
sequenceDiagram
    frondends(js)->>LB(nginx):https
    LB(nginx)->>BFF:http
    BFF->>APIServer:GRPC
    APIServer->>AttchedResourece:tcp(mysql)
```

- Backends For Frontends が 両方混在する場合

![bff-mixed](https://raw.githubusercontent.com/taguch1/try-bff/images/readme/bff-mixed.png?token=AEy1_t0e8-X3qJ1Cho2GnrGC4Vlmzt92ks5cqr0cwA%3D%3D)

```mermaid
sequenceDiagram
    frondends(js)->>LB(nginx):https
    LB(nginx)->>BFF:http
    BFF->>GRPCGateWay:REST
    GRPCGateWay->APIServer:GRPC
    BFF->>APIServer:GRPC
    APIServer->>AttchedResourece:tcp(mysql)
```

## タスク

- 各モジュール作成
- Docker イメージ
- Helm
- CI&CD
- 監視
- バッチ処理


## 作業内容概要

### 各モジュール作成

- frondend
- bff-grpc
- backend-api
- bff-http
- grpc-gateway

### Docker イメージ

- nginx
- proxysql
- mysql

### Helm

- nginx
- proxysql
- mysql

### CI(circleci)

各モジュールのunit testとdocker imageのbuildまで

### CD(予定は未定)

CircleCIがデブロイ先(docker for mac)まで届かないので省略

### 監視(予定は未定)

- prometeus
- graphana

### バッチ処理(予定は未定)

- CronJob
- k8sAPIを使ったシンプルなJobManager(シーケンシャルにJobを実行するやつ)
