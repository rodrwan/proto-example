package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/rodrwan/proto-example/proto"
	"github.com/rodrwan/proto-example/service"
	"google.golang.org/grpc"
)

func main() {
	port := 8080
	srv := grpc.NewServer()
	sg := &service.Greetings{}
	pb.RegisterGreetingServer(srv, sg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Users Service...")
	log.Println(fmt.Sprintf("Listening on: %d", port))
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
