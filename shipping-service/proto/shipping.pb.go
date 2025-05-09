// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: proto/shipping.proto

package shippingservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartShippingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartShippingRequest) Reset() {
	*x = StartShippingRequest{}
	mi := &file_proto_shipping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingRequest) ProtoMessage() {}

func (x *StartShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingRequest.ProtoReflect.Descriptor instead.
func (*StartShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *StartShippingRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *StartShippingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type StartShippingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartShippingResponse) Reset() {
	*x = StartShippingResponse{}
	mi := &file_proto_shipping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingResponse) ProtoMessage() {}

func (x *StartShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingResponse.ProtoReflect.Descriptor instead.
func (*StartShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *StartShippingResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *StartShippingResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StartShippingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CancelShippingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelShippingRequest) Reset() {
	*x = CancelShippingRequest{}
	mi := &file_proto_shipping_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingRequest) ProtoMessage() {}

func (x *CancelShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingRequest.ProtoReflect.Descriptor instead.
func (*CancelShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *CancelShippingRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelShippingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelShippingResponse) Reset() {
	*x = CancelShippingResponse{}
	mi := &file_proto_shipping_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingResponse) ProtoMessage() {}

func (x *CancelShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingResponse.ProtoReflect.Descriptor instead.
func (*CancelShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{3}
}

func (x *CancelShippingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CancelShippingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ShippedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShippedRequest) Reset() {
	*x = ShippedRequest{}
	mi := &file_proto_shipping_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShippedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippedRequest) ProtoMessage() {}

func (x *ShippedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippedRequest.ProtoReflect.Descriptor instead.
func (*ShippedRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{4}
}

func (x *ShippedRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type ShippedResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShippedResponse) Reset() {
	*x = ShippedResponse{}
	mi := &file_proto_shipping_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShippedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippedResponse) ProtoMessage() {}

func (x *ShippedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippedResponse.ProtoReflect.Descriptor instead.
func (*ShippedResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{5}
}

func (x *ShippedResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ShippedResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_shipping_proto protoreflect.FileDescriptor

const file_proto_shipping_proto_rawDesc = "" +
	"\n" +
	"\x14proto/shipping.proto\x12\x0fshippingservice\"K\n" +
	"\x14StartShippingRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\"d\n" +
	"\x15StartShippingResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\"2\n" +
	"\x15CancelShippingRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"L\n" +
	"\x16CancelShippingResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"+\n" +
	"\x0eShippedRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"D\n" +
	"\x0fShippedResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status2\xa2\x02\n" +
	"\x0fShippingService\x12^\n" +
	"\rStartShipping\x12%.shippingservice.StartShippingRequest\x1a&.shippingservice.StartShippingResponse\x12a\n" +
	"\x0eCancelShipping\x12&.shippingservice.CancelShippingRequest\x1a'.shippingservice.CancelShippingResponse\x12L\n" +
	"\aShipped\x12\x1f.shippingservice.ShippedRequest\x1a .shippingservice.ShippedResponseB\x19Z\x17proto/;shippingservice;b\x06proto3"

var (
	file_proto_shipping_proto_rawDescOnce sync.Once
	file_proto_shipping_proto_rawDescData []byte
)

func file_proto_shipping_proto_rawDescGZIP() []byte {
	file_proto_shipping_proto_rawDescOnce.Do(func() {
		file_proto_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_shipping_proto_rawDesc), len(file_proto_shipping_proto_rawDesc)))
	})
	return file_proto_shipping_proto_rawDescData
}

var file_proto_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_shipping_proto_goTypes = []any{
	(*StartShippingRequest)(nil),   // 0: shippingservice.StartShippingRequest
	(*StartShippingResponse)(nil),  // 1: shippingservice.StartShippingResponse
	(*CancelShippingRequest)(nil),  // 2: shippingservice.CancelShippingRequest
	(*CancelShippingResponse)(nil), // 3: shippingservice.CancelShippingResponse
	(*ShippedRequest)(nil),         // 4: shippingservice.ShippedRequest
	(*ShippedResponse)(nil),        // 5: shippingservice.ShippedResponse
}
var file_proto_shipping_proto_depIdxs = []int32{
	0, // 0: shippingservice.ShippingService.StartShipping:input_type -> shippingservice.StartShippingRequest
	2, // 1: shippingservice.ShippingService.CancelShipping:input_type -> shippingservice.CancelShippingRequest
	4, // 2: shippingservice.ShippingService.Shipped:input_type -> shippingservice.ShippedRequest
	1, // 3: shippingservice.ShippingService.StartShipping:output_type -> shippingservice.StartShippingResponse
	3, // 4: shippingservice.ShippingService.CancelShipping:output_type -> shippingservice.CancelShippingResponse
	5, // 5: shippingservice.ShippingService.Shipped:output_type -> shippingservice.ShippedResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_shipping_proto_init() }
func file_proto_shipping_proto_init() {
	if File_proto_shipping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_shipping_proto_rawDesc), len(file_proto_shipping_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shipping_proto_goTypes,
		DependencyIndexes: file_proto_shipping_proto_depIdxs,
		MessageInfos:      file_proto_shipping_proto_msgTypes,
	}.Build()
	File_proto_shipping_proto = out.File
	file_proto_shipping_proto_goTypes = nil
	file_proto_shipping_proto_depIdxs = nil
}
