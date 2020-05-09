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

var _ = Describe("Replace", func() {
	It("Replaces a car successfully", func() {
		ctx := context.Background()
		req := proto.CarRequest{
			Id:    12345,
			Make:  "new-make",
			Model: "new-model",
			Vin:   "new-vin",
		}
		mockDB := &test.MockDB{
			ReplaceCarMethod: test.ReplaceCarMethodStruct{},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Replace(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusOK)))

		Expect(mockDB.ReplaceCarMethod.Called).To(BeTrue())
		Expect(mockDB.ReplaceCarMethod.GivenCar.GetId()).To(Equal(int64(12345)))
		Expect(mockDB.ReplaceCarMethod.GivenCar.GetMake()).To(Equal("new-make"))
		Expect(mockDB.ReplaceCarMethod.GivenCar.GetModel()).To(Equal("new-model"))
		Expect(mockDB.ReplaceCarMethod.GivenCar.GetVin()).To(Equal("new-vin"))
	})
	It("Fails while replacing a car", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			ReplaceCarMethod: test.ReplaceCarMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Replace(ctx, &req)

		Expect(mockDB.ReplaceCarMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
	})
})
