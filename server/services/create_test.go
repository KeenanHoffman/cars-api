package services_test

import (
	"errors"
	"github.com/keenanhoffman/cars-api/proto"
	. "github.com/keenanhoffman/cars-api/server/services"
	"github.com/keenanhoffman/cars-api/server/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"context"
	"net/http"
)


var _ = Describe("Create", func() {
	It("Creates a new car successfully", func() {
		ctx := context.Background()
		req := proto.CarRequest{
			Make:  "test-make",
			Model: "test-model",
			Vin:   "test-vin",
		}
		mockDB := &test.MockDB{}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Create(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusCreated)))

		Expect(mockDB.CreateMethod.Called).To(BeTrue())
		Expect(mockDB.CreateMethod.GivenCar.GetMake()).To(Equal("test-make"))
		Expect(mockDB.CreateMethod.GivenCar.GetModel()).To(Equal("test-model"))
		Expect(mockDB.CreateMethod.GivenCar.GetVin()).To(Equal("test-vin"))
	})
	It("Fails while storing the car", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			CreateMethod: test.CreateMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Create(ctx, &req)

		Expect(mockDB.CreateMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
	})
})
