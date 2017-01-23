package main

import (
	"log"

	"github.com/adamwg/proto-example/model"
	"github.com/adamwg/proto-example/web"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
)

var (
	rpcAddr = envflag.String("RPC_ADDR", "localhost:8080", "The RPC address for the Posts service")
)

func main() {
	envflag.Parse()
	client := getPostsClient()

	srv := web.NewServer(client)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getPostsClient() model.PostsClient {
	if *rpcAddr == "" {
		log.Fatal("must provide an RPC address")
	}

	conn, err := grpc.Dial(*rpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return model.NewPostsClient(conn)
}
