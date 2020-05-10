package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
	"strconv"
)

func DeleteCar(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		idNum, _ := strconv.ParseInt(id, 10, 32)
		deleteCarRequest := &proto.CarRequest{
			Id: idNum,
		}
		clientResponse, err := client.Delete(ctx, deleteCarRequest)
		if err != nil {
			ctx.JSON(int(clientResponse.Status), gin.H{
				"error": fmt.Sprintf(`grcp client: %s`, err.Error()),
			})
			ctx.Abort()
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)
		ctx.Writer.Write([]byte(""))
		ctx.Abort()
		return
	}
}
