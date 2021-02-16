hello:
	echo "Hello World"
	
build:
	go build -o bin/main main.go

run:
	go run ./server/server.go