package db

import (
	"github.com/go-pg/pg"
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

type Postgres struct {
	DB *pg.DB
}

func (p *Postgres) CreateCar(car proto.Car) error {
	err := p.DB.Insert(&car)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetCarById(id int64) (proto.Car, error) {
	car := proto.Car{}
	err := p.DB.Model(&car).
		Where("car.id = ?", id).
		Select()
	if err != nil {
		return proto.Car{}, err
	}

	return proto.Car{
		Id: car.Id,
		Make: car.Make,
		Model: car.Model,
		Vin: car.Vin,
	}, nil
}

func (p *Postgres) GetCars() ([]*proto.Car, error) {
	cars := []*proto.Car{}
	err := p.DB.Model(&cars).Select()
	if err != nil {
		return []*proto.Car{}, err
	}

	return cars, nil
}

func (p *Postgres) UpdateCar(car proto.Car) error {
	currentCar := proto.Car{}
	err := p.DB.Model(&currentCar).
		Where("car.id = ?", car.Id).
		Select()
	if err != nil {
		return err
	}
	if car.Make == "" {
		car.Make = currentCar.Make
	}
	if car.Model == "" {
		car.Model = currentCar.Model
	}
	if car.Vin == "" {
		car.Vin = currentCar.Vin
	}

	err = p.DB.Update(&car)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) ReplaceCar(car proto.Car) error {
	err := p.DB.Update(&car)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteCar(id int64) error {
	err := p.DB.Delete(&proto.Car{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}
