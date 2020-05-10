package services

import (
	"github.com/keenanhoffman/cars-api/proto"
	"context"
	"net/http"
)

func (s *Services) GetAll(ctx context.Context, request *proto.CarRequest) (*proto.CarsResponse, error) {
	cars, err := s.DB.GetCars()
	if err != nil {
		return &proto.CarsResponse{
			Status: http.StatusServiceUnavailable,
		}, err
	}
	return &proto.CarsResponse{
		Status: 200,
		Cars: cars,
	}, nil
}
