package services

import (
	"github.com/keenanhoffman/cars-api/proto"
	"context"
	"net/http"
)

func (s *Services) Delete(ctx context.Context, request *proto.CarRequest) (*proto.SimpleResponse, error) {
	err := s.DB.DeleteCar(request.GetId())
	if err != nil {
		return &proto.SimpleResponse{
			Status: http.StatusServiceUnavailable,
		}, err
	}
	return &proto.SimpleResponse{
		Status: http.StatusOK,
	}, nil
}
