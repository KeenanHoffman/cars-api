package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/keenanhoffman/cars-api/db"
	"github.com/keenanhoffman/cars-api/stuff/structs"
  "log"
  "net/http"
)

func CreateCar(carDB db.DB) gin.HandlerFunc {
  return func(context *gin.Context) {
    carName, _ := context.Params.Get("car_name")
    var newCar structs.Car
    err := context.BindJSON(&newCar)
    if err != nil {
     log.Fatal(err)
    }
    newCar.Name = carName
  	err = carDB.CreateCar(newCar)
  	if err != nil {
  		log.Fatal(err)
    }

    context.Writer.WriteHeader(http.StatusCreated)
    context.Writer.Write([]byte(""))
    context.Abort()
  }
}
