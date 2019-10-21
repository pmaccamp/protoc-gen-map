// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/gentest/unary_type_test.proto

package gentest

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{1}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type EMpTyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EMpTyResponse) Reset()         { *m = EMpTyResponse{} }
func (m *EMpTyResponse) String() string { return proto.CompactTextString(m) }
func (*EMpTyResponse) ProtoMessage()    {}
func (*EMpTyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{2}
}

func (m *EMpTyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EMpTyResponse.Unmarshal(m, b)
}
func (m *EMpTyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EMpTyResponse.Marshal(b, m, deterministic)
}
func (m *EMpTyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EMpTyResponse.Merge(m, src)
}
func (m *EMpTyResponse) XXX_Size() int {
	return xxx_messageInfo_EMpTyResponse.Size(m)
}
func (m *EMpTyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EMpTyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EMpTyResponse proto.InternalMessageInfo

type NilResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NilResponse) Reset()         { *m = NilResponse{} }
func (m *NilResponse) String() string { return proto.CompactTextString(m) }
func (*NilResponse) ProtoMessage()    {}
func (*NilResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{3}
}

func (m *NilResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NilResponse.Unmarshal(m, b)
}
func (m *NilResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NilResponse.Marshal(b, m, deterministic)
}
func (m *NilResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NilResponse.Merge(m, src)
}
func (m *NilResponse) XXX_Size() int {
	return xxx_messageInfo_NilResponse.Size(m)
}
func (m *NilResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NilResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NilResponse proto.InternalMessageInfo

type NullResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullResponse) Reset()         { *m = NullResponse{} }
func (m *NullResponse) String() string { return proto.CompactTextString(m) }
func (*NullResponse) ProtoMessage()    {}
func (*NullResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{4}
}

func (m *NullResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullResponse.Unmarshal(m, b)
}
func (m *NullResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullResponse.Marshal(b, m, deterministic)
}
func (m *NullResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullResponse.Merge(m, src)
}
func (m *NullResponse) XXX_Size() int {
	return xxx_messageInfo_NullResponse.Size(m)
}
func (m *NullResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NullResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NullResponse proto.InternalMessageInfo

type NuLlResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NuLlResponse) Reset()         { *m = NuLlResponse{} }
func (m *NuLlResponse) String() string { return proto.CompactTextString(m) }
func (*NuLlResponse) ProtoMessage()    {}
func (*NuLlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{5}
}

func (m *NuLlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NuLlResponse.Unmarshal(m, b)
}
func (m *NuLlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NuLlResponse.Marshal(b, m, deterministic)
}
func (m *NuLlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NuLlResponse.Merge(m, src)
}
func (m *NuLlResponse) XXX_Size() int {
	return xxx_messageInfo_NuLlResponse.Size(m)
}
func (m *NuLlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NuLlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NuLlResponse proto.InternalMessageInfo

type SampleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleResponse) Reset()         { *m = SampleResponse{} }
func (m *SampleResponse) String() string { return proto.CompactTextString(m) }
func (*SampleResponse) ProtoMessage()    {}
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35369211b23c51f, []int{6}
}

func (m *SampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleResponse.Unmarshal(m, b)
}
func (m *SampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleResponse.Marshal(b, m, deterministic)
}
func (m *SampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleResponse.Merge(m, src)
}
func (m *SampleResponse) XXX_Size() int {
	return xxx_messageInfo_SampleResponse.Size(m)
}
func (m *SampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SampleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "gentest.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "gentest.EmptyResponse")
	proto.RegisterType((*EMpTyResponse)(nil), "gentest.EMpTyResponse")
	proto.RegisterType((*NilResponse)(nil), "gentest.NilResponse")
	proto.RegisterType((*NullResponse)(nil), "gentest.NullResponse")
	proto.RegisterType((*NuLlResponse)(nil), "gentest.nuLlResponse")
	proto.RegisterType((*SampleResponse)(nil), "gentest.SampleResponse")
}

func init() {
	proto.RegisterFile("testdata/gentest/unary_type_test.proto", fileDescriptor_e35369211b23c51f)
}

