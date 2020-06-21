package hello

import (
	"context"
	"fmt"
	"simple-hello-grpc-server/pkg/api"
)

type Hello struct {
}

func (s *Hello) Hello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	name := req.Name
	hello := fmt.Sprintf("Hello, %s", name)
	response := &api.HelloResponse{Result: hello}
	return response, nil
}
