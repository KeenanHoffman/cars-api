package routes_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/keenanhoffman/cars-api/client/test"
	"github.com/keenanhoffman/cars-api/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	. "github.com/keenanhoffman/cars-api/client/routes"
)

var _ = Describe("GetCars", func() {
	Context("When given accept json header", func() {
		It("Gets cars successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{
				GetAllMethod: test.GetAllMethodStruct{
					ReturnCarsResponse: &proto.CarsResponse{
						Status: http.StatusOK,
						Cars: []*proto.Car{
							{
								Id:    12345,
								Make:  "test-make",
								Model: "test-model",
								Vin:   "test-vin",
							},
						},
					},
				},
			}
			router.GET("/cars", GetCars(&mockClient))
			request, err := http.NewRequest("GET", "/cars", bytes.NewBufferString(""))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/json")

			router.ServeHTTP(respRecorder, request)

			expetedResponseMap := []map[string]interface{}{
				{
					"id":    12345,
					"make":  "test-make",
					"model": "test-model",
					"vin":   "test-vin",
				},
			}
			expectedResponse, err := json.Marshal(expetedResponseMap)
			Expect(err).ToNot(HaveOccurred())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.GetAllMethod.Called).To(BeTrue())
		})
	})
	Context("When given accept xml header", func() {
		It("Gets cars successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{
				GetAllMethod: test.GetAllMethodStruct{
					ReturnCarsResponse: &proto.CarsResponse{
						Status: http.StatusOK,
						Cars: []*proto.Car{
							{
								Id:    12345,
								Make:  "test-make",
								Model: "test-model",
								Vin:   "test-vin",
							},
						},
					},
				},
			}
			router.GET("/cars", GetCars(&mockClient))
			request, err := http.NewRequest("GET", "/cars", bytes.NewBufferString(""))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/xml")

			router.ServeHTTP(respRecorder, request)

			expectedResponseCar := []*proto.Car{
				{
					Id:    12345,
					Make:  "test-make",
					Model: "test-model",
					Vin:   "test-vin",
				},
			}
			expectedResponse, err := xml.Marshal(expectedResponseCar)
			Expect(err).ToNot(HaveOccurred())
			Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.GetAllMethod.Called).To(BeTrue())
		})
	})
	Context("When given accept protobuf header", func() {
		It("Gets a car given an id successfully", func() {
			respRecorder := httptest.NewRecorder()
			_, router := gin.CreateTestContext(respRecorder)
			mockClient := test.MockClient{
				GetAllMethod: test.GetAllMethodStruct{
					ReturnCarsResponse: &proto.CarsResponse{
						Status: http.StatusOK,
						Cars: []*proto.Car{
							{
								Id:    12345,
								Make:  "test-make",
								Model: "test-model",
								Vin:   "test-vin",
							},
						},
					},
				},
			}
			router.GET("/cars", GetCars(&mockClient))
			request, err := http.NewRequest("GET", "/cars", bytes.NewBufferString(""))
			Expect(err).ToNot(HaveOccurred())
			request.Header.Add("Accept", "application/protobuf")

			router.ServeHTTP(respRecorder, request)

			//expectedResponseCars := &proto.CarsResponse{
			//	Cars: []*proto.Car{
			//		{
			//			Id:    12345,
			//			Make:  "test-make",
			//			Model: "test-model",
			//			Vin:   "test-vin",
			//		},
			//
			//	},
			//}
			//expectedResponse, err := protobuf.Marshal(expectedResponseCars)
			//Expect(err).ToNot(HaveOccurred())
			carsResponse := &proto.CarsResponse{}
			err = protobuf.Unmarshal(respRecorder.Body.Bytes(), carsResponse)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(carsResponse.GetCars())).To(Equal(1))
			Expect(carsResponse.GetCars()[0].GetId()).To(Equal(int64(12345)))
			Expect(carsResponse.GetCars()[0].GetMake()).To(Equal("test-make"))
			Expect(carsResponse.GetCars()[0].GetModel()).To(Equal("test-model"))
			Expect(carsResponse.GetCars()[0].GetVin()).To(Equal("test-vin"))
			Expect(respRecorder.Code).To(Equal(http.StatusOK))

			Expect(mockClient.GetAllMethod.Called).To(BeTrue())
		})
	})
	//It(`Returns the client status and error when client.GetById fails`, func() {
	//	respRecorder := httptest.NewRecorder()
	//	_, router := gin.CreateTestContext(respRecorder)
	//	clientError := errors.New("Client Error")
	//	mockClient := test.MockClient{
	//		GetByIdMethod: test.GetByIdMethodStruct{
	//			ReturnCarResponse: &proto.CarResponse{
	//				Status: http.StatusServiceUnavailable,
	//			},
	//			ReturnError: clientError,
	//		},
	//	}
	//	router.GET("/cars/:id", GetCarById(&mockClient))
	//	jsonBody := "{}"
	//	request, err := http.NewRequest("GET", "/cars/12345", bytes.NewBufferString(jsonBody))
	//	Expect(err).ToNot(HaveOccurred())
	//	request.Header.Add("Content-type", "application/json")
	//
	//	router.ServeHTTP(respRecorder, request)
	//	expectedResponseMap := map[string]string{
	//		"error": fmt.Sprintf(`grcp client: %s`, clientError.Error()),
	//	}
	//	expectedResponse, err := json.Marshal(expectedResponseMap)
	//	Expect(err).ToNot(HaveOccurred())
	//
	//	Expect(mockClient.GetByIdMethod.Called).To(BeTrue())
	//	Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
	//	Expect(respRecorder.Code).To(Equal(http.StatusServiceUnavailable))
	//})
})
