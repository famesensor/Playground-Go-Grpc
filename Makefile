run-service:
	go run ./server/server.go

run-client:
	go run ./client/client.go

gen-proto:
	protoc --proto_path=./proto --go_out=plugins=grpc:./ service.proto

	# protoc --proto_path=./proto --go_out=plugins=grpc:. --go_opt=paths=source_relative service.proto