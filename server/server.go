package main

import (
	"fmt"
	"os"

	"github.com/Sami1309/go-grpc-server/middleware"
	"github.com/Sami1309/go-grpc-server/router"
)

func main() {

	r := router.Router()

	middleware.ConnectGRPC()

	os.Setenv("PORT", "8080")

	myport := fmt.Sprintf(":%s", os.Getenv("PORT"))

	r.Run(myport)
}
