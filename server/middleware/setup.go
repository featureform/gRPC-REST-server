package middleware

import (
	_ "github.com/Sami1309/go-grpc-server/grpc"
)

// var CLIENT *grpc.embeddingHubClient

const (
	address = "localhost:7462"
	name    = "hello"
)

func ConnectGRPC() {

	// // Set up a connection to the server.
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	// CLIENT = pb.NewEmbeddingHubClient(conn)
}
