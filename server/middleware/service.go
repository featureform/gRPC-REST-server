package middleware

import (
	//remote packages
	"context"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/Sami1309/go-grpc-server/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetSpaces(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	CLIENT := pb.NewEmbeddingHubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, getSpacesError := CLIENT.DownloadSpaces(ctx, &pb.DownloadSpacesRequest{})
	if getSpacesError != nil {
		log.Fatalf("Error message: ( %v)", getSpacesError)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Space received: %s", resp.String())
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
	c.JSON(http.StatusOK, gin.H{"Spaces": "space"})
}

func GetSpace(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	CLIENT := pb.NewEmbeddingHubClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// name := c.Param("name")
	getResponse, getResponseErr := CLIENT.Get(ctx, &pb.GetRequest{Key: "test-key", Space: "test-space"})
	if getResponseErr != nil {
		log.Fatalf("Error message: ( %v)", getResponseErr)
	}

	log.Printf("Greeting: %s", getResponse.GetEmbedding())
	c.JSON(http.StatusOK, gin.H{"Space": "space"})

}

func GetSpaceVectors(c *gin.Context) {
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	// CLIENT := pb.NewEmbeddingHubClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	c.JSON(http.StatusOK, gin.H{"space vectors": "vec"})
}

func GetNearestNeighbors(c *gin.Context) {
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	// CLIENT := pb.NewEmbeddingHubClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	c.JSON(http.StatusOK, gin.H{"Nearest neighbors": "nn"})
}
