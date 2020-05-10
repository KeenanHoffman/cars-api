package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
)

func CreateCar(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCarRequest := &proto.CarRequest{}
		err := ctx.BindJSON(newCarRequest)
		if err != nil {
			panic(err)
		}
		_, err = client.Create(ctx, newCarRequest)
		if err != nil {
			panic(err)
			//ctx.JSON(http.StatusServiceUnavailable, gin.H{
			//	"error": err.Error(),
			//})
			//ctx.Abort()
		}
		ctx.JSON(http.StatusCreated, gin.H{})
		ctx.Abort()
	}
}
