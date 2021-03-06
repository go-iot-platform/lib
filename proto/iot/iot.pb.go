// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: iot/iot.proto

package iot

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ThingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *ThingRequest) Reset() {
	*x = ThingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThingRequest) ProtoMessage() {}

func (x *ThingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThingRequest.ProtoReflect.Descriptor instead.
func (*ThingRequest) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{0}
}

func (x *ThingRequest) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

type ThingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string              `protobuf:"bytes,1,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Serial      string              `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Properties  []*PropertyResponse `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
	Name        string              `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Template    *Template           `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *ThingResponse) Reset() {
	*x = ThingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThingResponse) ProtoMessage() {}

func (x *ThingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThingResponse.ProtoReflect.Descriptor instead.
func (*ThingResponse) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{1}
}

func (x *ThingResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ThingResponse) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ThingResponse) GetProperties() []*PropertyResponse {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ThingResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ThingResponse) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type PropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PropertyResponse) Reset() {
	*x = PropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyResponse) ProtoMessage() {}

func (x *PropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyResponse.ProtoReflect.Descriptor instead.
func (*PropertyResponse) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{2}
}

func (x *PropertyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PropertyResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{3}
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerNumber   string       `protobuf:"bytes,1,opt,name=customerNumber,proto3" json:"customerNumber,omitempty"`
	ThingName        string       `protobuf:"bytes,2,opt,name=thingName,proto3" json:"thingName,omitempty"`
	ThingDisplayName string       `protobuf:"bytes,3,opt,name=thingDisplayName,proto3" json:"thingDisplayName,omitempty"`
	ThingSerial      string       `protobuf:"bytes,4,opt,name=thingSerial,proto3" json:"thingSerial,omitempty"`
	PropertyName     string       `protobuf:"bytes,5,opt,name=propertyName,proto3" json:"propertyName,omitempty"`
	DataType         int32        `protobuf:"varint,6,opt,name=dataType,proto3" json:"dataType,omitempty"`
	Value            string       `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	PropertyId       int32        `protobuf:"varint,8,opt,name=propertyId,proto3" json:"propertyId,omitempty"`
	ThingId          int32        `protobuf:"varint,9,opt,name=thingId,proto3" json:"thingId,omitempty"`
	Alerts           []*AlertData `protobuf:"bytes,10,rep,name=alerts,proto3" json:"alerts,omitempty"`
	TemplateId       int32        `protobuf:"varint,11,opt,name=templateId,proto3" json:"templateId,omitempty"`
	TemplateName     string       `protobuf:"bytes,12,opt,name=templateName,proto3" json:"templateName,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{4}
}

func (x *DataRequest) GetCustomerNumber() string {
	if x != nil {
		return x.CustomerNumber
	}
	return ""
}

func (x *DataRequest) GetThingName() string {
	if x != nil {
		return x.ThingName
	}
	return ""
}

func (x *DataRequest) GetThingDisplayName() string {
	if x != nil {
		return x.ThingDisplayName
	}
	return ""
}

func (x *DataRequest) GetThingSerial() string {
	if x != nil {
		return x.ThingSerial
	}
	return ""
}

func (x *DataRequest) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *DataRequest) GetDataType() int32 {
	if x != nil {
		return x.DataType
	}
	return 0
}

func (x *DataRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *DataRequest) GetPropertyId() int32 {
	if x != nil {
		return x.PropertyId
	}
	return 0
}

func (x *DataRequest) GetThingId() int32 {
	if x != nil {
		return x.ThingId
	}
	return 0
}

func (x *DataRequest) GetAlerts() []*AlertData {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *DataRequest) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *DataRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

type AlertData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value        string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	SettingValue string     `protobuf:"bytes,3,opt,name=settingValue,proto3" json:"settingValue,omitempty"`
	AlertModule  string     `protobuf:"bytes,4,opt,name=alertModule,proto3" json:"alertModule,omitempty"`
	Status       int32      `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	PriorityId   int32      `protobuf:"varint,6,opt,name=priorityId,proto3" json:"priorityId,omitempty"`
	Triggers     []*Trigger `protobuf:"bytes,7,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *AlertData) Reset() {
	*x = AlertData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertData) ProtoMessage() {}

func (x *AlertData) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertData.ProtoReflect.Descriptor instead.
func (*AlertData) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{5}
}

func (x *AlertData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AlertData) GetSettingValue() string {
	if x != nil {
		return x.SettingValue
	}
	return ""
}

func (x *AlertData) GetAlertModule() string {
	if x != nil {
		return x.AlertModule
	}
	return ""
}

func (x *AlertData) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AlertData) GetPriorityId() int32 {
	if x != nil {
		return x.PriorityId
	}
	return 0
}

func (x *AlertData) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Micro     string `protobuf:"bytes,1,opt,name=micro,proto3" json:"micro,omitempty"`
	Endpoint  string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Body      string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Params    string `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	Namespace string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{6}
}

