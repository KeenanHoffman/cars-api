package db

import (
	"github.com/keenanhoffman/cars-api/structs"
)

type DB interface {
	CreateCar(structs.Car) (error)
}

