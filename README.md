### REST Application demonstrating GRPC usage in Go

#### Folder structure
 - /cmd: contains local app main function
 - /rpc: contains /service which is the implementation of the generated protobuf server interface.
  And /client containing a client-connection helper
 - /handler: contains http.Handlers that invoke the remote grpc functions via the client-connection
 - /proto contains the rpc .proto file and generated interfaces

#### Usage:

The RPC server exposes a simple user management API
#### Start rpc server
```shell
go run ./rpc/service
```
#### Start app server
```shell
 go run ./cmd/app.go
```

 #### Create user
```shell
curl --location --request POST 'localhost:8081/api/user/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Name":"Adam",
    "Age":10
}'
```

#### Fetch all users
```shell
curl --location --request GET 'localhost:8081/api/users'
```

#### Fetch single user by ID
```shell
curl --location --request GET 'localhost:8081/api/user?user_id=8f37caae-a76d-45b8-8cb4-be79c21f6ef1'
```

#### To regenerate rpc code from .proto file (you must have grpc setup locally)
```shell
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     ./proto/def.proto

```