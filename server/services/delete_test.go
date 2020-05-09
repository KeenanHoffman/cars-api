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

var _ = Describe("Delete", func() {
	It("Deletes a car successfully", func() {
		ctx := context.Background()
		req := proto.CarRequest{
			Id:    12345,
		}
		mockDB := &test.MockDB{
			DeleteCarMethod: test.DeleteCarMethodStruct{},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Delete(ctx, &req)

		Expect(err).ToNot(HaveOccurred())
		Expect(response.GetStatus()).To(Equal(int32(http.StatusOK)))

		Expect(mockDB.DeleteCarMethod.Called).To(BeTrue())
		Expect(mockDB.DeleteCarMethod.GivenId).To(Equal(int64(12345)))
	})
	It("Fails while deleting a car", func() {
		ctx := context.Background()
		req := proto.CarRequest{}
		dbError := errors.New("DB Error")
		mockDB := &test.MockDB{
			DeleteCarMethod: test.DeleteCarMethodStruct{
				ReturnError: dbError,
			},
		}
		server := Server{
			DB: mockDB,
		}
		response, err := server.Delete(ctx, &req)

		Expect(mockDB.DeleteCarMethod.Called).To(BeTrue())
		Expect(err).To(Equal(dbError))
		Expect(response.GetStatus()).To(Equal(int32(http.StatusServiceUnavailable)))
	})
})
