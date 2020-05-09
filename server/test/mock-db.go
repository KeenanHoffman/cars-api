package test

import (
	"github.com/keenanhoffman/cars-api/proto"
)

type MockDB struct {
	CreateMethod     CreateMethodStruct
	GetCarByIdMethod GetCarByIdMethodStruct
	GetCarsMethod    GetCarsMethodStruct
	UpdateCarMethod  UpdateCarMethodStruct
	ReplaceCarMethod  ReplaceCarMethodStruct
	DeleteCarMethod  DeleteCarMethodStruct
}

type CreateMethodStruct struct {
	Called      bool
	ReturnError error
	GivenCar    proto.Car
}
func(m *MockDB) CreateCar(car proto.Car) error {
	m.CreateMethod.Called = true
	m.CreateMethod.GivenCar = car
	return m.CreateMethod.ReturnError
}

type GetCarByIdMethodStruct struct {
	Called      bool
	GivenId     int64
	ReturnCar   proto.Car
	ReturnError error
}
func(m *MockDB) GetCarById(id int64) (proto.Car, error) {
	m.GetCarByIdMethod.Called = true
	m.GetCarByIdMethod.GivenId = id
	return m.GetCarByIdMethod.ReturnCar, m.GetCarByIdMethod.ReturnError
}

type GetCarsMethodStruct struct {
	Called      bool
	ReturnCars  []*proto.Car
	ReturnError error
}
func(m *MockDB) GetCars() ([]*proto.Car, error) {
	m.GetCarsMethod.Called = true
	return m.GetCarsMethod.ReturnCars, m.GetCarsMethod.ReturnError
}

type UpdateCarMethodStruct struct {
	Called      bool
	GivenCar    proto.Car
	ReturnError error
}
func(m *MockDB) UpdateCar(car proto.Car) error {
	m.UpdateCarMethod.Called = true
	m.UpdateCarMethod.GivenCar = car
	return m.UpdateCarMethod.ReturnError
}

type ReplaceCarMethodStruct struct {
	Called      bool
	GivenCar    proto.Car
	ReturnError error
}
func(m *MockDB) ReplaceCar(car proto.Car) error {
	m.ReplaceCarMethod.Called = true
	m.ReplaceCarMethod.GivenCar = car
	return m.ReplaceCarMethod.ReturnError
}

type DeleteCarMethodStruct struct {
	Called      bool
	GivenId     int64
	ReturnError error
}
func(m *MockDB) DeleteCar(id int64) error {
	m.DeleteCarMethod.Called = true
	m.DeleteCarMethod.GivenId = id
	return m.DeleteCarMethod.ReturnError
}
