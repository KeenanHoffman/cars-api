package routes_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/routes"
	"github.com/keenanhoffman/cars-api/structs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	//"github.com/golang/protobuf/proto"
)

type MockDB struct { Called      bool
	ReturnError error
}

func(m *MockDB) CreateCar(car structs.Car) (error) {
	m.Called = true
	return m.ReturnError
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
	Context("When given Protobuf", func() {
		It("Creates a new car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockDB := MockDB{ReturnError: nil}
			router.POST("/cars/:car_name", routes.CreateCar(&mockDB))
			protobufBody :=
			request, err := http.NewRequest("POST", "/cars/test_car", bytes.NewBufferString("{}"))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/protobuf")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusCreated))
		})
	})
})
