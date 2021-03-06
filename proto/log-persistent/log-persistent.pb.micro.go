// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: log-persistent/log-persistent.proto

package logpersistent

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

// Api Endpoints for LogPersistentSvc service

func NewLogPersistentSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LogPersistentSvc service

type LogPersistentSvcService interface {
	Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error)
}

type logPersistentSvcService struct {
	c    client.Client
	name string
}

func NewLogPersistentSvcService(name string, c client.Client) LogPersistentSvcService {
	return &logPersistentSvcService{
		c:    c,
		name: name,
	}
}

func (c *logPersistentSvcService) Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error) {
	req := c.c.NewRequest(c.name, "LogPersistentSvc.Log", in)
	out := new(DataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogPersistentSvc service

type LogPersistentSvcHandler interface {
	Log(context.Context, *DataRequest, *DataResponse) error
}

func RegisterLogPersistentSvcHandler(s server.Server, hdlr LogPersistentSvcHandler, opts ...server.HandlerOption) error {
	type logPersistentSvc interface {
		Log(ctx context.Context, in *DataRequest, out *DataResponse) error
	}
	type LogPersistentSvc struct {
		logPersistentSvc
	}
	h := &logPersistentSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&LogPersistentSvc{h}, opts...))
}

type logPersistentSvcHandler struct {
	LogPersistentSvcHandler
}

func (h *logPersistentSvcHandler) Log(ctx context.Context, in *DataRequest, out *DataResponse) error {
	return h.LogPersistentSvcHandler.Log(ctx, in, out)
}
