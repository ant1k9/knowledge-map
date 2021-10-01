[grpcurl](https://github.com/fullstorydev/grpcurl)

*Description*: grpcurl is a command-line tool that lets you interact with gRPC servers. It's basically curl for gRPC servers.

*Labels*: #Go #GRPC

*Examples*:

```bash
$ grpcurl grpc.server.com:443 list
$ grpcurl grpc.server.com:443 describe server.Service
$ grpcurl -H 'Authentication: Bearer <token>' \
    -d '{"foo": "bar"}' \
    grpc.server.com:443 server.Service.Method
```