func (x *Trigger) GetMicro() string {
	if x != nil {
		return x.Micro
	}
	return ""
}

func (x *Trigger) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Trigger) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Trigger) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *Trigger) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//repeated Content listContent = 1;
	StatusCode int64 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_iot_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iot_iot_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_iot_iot_proto_rawDescGZIP(), []int{7}
}

func (x *DataResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_iot_iot_proto protoreflect.FileDescriptor

var file_iot_iot_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6f, 0x74, 0x2f, 0x69, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f,
	0x74, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0xd9, 0x01, 0x0a, 0x0d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x42, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69,
	0x6f, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xaa, 0x03, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xea, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x85, 0x01,
	0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xfb, 0x01, 0x0a, 0x06, 0x49, 0x6f, 0x74, 0x53, 0x76, 0x63,
	0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f,
	0x74, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f,
	0x74, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f, 0x74,
	0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x69, 0x6f, 0x74,
	0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6f, 0x74, 0x3b,
	0x69, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iot_iot_proto_rawDescOnce sync.Once
	file_iot_iot_proto_rawDescData = file_iot_iot_proto_rawDesc
)

func file_iot_iot_proto_rawDescGZIP() []byte {
	file_iot_iot_proto_rawDescOnce.Do(func() {
		file_iot_iot_proto_rawDescData = protoimpl.X.CompressGZIP(file_iot_iot_proto_rawDescData)
	})
	return file_iot_iot_proto_rawDescData
}

var file_iot_iot_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_iot_iot_proto_goTypes = []interface{}{
	(*ThingRequest)(nil),     // 0: go.micro.srv.iot.ThingRequest
	(*ThingResponse)(nil),    // 1: go.micro.srv.iot.ThingResponse
	(*PropertyResponse)(nil), // 2: go.micro.srv.iot.PropertyResponse
	(*Template)(nil),         // 3: go.micro.srv.iot.Template
	(*DataRequest)(nil),      // 4: go.micro.srv.iot.DataRequest
	(*AlertData)(nil),        // 5: go.micro.srv.iot.AlertData
	(*Trigger)(nil),          // 6: go.micro.srv.iot.Trigger
	(*DataResponse)(nil),     // 7: go.micro.srv.iot.DataResponse
}
var file_iot_iot_proto_depIdxs = []int32{
	2, // 0: go.micro.srv.iot.ThingResponse.properties:type_name -> go.micro.srv.iot.PropertyResponse
	3, // 1: go.micro.srv.iot.ThingResponse.template:type_name -> go.micro.srv.iot.Template
	5, // 2: go.micro.srv.iot.DataRequest.alerts:type_name -> go.micro.srv.iot.AlertData
	6, // 3: go.micro.srv.iot.AlertData.triggers:type_name -> go.micro.srv.iot.Trigger
	0, // 4: go.micro.srv.iot.IotSvc.GetGateway:input_type -> go.micro.srv.iot.ThingRequest
	0, // 5: go.micro.srv.iot.IotSvc.GetThing:input_type -> go.micro.srv.iot.ThingRequest
	4, // 6: go.micro.srv.iot.IotSvc.ExcuteTemplate:input_type -> go.micro.srv.iot.DataRequest
	1, // 7: go.micro.srv.iot.IotSvc.GetGateway:output_type -> go.micro.srv.iot.ThingResponse
	1, // 8: go.micro.srv.iot.IotSvc.GetThing:output_type -> go.micro.srv.iot.ThingResponse
	7, // 9: go.micro.srv.iot.IotSvc.ExcuteTemplate:output_type -> go.micro.srv.iot.DataResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_iot_iot_proto_init() }
func file_iot_iot_proto_init() {
	if File_iot_iot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iot_iot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iot_iot_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iot_iot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_iot_proto_goTypes,
		DependencyIndexes: file_iot_iot_proto_depIdxs,
		MessageInfos:      file_iot_iot_proto_msgTypes,
	}.Build()
	File_iot_iot_proto = out.File
	file_iot_iot_proto_rawDesc = nil
	file_iot_iot_proto_goTypes = nil
	file_iot_iot_proto_depIdxs = nil
}
