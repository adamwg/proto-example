package main

import (
	"log"
	"net"

	"github.com/adamwg/proto-example/model"
	"github.com/adamwg/proto-example/server"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen")
	}

	srv := grpc.NewServer()
	model.RegisterPostsServer(srv, server.New())
	srv.Serve(l)
}
