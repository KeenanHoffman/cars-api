package services

import (
	"context"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
)

func (s *Services) Create(ctx context.Context, request *proto.CarRequest) (*proto.SimpleResponse, error) {
	err := s.DB.CreateCar(proto.Car{
		Make: request.GetMake(),
		Model: request.GetModel(),
		Vin: request.GetVin(),
	})
	if err != nil {
		return &proto.SimpleResponse{
			Status: http.StatusServiceUnavailable,
		}, err
	}

	return &proto.SimpleResponse{
		Status: http.StatusCreated,
	}, nil
}
