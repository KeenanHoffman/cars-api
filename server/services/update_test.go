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

var _ = Describe("Update", func() {
	It("Gets cars successfully", func() {
		ctx := context.Background()
		req := proto.CarRequest{
			Id:    12345,
			Make:  "new-make",
			Model: "new-model",
			Vin:   "new-vin",
		}
		mockDB := &test.MockDB{
			UpdateCarMethod: test.UpdateCarMethodStruct{
				GivenCar: proto.Car{
					Id:    12345,
					Make:  "new-make",
					Model: "new-model",
					Vin:   "new-vin",
				},
			},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Update(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusOK)))

		Expect(mockDB.UpdateCarMethod.Called).To(BeTrue())
		Expect(mockDB.UpdateCarMethod.GivenCar.GetId()).To(Equal(int64(12345)))
		Expect(mockDB.UpdateCarMethod.GivenCar.GetMake()).To(Equal("new-make"))
		Expect(mockDB.UpdateCarMethod.GivenCar.GetModel()).To(Equal("new-model"))
		Expect(mockDB.UpdateCarMethod.GivenCar.GetVin()).To(Equal("new-vin"))
	})
	It("Fails while updating a car", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			UpdateCarMethod: test.UpdateCarMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Update(ctx, &req)

		Expect(mockDB.UpdateCarMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
	})
})
