package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/keenanhoffman/cars-api/proto"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UpdateCar(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		updateCarRequest := &proto.CarRequest{}
		accept := ctx.GetHeader("Content-type")
		if accept == "application/json" {
			err := ctx.BindJSON(updateCarRequest)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid json: %s", err.Error()),
				})
				ctx.Abort()
				return
			}
		}
		if accept == "application/xml" {
			err := ctx.BindXML(updateCarRequest)
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
			err = protobuf.Unmarshal(body, updateCarRequest)
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

		id, _ := ctx.Params.Get("id")
		idNum, _ := strconv.ParseInt(id, 10, 32)
		updateCarRequest.Id = idNum
		clientResponse, err := client.Update(ctx, updateCarRequest)
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
