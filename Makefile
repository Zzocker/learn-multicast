.PHONY: protos server

protos:
	protoc -I protos --go_out ${GOPATH}/src --go-grpc_out ${GOPATH}/src  protos/server.proto

cli:
	go build -o .build/multicast-cli cmd/cli/main.go

server:
	go build -o .build/multicast cmd/server/main.go