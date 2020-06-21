package main

import (
	"google.golang.org/grpc"
	"net"
	"simple-hello-grpc-server/pkg/api"
	"simple-hello-grpc-server/pkg/hello"
)

func main() {
	server := grpc.NewServer()
	helloSrv := &hello.Hello{}

	api.RegisterHelloServer(server, helloSrv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(l); err != nil {
		panic(err)
	}
}


