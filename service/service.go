package service

import (
	"context"
	"fmt"

	pb "github.com/rodrwan/proto-example/proto"
)

type Greetings struct {
}

func (g *Greetings) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	msg := req.GetMessage()

	return &pb.Response{
		Data: fmt.Sprintf("Hello: %s\n", msg),
	}, nil
}
