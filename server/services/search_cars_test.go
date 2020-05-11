package services_test

import (
	"context"
	"github.com/keenanhoffman/cars-api/proto"
	"github.com/keenanhoffman/cars-api/server/test"
	. "github.com/keenanhoffman/cars-api/server/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"google.golang.org/grpc/test/bufconn"
)

var _ = Describe("SearchCars", func() {
	BeforeEach(func() {
	})
	It("Searches for cars successfully", func() {
		mockDB := &test.MockDB{
			SearchCarMethod: test.SearchCarMethodStruct{
				ReturnCars: []*proto.Car{
					{
						Id: 12345,
						Make: "test-make",
						Model: "test-model",
						Vin: "test-vin",
					},
					{
						Id: 123456,
						Make: "test-make",
						Model: "other-model",
						Vin: "other-vin",
					},
					{
						Id: 1234567,
						Make: "test-make",
						Model: "another-model",
						Vin: "another-vin",
					},
				},
			},
		}
		server := Services{
			DB: mockDB,
		}
		listener := bufconn.Listen(1024 * 1024)
		s := grpc.NewServer()
		proto.RegisterAddCarServiceServer(s, &server)
		go func() {
			if err := s.Serve(listener); err != nil {
				log.Fatalf("Server exited with error: %v", err)
			}
		}()
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func (context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewAddCarServiceClient(conn)

		req := proto.CarRequest{
			Make:  "test-make",
		}
		stream, err := client.Search(ctx, &req)
		Expect(err).ToNot(HaveOccurred())
		foundCars := []*proto.Car{}
		for {
			car, err := stream.Recv()
			if err == io.EOF {
				break
			} else {
				Expect(err).ToNot(HaveOccurred())
			}
			foundCars = append(foundCars, car)
		}
		Expect(len(foundCars)).To(Equal(3))
		Expect(foundCars[0].GetModel()).To(Equal("test-model"))
		Expect(foundCars[1].GetModel()).To(Equal("other-model"))
		Expect(foundCars[2].GetModel()).To(Equal("another-model"))

		Expect(mockDB.SearchCarMethod.Called).To(BeTrue())
		Expect(mockDB.SearchCarMethod.GivenCar.GetMake()).To(Equal("test-make"))
	})
})
