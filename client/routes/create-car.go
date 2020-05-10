package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/proto"
	"io/ioutil"
	"net/http"
	protobuf "github.com/golang/protobuf/proto"
)

func CreateCar(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCarRequest := &proto.CarRequest{}
		accept := ctx.GetHeader("Accept")
		if accept == "application/json" {
			err := ctx.BindJSON(newCarRequest)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid json: %s", err.Error()),
				})
				ctx.Abort()
				return
			}
		}
		if accept == "application/xml" {
			err := ctx.BindXML(newCarRequest)
			if err != nil {
				ctx.XML(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid xml: %s", err.Error()),
				})
				ctx.Abort()
				return
			}
		}
		if accept == "application/protobuf" {
			defer ctx.Request.Body.Close()
			body, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				responseError := &proto.SimpleError{
					Error: fmt.Sprintf(`failure reading protobuf body: %s`, err.Error()),
				}
				protobufResponse, _ := protobuf.Marshal(responseError)
				ctx.Writer.WriteHeader(http.StatusServiceUnavailable)
				ctx.Writer.Write(protobufResponse)
				ctx.Abort()
				return
			}
			err = protobuf.Unmarshal(body, newCarRequest)
			if err != nil {
				responseError := &proto.SimpleError{
					Error: fmt.Sprintf(`invalid protobuf: %s`, err.Error()),
				}
				protobufResponse, err := protobuf.Marshal(responseError)
				if err != nil {
					responseError := &proto.SimpleError{
						Error: fmt.Sprintf(`failure marshaling protobuf error response: %s`, err.Error()),
					}
					protobufResponse, _ := protobuf.Marshal(responseError)
					ctx.Writer.WriteHeader(http.StatusServiceUnavailable)
					ctx.Writer.Write(protobufResponse)
					ctx.Abort()
					return
				}
				ctx.Writer.WriteHeader(http.StatusBadRequest)
				ctx.Writer.Write(protobufResponse)
				ctx.Abort()
				return
			}
		}
		clientResponse, err := client.Create(ctx, newCarRequest)
		if err != nil {
			ctx.JSON(int(clientResponse.Status), gin.H{
				"error": fmt.Sprintf(`grcp client: %s`, err.Error()),
			})
			ctx.Abort()
			return
		}
		ctx.Writer.WriteHeader(http.StatusCreated)
		ctx.Writer.Write([]byte(""))
		ctx.Abort()
		return
	}
}
