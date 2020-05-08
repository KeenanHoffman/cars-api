package db

import (
	"github.com/keenanhoffman/cars-api/stuff/structs"
)

type DB interface {
	CreateCar(structs.Car) (error)
}

