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
	//apiRouter.GET("/cars/:car_id", )
	////Read All
	//apiRouter.GET("/cars", )
	////Update
	//apiRouter.PATCH("/cars/:car_id", )
	////Replace
	//apiRouter.PUT("/cars/:car_id", )
	////Delete
	//apiRouter.DELETE("/cars/:car_id", )

	return apiRouter
}
