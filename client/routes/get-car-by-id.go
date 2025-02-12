package routes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
	"strconv"
)

func GetCarById(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		idNum, _ := strconv.ParseInt(id, 10, 32)
		newCarRequest := &proto.CarRequest{
			Id: idNum,
		}
		clientResponse, err := client.GetById(ctx, newCarRequest)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"error": fmt.Sprintf(`grcp client: %s`, err.Error()),
			})
			ctx.Abort()
			return
		}
		accept := ctx.GetHeader("Accept")
		if accept == "application/json" {
			jsonResponse, _ := json.Marshal(clientResponse.Car)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(jsonResponse)
			ctx.Abort()
			return
		}
		if accept == "application/xml" {
			xmlResponse, _ := xml.Marshal(clientResponse.Car)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(xmlResponse)
			ctx.Abort()
			return
		}
		if accept == "application/protobuf" {
			protobufResponse, _ := protobuf.Marshal(clientResponse.Car)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(protobufResponse)
			ctx.Abort()
			return
		}
	}
}
