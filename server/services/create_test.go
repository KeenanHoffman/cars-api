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
		server := Services{
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
		server := Services{
			DB: mockDB,
		}
		response, err := server.Create(ctx, &req)

		Expect(mockDB.CreateMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
	})
	It("Fails when given an ID", func() {
		ctx := context.Background()
		req := proto.CarRequest{
			Id: 12345,
		}
		mockDB := &test.MockDB{}
		server := Services{
			DB: mockDB,
		}
		response, err := server.Create(ctx, &req)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("ID provided for a new car"))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusBadRequest)))

		Expect(mockDB.CreateMethod.Called).To(BeFalse())
	})
})
