package services

import (
	"github.com/keenanhoffman/cars-api/proto"
)

func (s *Services) Search(request *proto.CarRequest, stream proto.AddCarService_SearchServer) error {

	cars, err := s.DB.SearchCars(proto.Car{
		Make: request.GetMake(),
		Model: request.GetModel(),
		Vin: request.GetVin(),
	})
	if err != nil {
		panic(err)
	}
	for _, car := range cars {
		err = stream.Send(car)
		if err != nil {
			return err
		}
	}
	return nil
}
