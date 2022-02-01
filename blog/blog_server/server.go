package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hasimwebapps/go_blog/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
	blogpb.BlogServiceServer
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("starting blog server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
