package routes_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/keenanhoffman/cars-api/client/test"
	. "github.com/keenanhoffman/cars-api/client/routes"
	"github.com/keenanhoffman/cars-api/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("UpdateCar", func() {
	Context("When given accept json header", func() {
		It("Updates a car given an id successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			reqBody := map[string]string{
				"make":  "new-make",
			}
			jsonBody, err := json.Marshal(reqBody)
			Expect(err).ToNot(HaveOccurred())
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(string(jsonBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/json")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.UpdateMethod.Called).To(BeTrue())
			Expect(mockClient.UpdateMethod.GivenReq.GetMake()).To(Equal("new-make"))
			Expect(mockClient.UpdateMethod.GivenReq.GetId()).To(Equal(int64(12345)))
		})
		It(`Returns "Bad Request" when given invalid json`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			jsonBody := "}{"
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(jsonBody))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/json")

			router.ServeHTTP(respRecorder, request)
			expectedResponseMap := map[string]string{
				"error": `invalid json: invalid character '}' looking for beginning of value`,
			}
			expectedResponse, err := json.Marshal(expectedResponseMap)
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.UpdateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	Context("When given accept protobuf header", func() {
		It("Updates a car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			reqBody := proto.CarRequest{
				Make: "new-make",
			}
			protobufBody, err := protobuf.Marshal(&reqBody)
			Expect(err).ToNot(HaveOccurred())
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(string(protobufBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/protobuf")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.UpdateMethod.Called).To(BeTrue())
			Expect(mockClient.UpdateMethod.GivenReq.GetMake()).To(Equal("new-make"))
			Expect(mockClient.UpdateMethod.GivenReq.GetId()).To(Equal(int64(12345)))
		})
		It(`Returns "Bad Request" when given invalid protobuf`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			protobufBody := "}{"
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(string(protobufBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/protobuf")

			router.ServeHTTP(respRecorder, request)
			expectedResponseStruct := &proto.SimpleError{
				Error: `invalid protobuf: unexpected EOF`,
			}
			expectedResponse, err := protobuf.Marshal(expectedResponseStruct)
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.UpdateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	Context("When given accept xml header", func() {
		It("Updates a car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			reqBody := proto.CarRequest{
				Make: "new-make",
			}
			xmlBody, err := xml.Marshal(reqBody)
			Expect(err).ToNot(HaveOccurred())
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(string(xmlBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/xml")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.UpdateMethod.Called).To(BeTrue())
			Expect(mockClient.UpdateMethod.GivenReq.GetMake()).To(Equal("new-make"))
			Expect(mockClient.UpdateMethod.GivenReq.GetId()).To(Equal(int64(12345)))
		})
		It(`Returns "Bad Request" when given invalid xml`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.PATCH("/cars/:id", UpdateCar(&mockClient))
			xmlBody := "}{"
			request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(xmlBody))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Content-type", "application/xml")

			router.ServeHTTP(respRecorder, request)
			expectedResponse := "<map><error>invalid xml: EOF</error></map>"
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.UpdateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	It(`Returns "Service Unavailable" when client.Delete fails`, func() {
		respRecorder := httptest.NewRecorder()
		_, router := gin.CreateTestContext(respRecorder)
		clientError := errors.New("Client Error")
		mockClient := test.MockClient{
			UpdateMethod: test.UpdateMethodStruct{
				ReturnError: clientError,
				ReturnSimpleResponse: &proto.SimpleResponse{
					Status: http.StatusServiceUnavailable,
				},
			},
		}
		router.PATCH("/cars/:id", UpdateCar(&mockClient))
		jsonBody := "{}"
		request, err := http.NewRequest("PATCH", "/cars/12345", bytes.NewBufferString(jsonBody))
		Expect(err).ToNot(HaveOccurred())
		request.Header.Add("Content-type", "application/json")

		router.ServeHTTP(respRecorder, request)
		expectedResponseMap := map[string]string{
			"error": fmt.Sprintf(`grcp client: %s`, clientError.Error()),
		}
		expectedResponse, err := json.Marshal(expectedResponseMap)
		Expect(err).ToNot(HaveOccurred())

		Expect(mockClient.UpdateMethod.Called).To(BeTrue())
		Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
		Expect(respRecorder.Code).To(Equal(http.StatusServiceUnavailable))
	})
})
