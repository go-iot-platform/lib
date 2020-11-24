// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: log-serial/log-serial.proto

package logserial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/go-iot-platform/go-micro/api"
	client "github.com/go-iot-platform/go-micro/client"
	server "github.com/go-iot-platform/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for LogSerialSvc service

func NewLogSerialSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LogSerialSvc service

type LogSerialSvcService interface {
	Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error)
	LogHistory(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error)
}

type logSerialSvcService struct {
	c    client.Client
	name string
}

func NewLogSerialSvcService(name string, c client.Client) LogSerialSvcService {
	return &logSerialSvcService{
		c:    c,
		name: name,
	}
}

func (c *logSerialSvcService) Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error) {
	req := c.c.NewRequest(c.name, "LogSerialSvc.Log", in)
	out := new(DataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logSerialSvcService) LogHistory(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error) {
	req := c.c.NewRequest(c.name, "LogSerialSvc.LogHistory", in)
	out := new(DataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogSerialSvc service

type LogSerialSvcHandler interface {
	Log(context.Context, *DataRequest, *DataResponse) error
	LogHistory(context.Context, *DataRequest, *DataResponse) error
}

func RegisterLogSerialSvcHandler(s server.Server, hdlr LogSerialSvcHandler, opts ...server.HandlerOption) error {
	type logSerialSvc interface {
		Log(ctx context.Context, in *DataRequest, out *DataResponse) error
		LogHistory(ctx context.Context, in *DataRequest, out *DataResponse) error
	}
	type LogSerialSvc struct {
		logSerialSvc
	}
	h := &logSerialSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&LogSerialSvc{h}, opts...))
}

type logSerialSvcHandler struct {
	LogSerialSvcHandler
}

func (h *logSerialSvcHandler) Log(ctx context.Context, in *DataRequest, out *DataResponse) error {
	return h.LogSerialSvcHandler.Log(ctx, in, out)
}

func (h *logSerialSvcHandler) LogHistory(ctx context.Context, in *DataRequest, out *DataResponse) error {
	return h.LogSerialSvcHandler.LogHistory(ctx, in, out)
}