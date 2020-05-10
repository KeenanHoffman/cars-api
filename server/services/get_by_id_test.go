package services_test

import (
	"errors"
	"github.com/keenanhoffman/cars-api/proto"
	"github.com/keenanhoffman/cars-api/server/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"context"

	. "github.com/keenanhoffman/cars-api/server/services"
)

var _ = Describe("GetCarById", func() {
	It("Gets a car by ID successfully", func() {
		ctx := context.Background()
		id := int64(12345)
		req := proto.CarRequest{
			Id: id,
		}
		mockDB := &test.MockDB{
			GetCarByIdMethod: test.GetCarByIdMethodStruct{
				ReturnCar: proto.Car{
					Id:    id,
					Make:  "test-make",
					Model: "test-model",
					Vin:   "test-vin",
				},
			},
		}
		server := Services{
			DB: mockDB,
		}
		response, err := server.GetById(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusOK)))
		Expect(response.Car.GetId()).To(Equal(id))
		Expect(response.Car.GetMake()).To(Equal("test-make"))
		Expect(response.Car.GetModel()).To(Equal("test-model"))
		Expect(response.Car.GetVin()).To(Equal("test-vin"))

		Expect(mockDB.GetCarByIdMethod.Called).To(BeTrue())
	})
	It("Fails while getting the car", func() {
		ctx := context.Background()
		id := int64(12345)
		req := proto.CarRequest{
			Id: id,
		}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			GetCarByIdMethod: test.GetCarByIdMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Services{
			DB: mockDB,
		}
		response, err := server.GetById(ctx, &req)

		Expect(mockDB.GetCarByIdMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
		Expect(response.GetCar()).To(BeNil())
	})
})
