package routes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/keenanhoffman/cars-api/proto"
	"net/http"
)

func GetCars(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientResponse, err := client.GetAll(ctx, &proto.CarRequest{})
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"error": fmt.Sprintf(`grcp client: %s`, err.Error()),
			})
			ctx.Abort()
			return
		}
		accept := ctx.GetHeader("Accept")
		if accept == "application/json" {
			jsonResponse, _ := json.Marshal(clientResponse.Cars)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(jsonResponse)
			ctx.Abort()
			return
		}
		if accept == "application/xml" {
			xmlResponse, _ := xml.Marshal(clientResponse.Cars)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(xmlResponse)
			ctx.Abort()
			return
		}
		if accept == "application/protobuf" {
			protobufResponse, _ := protobuf.Marshal(clientResponse)
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write(protobufResponse)
			ctx.Abort()
			return
		}
	}
}
