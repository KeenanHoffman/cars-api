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

var _ = Describe("GetAll", func() {
	It("Gets cars successfully", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		mockDB := &test.MockDB{
			GetCarsMethod: test.GetCarsMethodStruct{
				ReturnCars: []*proto.Car{
					{
						Id:    12345,
						Make:  "test-make",
						Model: "test-model",
						Vin:   "test-vin",
					},
				},
			},
		}
		server := Services{
			DB: mockDB,
		}
		response, err := server.GetAll(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusOK)))
		Expect(response.GetCars()[0].GetId()).To(Equal(int64(12345)))
		Expect(response.GetCars()[0].GetMake()).To(Equal("test-make"))
		Expect(response.GetCars()[0].GetModel()).To(Equal("test-model"))
		Expect(response.GetCars()[0].GetVin()).To(Equal("test-vin"))

		Expect(mockDB.GetCarsMethod.Called).To(BeTrue())
	})
	It("Fails while getting cars", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			GetCarsMethod: test.GetCarsMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Services{
			DB: mockDB,
		}
		response, err := server.GetAll(ctx, &req)

		Expect(mockDB.GetCarsMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
		Expect(len(response.GetCars())).To(Equal(0))
	})
})
