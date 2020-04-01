// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package api_customer_bri_v1_users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// RetrieveRequest is the request for types service.
type RetrieveRequest struct {
	UId                  string   `protobuf:"bytes,1,opt,name=uId,proto3" json:"uId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveRequest) Reset()         { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()    {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *RetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveRequest.Unmarshal(m, b)
}
func (m *RetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveRequest.Merge(m, src)
}
func (m *RetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveRequest.Size(m)
}
func (m *RetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveRequest proto.InternalMessageInfo

func (m *RetrieveRequest) GetUId() string {
	if m != nil {
		return m.UId
	}
	return ""
}

type Users struct {
	UId                  string   `protobuf:"bytes,1,opt,name=uId,proto3" json:"uId,omitempty"`
	UName                string   `protobuf:"bytes,2,opt,name=uName,proto3" json:"uName,omitempty"`
	UAddress             string   `protobuf:"bytes,3,opt,name=uAddress,proto3" json:"uAddress,omitempty"`
	Age                  string   `protobuf:"bytes,4,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUId() string {
	if m != nil {
		return m.UId
	}
	return ""
}

func (m *Users) GetUName() string {
	if m != nil {
		return m.UName
	}
	return ""
}

func (m *Users) GetUAddress() string {
	if m != nil {
		return m.UAddress
	}
	return ""
}

func (m *Users) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

// RetrieveResponse is the response for types service.
type RetrieveResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RespCode             string   `protobuf:"bytes,2,opt,name=respCode,proto3" json:"respCode,omitempty"`
	RespDesc             string   `protobuf:"bytes,3,opt,name=respDesc,proto3" json:"respDesc,omitempty"`
	User                 *Users   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveResponse) Reset()         { *m = RetrieveResponse{} }
