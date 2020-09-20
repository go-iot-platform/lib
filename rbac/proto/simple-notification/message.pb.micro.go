// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: simple-notification/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
	calling "github.com/go-iot-platform/rbac/proto/calling"
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

// Client API for NotificationSvc service

type NotificationSvcService interface {
	// send notification request to call & sms
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*calling.Response, error)
	CreateWithNotify(ctx context.Context, in *NotifyRequest, opts ...client.CallOption) (*calling.Response, error)
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...client.CallOption) (*GetNotificationResponse, error)
	GetNotificationAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type notificationSvcService struct {
	c    client.Client
	name string
}

func NewNotificationSvcService(name string, c client.Client) NotificationSvcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "notificationsvc"
	}
	return &notificationSvcService{
		c:    c,
		name: name,
	}
}

func (c *notificationSvcService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*calling.Response, error) {
	req := c.c.NewRequest(c.name, "NotificationSvc.Create", in)
	out := new(calling.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationSvcService) CreateWithNotify(ctx context.Context, in *NotifyRequest, opts ...client.CallOption) (*calling.Response, error) {
	req := c.c.NewRequest(c.name, "NotificationSvc.CreateWithNotify", in)
	out := new(calling.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationSvcService) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...client.CallOption) (*GetNotificationResponse, error) {
	req := c.c.NewRequest(c.name, "NotificationSvc.GetNotification", in)
	out := new(GetNotificationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationSvcService) GetNotificationAPI(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "NotificationSvc.GetNotificationAPI", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotificationSvc service

type NotificationSvcHandler interface {
	// send notification request to call & sms
	Create(context.Context, *Request, *calling.Response) error
	CreateWithNotify(context.Context, *NotifyRequest, *calling.Response) error
	GetNotification(context.Context, *GetNotificationRequest, *GetNotificationResponse) error
	GetNotificationAPI(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterNotificationSvcHandler(s server.Server, hdlr NotificationSvcHandler, opts ...server.HandlerOption) error {
	type notificationSvc interface {
		Create(ctx context.Context, in *Request, out *calling.Response) error
		CreateWithNotify(ctx context.Context, in *NotifyRequest, out *calling.Response) error
		GetNotification(ctx context.Context, in *GetNotificationRequest, out *GetNotificationResponse) error
		GetNotificationAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type NotificationSvc struct {
		notificationSvc
	}
	h := &notificationSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&NotificationSvc{h}, opts...))
}

type notificationSvcHandler struct {
	NotificationSvcHandler
}

func (h *notificationSvcHandler) Create(ctx context.Context, in *Request, out *calling.Response) error {
	return h.NotificationSvcHandler.Create(ctx, in, out)
}

func (h *notificationSvcHandler) CreateWithNotify(ctx context.Context, in *NotifyRequest, out *calling.Response) error {
	return h.NotificationSvcHandler.CreateWithNotify(ctx, in, out)
}

func (h *notificationSvcHandler) GetNotification(ctx context.Context, in *GetNotificationRequest, out *GetNotificationResponse) error {
	return h.NotificationSvcHandler.GetNotification(ctx, in, out)
}

func (h *notificationSvcHandler) GetNotificationAPI(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.NotificationSvcHandler.GetNotificationAPI(ctx, in, out)
}
