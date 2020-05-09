package services_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/stuff/routes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type MockDB struct { Called      bool
	ReturnError error
}

var _ = Describe("Create", func() {
	Context("When given JSON", func() {
		It("Creates a new car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockDB := MockDB{ReturnError: nil}
			router.POST("/cars/:car_name", routes.CreateCar(&mockDB))
			request, err := http.NewRequest("POST", "/cars/test_car", bytes.NewBufferString("{}"))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/json")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusCreated))
		})
	})
})
