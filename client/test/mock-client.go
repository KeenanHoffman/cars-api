package test

import (
	"context"
	"github.com/keenanhoffman/cars-api/proto"
	grcp "google.golang.org/grpc"
)

type MockClient struct {
	CreateMethod CreateMethodStruct
	GetByIdMethod GetByIdMethodStruct
	GetAllMethod GetAllMethodStruct
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

type GetByIdMethodStruct struct {
	GivenCtx          context.Context
	GivenReq          *proto.CarRequest
	Called            bool
	ReturnError       error
	ReturnCarResponse *proto.CarResponse
}
func (m *MockClient) GetById(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.CarResponse, error) {
	m.GetByIdMethod.Called = true
	m.GetByIdMethod.GivenCtx = ctx
	m.GetByIdMethod.GivenReq = req
	return m.GetByIdMethod.ReturnCarResponse, m.GetByIdMethod.ReturnError
}

type GetAllMethodStruct struct {
	GivenCtx          context.Context
	GivenReq          *proto.CarRequest
	Called            bool
	ReturnError       error
	ReturnCarsResponse *proto.CarsResponse
}
func (m *MockClient) GetAll(ctx context.Context, req *proto.CarRequest, options ...grcp.CallOption) (*proto.CarsResponse, error) {
	m.GetAllMethod.Called = true
	m.GetAllMethod.GivenCtx = ctx
	m.GetAllMethod.GivenReq = req
	return m.GetAllMethod.ReturnCarsResponse, m.GetAllMethod.ReturnError
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
