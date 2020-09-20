// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: synchronize/message.proto

package synchronize

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SynchronizeSvc service

type SynchronizeSvcService interface {
	// customer info
	Register(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ChangeGateway(ctx context.Context, in *ChangeGatewayRequest, opts ...client.CallOption) (*Response, error)
	InactiveDevice(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	RemoveDevice(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ZoneSync(ctx context.Context, in *ZoneRequest, opts ...client.CallOption) (*Response, error)
	SceneSync(ctx context.Context, in *SceneRequest, opts ...client.CallOption) (*Response, error)
	ScheduleSync(ctx context.Context, in *ScheduleRequest, opts ...client.CallOption) (*Response, error)
}

type synchronizeSvcService struct {
	c    client.Client
	name string
}

func NewSynchronizeSvcService(name string, c client.Client) SynchronizeSvcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "synchronize"
	}
	return &synchronizeSvcService{
		c:    c,
		name: name,
	}
}

func (c *synchronizeSvcService) Register(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) ChangeGateway(ctx context.Context, in *ChangeGatewayRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.ChangeGateway", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) InactiveDevice(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.InactiveDevice", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) RemoveDevice(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.RemoveDevice", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) ZoneSync(ctx context.Context, in *ZoneRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.ZoneSync", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) SceneSync(ctx context.Context, in *SceneRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.SceneSync", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizeSvcService) ScheduleSync(ctx context.Context, in *ScheduleRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SynchronizeSvc.ScheduleSync", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SynchronizeSvc service

type SynchronizeSvcHandler interface {
	// customer info
	Register(context.Context, *Request, *Response) error
	ChangeGateway(context.Context, *ChangeGatewayRequest, *Response) error
	InactiveDevice(context.Context, *Request, *Response) error
	RemoveDevice(context.Context, *Request, *Response) error
	ZoneSync(context.Context, *ZoneRequest, *Response) error
	SceneSync(context.Context, *SceneRequest, *Response) error
	ScheduleSync(context.Context, *ScheduleRequest, *Response) error
}

func RegisterSynchronizeSvcHandler(s server.Server, hdlr SynchronizeSvcHandler, opts ...server.HandlerOption) error {
	type synchronizeSvc interface {
		Register(ctx context.Context, in *Request, out *Response) error
		ChangeGateway(ctx context.Context, in *ChangeGatewayRequest, out *Response) error
		InactiveDevice(ctx context.Context, in *Request, out *Response) error
		RemoveDevice(ctx context.Context, in *Request, out *Response) error
		ZoneSync(ctx context.Context, in *ZoneRequest, out *Response) error
		SceneSync(ctx context.Context, in *SceneRequest, out *Response) error
		ScheduleSync(ctx context.Context, in *ScheduleRequest, out *Response) error
	}
	type SynchronizeSvc struct {
		synchronizeSvc
	}
	h := &synchronizeSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&SynchronizeSvc{h}, opts...))
}

type synchronizeSvcHandler struct {
	SynchronizeSvcHandler
}

func (h *synchronizeSvcHandler) Register(ctx context.Context, in *Request, out *Response) error {
	return h.SynchronizeSvcHandler.Register(ctx, in, out)
}

func (h *synchronizeSvcHandler) ChangeGateway(ctx context.Context, in *ChangeGatewayRequest, out *Response) error {
	return h.SynchronizeSvcHandler.ChangeGateway(ctx, in, out)
}

func (h *synchronizeSvcHandler) InactiveDevice(ctx context.Context, in *Request, out *Response) error {
	return h.SynchronizeSvcHandler.InactiveDevice(ctx, in, out)
}

func (h *synchronizeSvcHandler) RemoveDevice(ctx context.Context, in *Request, out *Response) error {
	return h.SynchronizeSvcHandler.RemoveDevice(ctx, in, out)
}

func (h *synchronizeSvcHandler) ZoneSync(ctx context.Context, in *ZoneRequest, out *Response) error {
	return h.SynchronizeSvcHandler.ZoneSync(ctx, in, out)
}

func (h *synchronizeSvcHandler) SceneSync(ctx context.Context, in *SceneRequest, out *Response) error {
	return h.SynchronizeSvcHandler.SceneSync(ctx, in, out)
}

func (h *synchronizeSvcHandler) ScheduleSync(ctx context.Context, in *ScheduleRequest, out *Response) error {
	return h.SynchronizeSvcHandler.ScheduleSync(ctx, in, out)
}
