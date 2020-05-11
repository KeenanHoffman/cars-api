package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/client/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	. "github.com/keenanhoffman/cars-api/client/routes"
)

var _ = Describe("DeleteCar", func() {
	It("Deletes a car given an id successfully", func() {
		respRecorder := httptest.NewRecorder()
		_, router := gin.CreateTestContext(respRecorder)
		mockClient := test.MockClient{}
		router.DELETE("/cars/:id", DeleteCar(&mockClient))
		request, err := http.NewRequest("DELETE", "/cars/12345", bytes.NewBufferString(""))
		Expect(err).ToNot(HaveOccurred())

		router.ServeHTTP(respRecorder, request)
		Expect(respRecorder.Body.String()).To(Equal(""))
		Expect(respRecorder.Code).To(Equal(http.StatusOK))

		Expect(mockClient.DeleteMethod.Called).To(BeTrue())
		Expect(mockClient.DeleteMethod.GivenReq.GetId()).To(Equal(int64(12345)))
	})
	It(`Returns "Service Unavailable" when client.Delete fails`, func() {
		respRecorder := httptest.NewRecorder()
		_, router := gin.CreateTestContext(respRecorder)
		clientError := errors.New("Client Error")
		mockClient := test.MockClient{
			DeleteMethod: test.DeleteMethodStruct{
				ReturnError: clientError,
			},
		}
		router.DELETE("/cars/:id", DeleteCar(&mockClient))
		jsonBody := "{}"
		request, err := http.NewRequest("DELETE", "/cars/12345", bytes.NewBufferString(jsonBody))
		Expect(err).ToNot(HaveOccurred())
		request.Header.Add("Content-type", "application/json")

		router.ServeHTTP(respRecorder, request)
		expectedResponseMap := map[string]string{
			"error": fmt.Sprintf(`grcp client: %s`, clientError.Error()),
		}
		expectedResponse, err := json.Marshal(expectedResponseMap)
		Expect(err).ToNot(HaveOccurred())

		Expect(mockClient.DeleteMethod.Called).To(BeTrue())
		Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
		Expect(respRecorder.Code).To(Equal(http.StatusServiceUnavailable))
	})
})
