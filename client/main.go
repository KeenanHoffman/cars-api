package main

import (
	"fmt"
	"github.com/keenanhoffman/cars-api/client/router"
	"github.com/keenanhoffman/cars-api/proto"
	grcp "google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grcp.Dial(fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_URL"),
		os.Getenv("SERVER_PORT")),
		grcp.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddCarServiceClient(conn)
	apiRouter := router.NewRouter(client)

	err = apiRouter.Run(fmt.Sprintf(":%s", os.Getenv("CLIENT_PORT")))
	if err != nil {
		panic(err)
	}
}
