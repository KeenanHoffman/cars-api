package services

import (
	"github.com/keenanhoffman/cars-api/proto"
	"context"
	"net/http"
)

func (s *Services) Replace(ctx context.Context, request *proto.CarRequest) (*proto.SimpleResponse, error) {
	err := s.DB.ReplaceCar(proto.Car{
		Id:    request.GetId(),
		Make:  request.GetMake(),
		Model: request.GetModel(),
		Vin:   request.GetVin(),
	})
	if err != nil {
		return &proto.SimpleResponse{
			Status: http.StatusServiceUnavailable,
		}, err
	}
	return &proto.SimpleResponse{
		Status: http.StatusOK,
	}, nil
}