var fileDescriptor_e35369211b23c51f = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0x77, 0x30, 0x4c, 0xeb, 0xc6, 0x0c, 0xf1, 0x47, 0xe2, 0xd1, 0x83, 0xb7, 0x41, 0xa2,
	0x27, 0xa7, 0x27, 0x75, 0x26, 0x26, 0x3a, 0x93, 0x81, 0x17, 0x2f, 0x4b, 0xc7, 0x5e, 0x18, 0x1b,
	0xb4, 0xb5, 0xb4, 0x53, 0xfe, 0x27, 0xff, 0x48, 0x53, 0xec, 0x60, 0x5d, 0x42, 0x5c, 0x38, 0xc1,
	0xf7, 0x03, 0x1f, 0xfa, 0xde, 0xe3, 0xa1, 0x4b, 0x01, 0x99, 0x98, 0x61, 0x81, 0xbd, 0x08, 0x88,
	0xba, 0xf7, 0x24, 0xc1, 0x3c, 0x9f, 0x88, 0x9c, 0xc1, 0x44, 0x65, 0x97, 0x71, 0x2a, 0xa8, 0xd3,
	0xd6, 0x8f, 0x2f, 0x6c, 0xd4, 0x19, 0xa6, 0x4c, 0xe4, 0x63, 0xf8, 0x94, 0x2a, 0xf7, 0x50, 0x57,
	0xe7, 0x8c, 0x51, 0x92, 0x41, 0x01, 0x5e, 0x59, 0x50, 0x81, 0x2e, 0x3a, 0x1c, 0xc5, 0x49, 0x19,
	0x6d, 0xd4, 0x19, 0xc9, 0xc4, 0xc8, 0x44, 0xbe, 0x54, 0xf9, 0x08, 0xd9, 0x3e, 0x4e, 0x59, 0x02,
	0x6b, 0x72, 0xf5, 0xb3, 0x87, 0x7a, 0xc3, 0x6f, 0x08, 0x83, 0x9c, 0x81, 0x0f, 0x7c, 0x15, 0x87,
	0xe0, 0x0c, 0x50, 0x5b, 0xa1, 0x37, 0x02, 0xce, 0x89, 0xab, 0x6b, 0x73, 0x37, 0x0b, 0x3b, 0x3f,
	0xdd, 0xc6, 0xfa, 0xfb, 0xad, 0xb5, 0x1b, 0x7c, 0xd1, 0x1d, 0x5c, 0xa3, 0x15, 0xe5, 0x1e, 0x14,
	0xee, 0x9c, 0x43, 0xed, 0xc9, 0xc7, 0x25, 0xde, 0xec, 0x5b, 0xb9, 0xfb, 0xca, 0x7d, 0xa2, 0x92,
	0xd7, 0xa9, 0x15, 0x36, 0x66, 0x54, 0xb9, 0xf1, 0x0a, 0xfe, 0x77, 0x8d, 0x79, 0xb6, 0x9c, 0x1b,
	0x64, 0x3d, 0x13, 0x1f, 0xc6, 0xa2, 0xc9, 0xa8, 0xac, 0x47, 0x48, 0x40, 0xd4, 0x1e, 0x7a, 0x56,
	0x62, 0xf3, 0xa7, 0xfd, 0xb9, 0xef, 0x6c, 0x86, 0x9b, 0xba, 0x0f, 0x1c, 0x1a, 0xb9, 0xf7, 0x77,
	0x1f, 0x83, 0x28, 0x16, 0x73, 0x39, 0x75, 0x43, 0x9a, 0x7a, 0x0b, 0x1c, 0x2e, 0xb3, 0xe5, 0xc2,
	0x2b, 0xd6, 0x38, 0xec, 0x47, 0x40, 0xfa, 0x29, 0x66, 0xde, 0xf6, 0xda, 0xdf, 0xea, 0xeb, 0xd4,
	0x2a, 0x5e, 0xbc, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xec, 0x3e, 0x88, 0x19, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecTypeServiceClient is the client API for ExecTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecTypeServiceClient interface {
	ExecOne(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	ExecTwo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EMpTyResponse, error)
	ExecThree(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NilResponse, error)
	ExecFour(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NullResponse, error)
	ExecFive(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NuLlResponse, error)
	InSeRt(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Delete(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error)
	Update(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error)
	Create(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error)
}

type execTypeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecTypeServiceClient(cc *grpc.ClientConn) ExecTypeServiceClient {
	return &execTypeServiceClient{cc}
}

func (c *execTypeServiceClient) ExecOne(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/ExecOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) ExecTwo(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EMpTyResponse, error) {
	out := new(EMpTyResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/ExecTwo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) ExecThree(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NilResponse, error) {
	out := new(NilResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/ExecThree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) ExecFour(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NullResponse, error) {
	out := new(NullResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/ExecFour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) ExecFive(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*NuLlResponse, error) {
	out := new(NuLlResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/ExecFive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) InSeRt(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/InSeRt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) Delete(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) Update(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *execTypeServiceClient) Create(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/gentest.ExecTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecTypeServiceServer is the server API for ExecTypeService service.
type ExecTypeServiceServer interface {
	ExecOne(context.Context, *EmptyRequest) (*EmptyResponse, error)
	ExecTwo(context.Context, *EmptyRequest) (*EMpTyResponse, error)
	ExecThree(context.Context, *EmptyRequest) (*NilResponse, error)
	ExecFour(context.Context, *EmptyRequest) (*NullResponse, error)
	ExecFive(context.Context, *EmptyRequest) (*NuLlResponse, error)
	InSeRt(context.Context, *EmptyRequest) (*EmptyResponse, error)
	Delete(context.Context, *EmptyRequest) (*SampleResponse, error)
	Update(context.Context, *EmptyRequest) (*SampleResponse, error)
	Create(context.Context, *EmptyRequest) (*SampleResponse, error)
}

// UnimplementedExecTypeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExecTypeServiceServer struct {
}

func (*UnimplementedExecTypeServiceServer) ExecOne(ctx context.Context, req *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecOne not implemented")
}
func (*UnimplementedExecTypeServiceServer) ExecTwo(ctx context.Context, req *EmptyRequest) (*EMpTyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecTwo not implemented")
}
func (*UnimplementedExecTypeServiceServer) ExecThree(ctx context.Context, req *EmptyRequest) (*NilResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecThree not implemented")
}
func (*UnimplementedExecTypeServiceServer) ExecFour(ctx context.Context, req *EmptyRequest) (*NullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecFour not implemented")
}
func (*UnimplementedExecTypeServiceServer) ExecFive(ctx context.Context, req *EmptyRequest) (*NuLlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecFive not implemented")
}
func (*UnimplementedExecTypeServiceServer) InSeRt(ctx context.Context, req *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InSeRt not implemented")
}
func (*UnimplementedExecTypeServiceServer) Delete(ctx context.Context, req *EmptyRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedExecTypeServiceServer) Update(ctx context.Context, req *EmptyRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedExecTypeServiceServer) Create(ctx context.Context, req *EmptyRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterExecTypeServiceServer(s *grpc.Server, srv ExecTypeServiceServer) {
	s.RegisterService(&_ExecTypeService_serviceDesc, srv)
}

func _ExecTypeService_ExecOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).ExecOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/ExecOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).ExecOne(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_ExecTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).ExecTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/ExecTwo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).ExecTwo(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_ExecThree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).ExecThree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/ExecThree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).ExecThree(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_ExecFour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).ExecFour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/ExecFour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).ExecFour(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_ExecFive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).ExecFive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/ExecFive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).ExecFive(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_InSeRt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).InSeRt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/InSeRt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).InSeRt(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).Delete(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).Update(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gentest.ExecTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecTypeServiceServer).Create(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecTypeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gentest.ExecTypeService",
	HandlerType: (*ExecTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecOne",
			Handler:    _ExecTypeService_ExecOne_Handler,
		},
		{
			MethodName: "ExecTwo",
			Handler:    _ExecTypeService_ExecTwo_Handler,
		},
		{
			MethodName: "ExecThree",
			Handler:    _ExecTypeService_ExecThree_Handler,
		},
		{
			MethodName: "ExecFour",
			Handler:    _ExecTypeService_ExecFour_Handler,
		},
		{
			MethodName: "ExecFive",
			Handler:    _ExecTypeService_ExecFive_Handler,
		},
		{
			MethodName: "InSeRt",
			Handler:    _ExecTypeService_InSeRt_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ExecTypeService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ExecTypeService_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ExecTypeService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/gentest/unary_type_test.proto",
}
