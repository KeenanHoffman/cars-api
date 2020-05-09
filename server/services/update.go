package services

import (
	"context"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
)

func (s *Server) Update(ctx context.Context, request *proto.CarRequest) (*proto.SimpleResponse, error) {
	err := s.DB.UpdateCar(proto.Car{
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

