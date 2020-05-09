package db

import (
	"github.com/keenanhoffman/cars-api/proto"
)

type Database interface {
	CreateCar(proto.Car) error
	GetCarById(int64) (proto.Car, error)
	GetCars() ([]*proto.Car, error)
	UpdateCar(proto.Car) error
	ReplaceCar(proto.Car) error
	DeleteCar(int64) error
}

type Postgres struct {}

func (p *Postgres) CreateCar(car proto.Car) error {
	return nil
}

func (p *Postgres) GetById(id int64) (proto.Car, error) {
	return proto.Car{}, nil
}

func (p *Postgres) GetCars() ([]*proto.Car, error) {
	return []*proto.Car{}, nil
}

func (p *Postgres) UpdateCar(car proto.Car) error {
	return nil
}

func (p *Postgres) ReplaceCar(car proto.Car) error {
	return nil
}

func (p *Postgres) DeleteCar(id int64) error {
	return nil
}
