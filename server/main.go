package main

import (
	"github.com/keenanhoffman/cars-api/proto"
	"golang.org/x/net/context"
	"golang.org/x/text/message/catalog"
)

type server struct{}

func (s *server) Create(ctx context.Context, request *proto.CarRequest) (*proto.StatusResponse, error) {

}

func main() {

}