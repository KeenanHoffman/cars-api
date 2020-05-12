package routes_test

import (
	//"bytes"
	//"encoding/json"
	//"github.com/gin-gonic/gin"
	//"github.com/keenanhoffman/cars-api/client/test"
	//"github.com/keenanhoffman/cars-api/proto"
	. "github.com/onsi/ginkgo"
	//"net/http"
	//"net/http/httptest"
	//rgmock "google.golang.org/grpc/examples/route_guide/mock_routeguide"
	//rgpb "google.golang.org/grpc/examples/route_guide/routeguide"
	//
	//. "github.com/onsi/gomega"
	//
	//. "github.com/keenanhoffman/cars-api/client/routes"
)

var _ = Describe("Search", func() {
	Context("When given accept json header", func() {
		Context("When given content-type json header", func() {
			//It("Searches for cars successfully", func() {
			//	respRecorder := httptest.NewRecorder()
			//	_, router := gin.CreateTestContext(respRecorder)
			//
			//	stream := rgmock.NewMockRouteGuideClient(proto.NewAddCarServiceClient)
			//
			//	router.GET("/cars/:id", GetCarById(&mockClient))
			//	request, err := http.NewRequest("GET", "/cars/12345", bytes.NewBufferString(""))
			//	Expect(err).ToNot(HaveOccurred())
			//	request.Header.Add("Accept", "application/json")
			//
			//	router.ServeHTTP(respRecorder, request)
			//
			//	expetedResponseMap := map[string]interface{}{
			//		"id": 12345,
			//		"make": "test-make",
			//		"model": "test-model",
			//		"vin": "test-vin",
			//	}
			//	expectedResponse, err := json.Marshal(expetedResponseMap)
			//	Expect(err).ToNot(HaveOccurred())
			//	Expect(respRecorder.Body.String()).To(Equal(string(expectedResponse)))
			//	Expect(respRecorder.Code).To(Equal(http.StatusOK))
			//
			//	Expect(mockClient.GetByIdMethod.Called).To(BeTrue())
			//	Expect(mockClient.GetByIdMethod.GivenReq.GetId()).To(Equal(int64(12345)))
			//})
		})
	})
})
