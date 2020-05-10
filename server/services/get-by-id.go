package services

import (
	"github.com/keenanhoffman/cars-api/proto"
	"context"
	"net/http"
)

func (s *Services) GetById(ctx context.Context, request *proto.CarRequest) (*proto.CarResponse, error) {
	car, err := s.DB.GetCarById(request.GetId())
	if err != nil {
		return &proto.CarResponse{
			Status: http.StatusServiceUnavailable,
			Car: nil,
		}, err
	}

	return &proto.CarResponse{
		Status: http.StatusOK,
		Car: &car,
	}, nil
}
