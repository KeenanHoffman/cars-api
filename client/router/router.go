package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/client/routes"
	"github.com/keenanhoffman/cars-api/proto"
)

func NewRouter(client proto.AddCarServiceClient) *gin.Engine {
	apiRouter := gin.Default()
	//Create
	apiRouter.POST("/cars", routes.CreateCar(client))
	////Read One
	apiRouter.GET("/cars/:id", routes.GetCarById(client))
	////Read All
	apiRouter.GET("/cars", routes.GetCars(client))
	////Update
	apiRouter.PATCH("/cars/:id", routes.UpdateCar(client))
	////Replace
	apiRouter.PUT("/cars/:id", routes.ReplaceCar(client))
	////Delete
	apiRouter.DELETE("/cars/:id", routes.DeleteCar(client))
	return apiRouter
}
