package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/routes"
	"github.com/keenanhoffman/cars-api/structs"
)

type MockDB struct { Called      bool
	ReturnError error
}

func(m *MockDB) CreateCar(car structs.Car) (error) {
	m.Called = true
	return m.ReturnError
}

func main() {
	router := gin.Default()
	mockDB := MockDB{ReturnError: nil}
	router.GET("/cars", routes.CreateCar(&mockDB))
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}