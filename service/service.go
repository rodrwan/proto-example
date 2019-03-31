package service

import (
	"context"
	"fmt"

	pb "github.com/rodrwan/proto-example/proto"
)

// Greetings implements grpc interface.
type Greetings struct {
	Prefix string
}

// Hello send a new greetings message.
func (g *Greetings) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	msg := req.GetMessage()

	return &pb.Response{
		Data: fmt.Sprintf("%s: %s\n", g.Prefix, msg),
	}, nil
}
