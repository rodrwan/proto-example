package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/pkg/errors"
	pb "github.com/rodrwan/proto-example/proto"
)

func main() {
	IP := fmt.Sprintf("%s:%d", "", 8080)
	fmt.Printf("GRPC Service IP: %s\n", IP)
	conn, err := grpc.Dial(IP, grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrap(err, "did not connect"))
	}

	defer conn.Close()

	usc := pb.NewGreetingClient(conn)

	resp, err := usc.Hello(context.Background(), &pb.Request{
		Message: "grpc",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetData())
}
