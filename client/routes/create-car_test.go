package routes_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenanhoffman/cars-api/client/test"
	"github.com/keenanhoffman/cars-api/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	. "github.com/keenanhoffman/cars-api/client/routes"
	protobuf "github.com/golang/protobuf/proto"
)

var _ = Describe("CreateCar", func() {
	Context("When given accept json header", func() {
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
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusCreated))

			Expect(mockClient.CreateMethod.Called).To(BeTrue())
			Expect(mockClient.CreateMethod.GivenReq.GetMake()).To(Equal("test-make"))
			Expect(mockClient.CreateMethod.GivenReq.GetModel()).To(Equal("test-model"))
			Expect(mockClient.CreateMethod.GivenReq.GetVin()).To(Equal("test-vin"))
		})
		It(`Returns "Bad Request" when given invalid json`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.POST("/cars", CreateCar(&mockClient))
			jsonBody := "}{"
			request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(jsonBody))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/json")

			router.ServeHTTP(respRecorder, request)
			expectedResponseMap := map[string]string{
				"error": `invalid json: invalid character '}' looking for beginning of value`,
			}
			expectedResponse, err := json.Marshal(expectedResponseMap)
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.CreateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	Context("When given accept protobuf header", func() {
		It("Creates a new car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.POST("/cars", CreateCar(&mockClient))
			reqBody := proto.CarRequest{
				Make: "test-make",
				Model: "test-model",
				Vin: "test-vin",
			}
			protobufBody, err := protobuf.Marshal(&reqBody)
			Expect(err).ToNot(HaveOccurred())
			request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(string(protobufBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/protobuf")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusCreated))

			Expect(mockClient.CreateMethod.Called).To(BeTrue())
			Expect(mockClient.CreateMethod.GivenReq.GetMake()).To(Equal("test-make"))
			Expect(mockClient.CreateMethod.GivenReq.GetModel()).To(Equal("test-model"))
			Expect(mockClient.CreateMethod.GivenReq.GetVin()).To(Equal("test-vin"))
		})
		It(`Returns "Bad Request" when given invalid protobuf`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.POST("/cars", CreateCar(&mockClient))
			protobufBody := "}{"
			request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(string(protobufBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/protobuf")

			router.ServeHTTP(respRecorder, request)
			expectedResponseStruct := &proto.SimpleError{
				Error: `invalid protobuf: unexpected EOF`,
			}
			expectedResponse, err := protobuf.Marshal(expectedResponseStruct)
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.CreateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	Context("When given accept xml header", func() {
		It("Creates a new car successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.POST("/cars", CreateCar(&mockClient))
			reqBody := proto.CarRequest{
				Make: "test-make",
				Model: "test-model",
				Vin: "test-vin",
			}
			xmlBody, err := xml.Marshal(reqBody)
			Expect(err).ToNot(HaveOccurred())
			request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(string(xmlBody)))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/xml")

			router.ServeHTTP(respRecorder, request)
			Expect(respRecorder.Body.String()).To(Equal(""))
			Expect(respRecorder.Code).To(Equal(http.StatusCreated))

			Expect(mockClient.CreateMethod.Called).To(BeTrue())
			Expect(mockClient.CreateMethod.GivenReq.GetMake()).To(Equal("test-make"))
			Expect(mockClient.CreateMethod.GivenReq.GetModel()).To(Equal("test-model"))
			Expect(mockClient.CreateMethod.GivenReq.GetVin()).To(Equal("test-vin"))
		})
		It(`Returns "Bad Request" when given invalid xml`, func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{}
			router.POST("/cars", CreateCar(&mockClient))
			xmlBody := "}{"
			request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(xmlBody))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/xml")

			router.ServeHTTP(respRecorder, request)
			expectedResponse := "<map><error>invalid xml: EOF</error></map>"
			Expect(err).ToNot(HaveOccurred())

			Expect(mockClient.CreateMethod.Called).To(BeFalse())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusBadRequest))
		})
	})
	It(`Returns the client status and error when client.Create fails`, func() {
		respRecorder := httptest.NewRecorder()
		_, router := gin.CreateTestContext(respRecorder)
		clientError := errors.New("Client Error")
		mockClient := test.MockClient{
			CreateMethod: test.CreateMethodStruct{
				ReturnError: clientError,
				ReturnSimpleResponse: &proto.SimpleResponse{
					Status: http.StatusServiceUnavailable,
				},
			},
		}
		router.POST("/cars", CreateCar(&mockClient))
		jsonBody := "{}"
		request, err := http.NewRequest("POST", "/cars", bytes.NewBufferString(jsonBody))
		Expect(err).ToNot(HaveOccurred())
		request.Header.Add("Accept", "application/json")

		router.ServeHTTP(respRecorder, request)
		expectedResponseMap := map[string]string{
			"error": fmt.Sprintf(`grcp client: %s`, clientError.Error()),
		}
		expectedResponse, err := json.Marshal(expectedResponseMap)
		Expect(err).ToNot(HaveOccurred())

		Expect(mockClient.CreateMethod.Called).To(BeTrue())
		Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
		Expect(respRecorder.Code).To(Equal(http.StatusServiceUnavailable))
	})
})
