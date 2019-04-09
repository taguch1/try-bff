## run server

```shell
make protoc
make release
./bin/grpc-server_darwin_amd64
```

## check ([grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe))

```shell
grpc-health-probe -addr=localhost:50051
> status: SERVING
```

## check ([grpcurl](https://github.com/fullstorydev/grpcurl))

health-check

```shell
grpcurl -plaintext localhost:50051 list
grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection

grpcurl -plaintext localhost:50051  grpc.health.v1.Health/Check
{
  "status": "SERVING"

}
```

todo crud

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
