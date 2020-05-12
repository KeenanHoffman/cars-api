package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/proto"
	"io"
)

func SearchCars(client proto.AddCarServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		car := proto.CarRequest{
			Make: ctx.Request.URL.Query().Get("make"),
			Model: ctx.Request.URL.Query().Get("model"),
			Vin: ctx.Request.URL.Query().Get("vin"),
		}
		stream, err := client.Search(ctx, &car)
		if err != nil {
			panic(err)
		}
		streamingCars := make(chan *proto.Car, 100)
		func() {
			for {
				foundCar, err := stream.Recv()
				if err == io.EOF {
					close(streamingCars)
					stream.CloseSend()
					break
				}
				if err != nil {
					panic(err)
				}
				fmt.Println(foundCar)
				streamingCars <- foundCar
			}
		}()
		ctx.Stream(func(w io.Writer) bool {
			select {
				case foundCar := <- streamingCars:
					someCar := proto.Car{
						Id: foundCar.GetId(),
						Make: foundCar.GetMake(),
						Model: foundCar.GetModel(),
						Vin: foundCar.GetVin(),
					}
					jsonCar, err := json.Marshal(someCar)
					if err != nil {
						panic(err)
					}
					ctx.SSEvent("", string(jsonCar))

				case <-streamingCars:
					return false
				}
				return true
		})
	}
}
