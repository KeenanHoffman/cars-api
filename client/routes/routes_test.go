package routes_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/client/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	. "github.com/keenanhoffman/cars-api/client/routes"
)

var _ = Describe("Routes", func() {
	Context("CreateCar", func() {
		Context("When given JSON", func() {
			It("Creates a new car successfully", func() {
				respRecorder := httptest.NewRecorder()
				_, router := gin.CreateTestContext(respRecorder)
				mockClient := test.MockClient{}
				router.POST("/cars", CreateCar(&mockClient))
				reqBody := map[string]string{
					"make": "test-make",
					"model": "test-model",
					"vin": "test-vin",
				}
				jsonBody, err := json.Marshal(reqBody)
				Expect(err).ToNot(HaveOccurred())
				request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(string(jsonBody)))
				Expect(err).ToNot(HaveOccurred())
				request.Header.Add("Accept", "application/json")

				router.ServeHTTP(respRecorder, request)
				Expect(respRecorder.Body.String()).To(Equal("{}"))
				Expect(respRecorder.Code).To(Equal(http.StatusCreated))

				Expect(mockClient.CreateMethod.Called).To(BeTrue())
				Expect(mockClient.CreateMethod.GivenReq.GetMake()).To(Equal("test-make"))
				Expect(mockClient.CreateMethod.GivenReq.GetModel()).To(Equal("test-model"))
				Expect(mockClient.CreateMethod.GivenReq.GetVin()).To(Equal("test-vin"))
			})
		})
		//Context("When given Protobuf", func() {
		//	It("Creates a new car successfully", func() {
		//		respRecorder := httptest.NewRecorder()
		//		_, router := gin.CreateTestContext(respRecorder)
		//		mockDB := MockDB{ReturnError: nil}
		//		router.POST("/cars/:car_name", routes.CreateCar(&mockDB))
		//		protobufBody :=
		//		request, err := http.NewRequest("POST", "/cars/test_car", bytes.NewBufferString("{}"))
		//		Expect(err).ToNot(HaveOccurred())
		//		request.Header.Add("Accept", "application/protobuf")
		//
		//		router.ServeHTTP(respRecorder, request)
		//		Expect(respRecorder.Body.String()).To(Equal(""))
		//		Expect(respRecorder.Code).To(Equal(http.StatusCreated))
		//	})
		//})
	})
})
