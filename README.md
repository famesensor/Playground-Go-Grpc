# Playground-Go-Grpc
_This project for self-learning client connect service golang by grpc_
#### GRPC connect type...
- Unary RPC
- Client streaming RPC
- Server streaming RPC
- Bidirectional streaming RPC

## Installation
```
# Go plugins for the protocol compiler
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Run Application
```
# Unary
proto : make gen-proto-unary
client : go run ./unary/clinet/client.go
service  : go run ./unary/server/server.go

# Client stream
proto : make gen-proto-client-stream
client : go run ./grpc_call/client_stream/client/client.go
service : go run ./grpc_call/client_stream/server/server.go

# Server stream
proto : make gen-proto-server-stream
client : go run ./grpc_call/server_stream/client/client.go
service : go run ./grpc_call/server_stream/server/server.go

# Bidirectional stream
proto : make gen-proto-bi-direction-stream
client : go run ./grpc_call/bi_directional_stream/client/client.go
service : go run ./grpc_call/bi_directional_stream/server/server.go
```