func (m *RetrieveResponse) String() string { return proto.CompactTextString(m) }
func (*RetrieveResponse) ProtoMessage()    {}
func (*RetrieveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *RetrieveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveResponse.Unmarshal(m, b)
}
func (m *RetrieveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveResponse.Marshal(b, m, deterministic)
}
func (m *RetrieveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveResponse.Merge(m, src)
}
func (m *RetrieveResponse) XXX_Size() int {
	return xxx_messageInfo_RetrieveResponse.Size(m)
}
func (m *RetrieveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveResponse proto.InternalMessageInfo

func (m *RetrieveResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RetrieveResponse) GetRespCode() string {
	if m != nil {
		return m.RespCode
	}
	return ""
}

func (m *RetrieveResponse) GetRespDesc() string {
	if m != nil {
		return m.RespDesc
	}
	return ""
}

func (m *RetrieveResponse) GetUser() *Users {
	if m != nil {
		return m.User
	}
	return nil
}

// RetrieveResponse is the response for types service.
type InsertResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RespCode             string   `protobuf:"bytes,2,opt,name=respCode,proto3" json:"respCode,omitempty"`
	RespDesc             string   `protobuf:"bytes,3,opt,name=respDesc,proto3" json:"respDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertResponse) Reset()         { *m = InsertResponse{} }
func (m *InsertResponse) String() string { return proto.CompactTextString(m) }
func (*InsertResponse) ProtoMessage()    {}
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *InsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertResponse.Unmarshal(m, b)
}
func (m *InsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertResponse.Marshal(b, m, deterministic)
}
func (m *InsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertResponse.Merge(m, src)
}
func (m *InsertResponse) XXX_Size() int {
	return xxx_messageInfo_InsertResponse.Size(m)
}
func (m *InsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsertResponse proto.InternalMessageInfo

func (m *InsertResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *InsertResponse) GetRespCode() string {
	if m != nil {
		return m.RespCode
	}
	return ""
}

func (m *InsertResponse) GetRespDesc() string {
	if m != nil {
		return m.RespDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*RetrieveRequest)(nil), "api.customer.bri.v1.users.RetrieveRequest")
	proto.RegisterType((*Users)(nil), "api.customer.bri.v1.users.Users")
	proto.RegisterType((*RetrieveResponse)(nil), "api.customer.bri.v1.users.RetrieveResponse")
	proto.RegisterType((*InsertResponse)(nil), "api.customer.bri.v1.users.InsertResponse")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xfd, 0xf3, 0xfb, 0xd5, 0xad, 0x68, 0x59, 0x3c, 0xc4, 0x80, 0x12, 0x52, 0x10, 0xad,
	0xb0, 0xb1, 0xd5, 0x93, 0x37, 0x51, 0x90, 0x5e, 0x44, 0xa2, 0x1e, 0x3d, 0xe4, 0xcf, 0x10, 0x02,
	0x6d, 0x36, 0xee, 0xec, 0x06, 0x45, 0xbc, 0x78, 0xf0, 0x2e, 0xde, 0xfc, 0x5a, 0x7e, 0x05, 0x3f,
	0x88, 0xec, 0xd6, 0x54, 0x54, 0x5a, 0xbd, 0x78, 0xcb, 0xcb, 0xbc, 0x79, 0x6f, 0xe6, 0xed, 0x90,
	0xb6, 0x42, 0x10, 0xc8, 0x0a, 0xc1, 0x25, 0xa7, 0xab, 0x61, 0x91, 0xb1, 0x58, 0xa1, 0xe4, 0x63,
	0x10, 0x2c, 0x12, 0x19, 0x2b, 0xfb, 0xcc, 0x10, 0x1c, 0x77, 0x94, 0x45, 0x7e, 0xca, 0x79, 0x3a,
	0x02, 0x3f, 0x2c, 0x32, 0x3f, 0xcc, 0x73, 0x2e, 0x43, 0x99, 0xf1, 0xfc, 0xbd, 0xd9, 0xeb, 0x92,
	0xe5, 0x00, 0xa4, 0xc8, 0xa0, 0x84, 0x00, 0xae, 0x14, 0xa0, 0xa4, 0x1d, 0x52, 0x57, 0xc3, 0xc4,
	0xb6, 0x5c, 0x6b, 0x73, 0x21, 0xd0, 0x9f, 0xde, 0x25, 0x69, 0x5e, 0x68, 0xbd, 0xef, 0x25, 0xba,
	0x42, 0x9a, 0xea, 0x24, 0x1c, 0x83, 0x5d, 0x33, 0xff, 0x26, 0x80, 0x3a, 0xa4, 0xa5, 0x0e, 0x92,
	0x44, 0x00, 0xa2, 0x5d, 0x37, 0x85, 0x29, 0xd6, 0x1a, 0x61, 0x0a, 0x76, 0x63, 0xa2, 0x11, 0xa6,
	0xe0, 0x3d, 0x5b, 0xa4, 0xf3, 0x31, 0x04, 0x16, 0x3c, 0x47, 0xa0, 0x36, 0xf9, 0x8f, 0x2a, 0x8e,
	0xb5, 0x82, 0xb6, 0x6b, 0x05, 0x15, 0xd4, 0xe2, 0x02, 0xb0, 0x38, 0xe4, 0x49, 0xe5, 0x3a, 0xc5,
	0x55, 0xed, 0x08, 0x30, 0xae, 0x8c, 0x2b, 0x4c, 0xf7, 0x48, 0x43, 0xa7, 0x62, 0x9c, 0xdb, 0x03,
	0x97, 0xcd, 0x8c, 0x8d, 0x99, 0x65, 0x03, 0xc3, 0xf6, 0x22, 0xb2, 0x34, 0xcc, 0x11, 0x84, 0xfc,
	0xbb, 0xc9, 0x06, 0x8f, 0x35, 0xb2, 0x78, 0x7e, 0x53, 0x00, 0x9e, 0x81, 0x28, 0xb3, 0x18, 0xe8,
	0x83, 0x45, 0xea, 0xc7, 0x20, 0x69, 0x6f, 0xce, 0x90, 0x5f, 0x9e, 0xcd, 0xd9, 0xfe, 0x15, 0x77,
	0xb2, 0x83, 0xb7, 0x71, 0xff, 0xf2, 0xfa, 0x54, 0x73, 0xe9, 0xba, 0x39, 0x8b, 0xaa, 0xc9, 0x2f,
	0xfb, 0x6c, 0xc7, 0x37, 0x2d, 0xfe, 0xad, 0x1a, 0x26, 0x77, 0xf4, 0x9a, 0x34, 0x4e, 0x39, 0x4a,
	0xfa, 0x63, 0x5a, 0xce, 0xd6, 0x1c, 0xc6, 0xe7, 0x00, 0xbd, 0xae, 0x31, 0x5f, 0xf3, 0xec, 0x59,
	0xe6, 0xfb, 0x56, 0x2f, 0xfa, 0x67, 0xee, 0x73, 0xf7, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x90,
	0x67, 0x46, 0xeb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TypesServiceClient is the client API for TypesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TypesServiceClient interface {
	Get(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error)
	Post(ctx context.Context, in *Users, opts ...grpc.CallOption) (*InsertResponse, error)
}

type typesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTypesServiceClient(cc grpc.ClientConnInterface) TypesServiceClient {
	return &typesServiceClient{cc}
}

func (c *typesServiceClient) Get(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error) {
	out := new(RetrieveResponse)
	err := c.cc.Invoke(ctx, "/api.customer.bri.v1.users.TypesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typesServiceClient) Post(ctx context.Context, in *Users, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/api.customer.bri.v1.users.TypesService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypesServiceServer is the server API for TypesService service.
type TypesServiceServer interface {
	Get(context.Context, *RetrieveRequest) (*RetrieveResponse, error)
	Post(context.Context, *Users) (*InsertResponse, error)
}

// UnimplementedTypesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTypesServiceServer struct {
}

func (*UnimplementedTypesServiceServer) Get(ctx context.Context, req *RetrieveRequest) (*RetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTypesServiceServer) Post(ctx context.Context, req *Users) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterTypesServiceServer(s *grpc.Server, srv TypesServiceServer) {
	s.RegisterService(&_TypesService_serviceDesc, srv)
}

func _TypesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.customer.bri.v1.users.TypesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypesServiceServer).Get(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypesService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Users)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypesServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.customer.bri.v1.users.TypesService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypesServiceServer).Post(ctx, req.(*Users))
	}
	return interceptor(ctx, in, info, handler)
}

var _TypesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.customer.bri.v1.users.TypesService",
	HandlerType: (*TypesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TypesService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _TypesService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
