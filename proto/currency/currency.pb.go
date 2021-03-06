// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: proto/currency.proto

package currency

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   string `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *RateRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate float32 `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_proto_currency_proto protoreflect.FileDescriptor

var file_proto_currency_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x72, 0x61, 0x74, 0x65, 0x32, 0x6f, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_proto_rawDescOnce sync.Once
	file_proto_currency_proto_rawDescData = file_proto_currency_proto_rawDesc
)

func file_proto_currency_proto_rawDescGZIP() []byte {
	file_proto_currency_proto_rawDescOnce.Do(func() {
		file_proto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_proto_rawDescData)
	})
	return file_proto_currency_proto_rawDescData
}

var file_proto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_currency_proto_goTypes = []interface{}{
	(*RateRequest)(nil),  // 0: RateRequest
	(*RateResponse)(nil), // 1: RateResponse
}
var file_proto_currency_proto_depIdxs = []int32{
	0, // 0: ExchangeRateService.GetRate:input_type -> RateRequest
	0, // 1: ExchangeRateService.StreamRates:input_type -> RateRequest
	1, // 2: ExchangeRateService.GetRate:output_type -> RateResponse
	1, // 3: ExchangeRateService.StreamRates:output_type -> RateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_currency_proto_init() }
func file_proto_currency_proto_init() {
	if File_proto_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_proto_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
			RawDescriptor: file_proto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_proto_goTypes,
		DependencyIndexes: file_proto_currency_proto_depIdxs,
		MessageInfos:      file_proto_currency_proto_msgTypes,
	}.Build()
	File_proto_currency_proto = out.File
	file_proto_currency_proto_rawDesc = nil
	file_proto_currency_proto_goTypes = nil
	file_proto_currency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExchangeRateServiceClient is the client API for ExchangeRateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeRateServiceClient interface {
	//Unary
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	//Server-side streaming
	StreamRates(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (ExchangeRateService_StreamRatesClient, error)
}

type exchangeRateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeRateServiceClient(cc grpc.ClientConnInterface) ExchangeRateServiceClient {
	return &exchangeRateServiceClient{cc}
}

func (c *exchangeRateServiceClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/ExchangeRateService/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeRateServiceClient) StreamRates(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (ExchangeRateService_StreamRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExchangeRateService_serviceDesc.Streams[0], "/ExchangeRateService/StreamRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeRateServiceStreamRatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExchangeRateService_StreamRatesClient interface {
	Recv() (*RateResponse, error)
	grpc.ClientStream
}

type exchangeRateServiceStreamRatesClient struct {
	grpc.ClientStream
}

func (x *exchangeRateServiceStreamRatesClient) Recv() (*RateResponse, error) {
	m := new(RateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeRateServiceServer is the server API for ExchangeRateService service.
type ExchangeRateServiceServer interface {
	//Unary
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	//Server-side streaming
	StreamRates(*RateRequest, ExchangeRateService_StreamRatesServer) error
}

// UnimplementedExchangeRateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeRateServiceServer struct {
}

func (*UnimplementedExchangeRateServiceServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (*UnimplementedExchangeRateServiceServer) StreamRates(*RateRequest, ExchangeRateService_StreamRatesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRates not implemented")
}

func RegisterExchangeRateServiceServer(s *grpc.Server, srv ExchangeRateServiceServer) {
	s.RegisterService(&_ExchangeRateService_serviceDesc, srv)
}

func _ExchangeRateService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeRateServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExchangeRateService/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeRateServiceServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeRateService_StreamRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeRateServiceServer).StreamRates(m, &exchangeRateServiceStreamRatesServer{stream})
}

type ExchangeRateService_StreamRatesServer interface {
	Send(*RateResponse) error
	grpc.ServerStream
}

type exchangeRateServiceStreamRatesServer struct {
	grpc.ServerStream
}

func (x *exchangeRateServiceStreamRatesServer) Send(m *RateResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ExchangeRateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExchangeRateService",
	HandlerType: (*ExchangeRateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _ExchangeRateService_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRates",
			Handler:       _ExchangeRateService_StreamRates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/currency.proto",
}
