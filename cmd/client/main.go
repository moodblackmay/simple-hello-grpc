package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"simple-hello-grpc-server/pkg/api"
)

func main() {
	name := os.Args[1]

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewHelloClient(conn)

	resp, err := client.Hello(context.Background(), &api.HelloRequest{Name: name})
	fmt.Println(resp)
}