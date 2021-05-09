run-service:
	go run ./server/server.go

run-client:
	go run ./client/client.go

gen-proto-unary:
	protoc --proto_path=./unary/proto --go_out=plugins=grpc:./unary/proto service.proto

	# protoc --proto_path=./proto --go_out=plugins=grpc:. --go_opt=paths=source_relative service.proto

gen-proto-client-stream:
	protoc --proto_path=./grpc_call/client_stream/proto --go_out=plugins=grpc:./grpc_call/client_stream/proto api.proto

