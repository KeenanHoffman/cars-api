package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/keenanhoffman/cars-api/proto"
	"github.com/keenanhoffman/cars-api/server/db"
	"github.com/keenanhoffman/cars-api/server/services"
	grcp "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	database := pg.Connect(&pg.Options{
		User: os.Getenv("SERVER_DB_USER"),
		Password: os.Getenv("SERVER_DB_PASS"),
		Database: os.Getenv("SERVER_DB_NAME"),
		Addr: fmt.Sprintf("%s:%s", os.Getenv("SERVER_DB_URL"), os.Getenv("SERVER_DB_PORT")),
	})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		panic(err)
	}
	srv := grcp.NewServer()
	proto.RegisterAddCarServiceServer(srv, &services.Services{
		DB: &db.Postgres{
			DB: database,
		},
	})
	reflection.Register(srv)

	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}
