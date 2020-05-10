package test

import (
	"context"
	"github.com/keenanhoffman/cars-api/proto"
	grcp "google.golang.org/grpc"
)

type MockClient struct {
	CreateMethod CreateMethodStruct
}

type CreateMethodStruct struct {
	GivenCtx             context.Context
	GivenReq             *proto.CarRequest
	Called               bool
	ReturnError          error
	ReturnSimpleResponse *proto.SimpleResponse
}
func (m *MockClient) Create(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.SimpleResponse, error) {
	m.CreateMethod.Called = true
	m.CreateMethod.GivenCtx = ctx
	m.CreateMethod.GivenReq = req
	return m.CreateMethod.ReturnSimpleResponse, m.CreateMethod.ReturnError
}

func (m *MockClient) GetById(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.CarResponse, error) {
	return nil, nil
}

func (m *MockClient) GetAll(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.CarsResponse, error) {
	return nil, nil
}

func (m *MockClient) Replace(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.SimpleResponse, error) {
	return nil, nil
}

func (m *MockClient) Update(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.SimpleResponse, error) {
	return nil, nil
}

func (m *MockClient) Delete(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.SimpleResponse, error) {
	return nil, nil
}
