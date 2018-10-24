// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import/proto/import.proto

package import_package

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Name struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b964d71f84c0c1, []int{0}
}

func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_25b964d71f84c0c1, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Name)(nil), "import_package.Name")
	proto.RegisterType((*Empty)(nil), "import_package.Empty")
}

func init() { proto.RegisterFile("import/proto/import.proto", fileDescriptor_25b964d71f84c0c1) }

var fileDescriptor_25b964d71f84c0c1 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x70, 0xf4, 0xc0, 0x1c, 0x21, 0x3e,
	0x08, 0x2f, 0xbe, 0x20, 0x31, 0x39, 0x3b, 0x31, 0x3d, 0x55, 0x49, 0x8a, 0x8b, 0xc5, 0x2f, 0x31,
	0x37, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xcc, 0x56, 0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x34, 0x9a, 0xc9, 0xc8, 0xc5,
	0x1b, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0xea, 0x09, 0xd6, 0x2e, 0xe4, 0xcc, 0x25, 0x92, 0x9e,
	0x5a, 0x12, 0x58, 0x9a, 0x5a, 0x54, 0x09, 0x11, 0xf1, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x11, 0x12,
	0xd1, 0x43, 0x35, 0x5f, 0x0f, 0x64, 0xb8, 0x94, 0x28, 0xba, 0x28, 0xd8, 0x58, 0x90, 0x21, 0xb9,
	0xa5, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x64, 0x1b, 0x92, 0xc4, 0x06, 0xf6, 0x97, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0xa3, 0xdf, 0xad, 0xf4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceImportClient is the client API for ServiceImport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceImportClient interface {
	GetQueryImportMethod(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error)
	MutationImportMethod(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error)
}

type serviceImportClient struct {
	cc *grpc.ClientConn
}

func NewServiceImportClient(cc *grpc.ClientConn) ServiceImportClient {
	return &serviceImportClient{cc}
}

func (c *serviceImportClient) GetQueryImportMethod(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/import_package.ServiceImport/getQueryImportMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceImportClient) MutationImportMethod(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/import_package.ServiceImport/mutationImportMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceImportServer is the server API for ServiceImport service.
type ServiceImportServer interface {
	GetQueryImportMethod(context.Context, *Name) (*Empty, error)
	MutationImportMethod(context.Context, *Name) (*Empty, error)
}

func RegisterServiceImportServer(s *grpc.Server, srv ServiceImportServer) {
	s.RegisterService(&_ServiceImport_serviceDesc, srv)
}

func _ServiceImport_GetQueryImportMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceImportServer).GetQueryImportMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/import_package.ServiceImport/GetQueryImportMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceImportServer).GetQueryImportMethod(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceImport_MutationImportMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceImportServer).MutationImportMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/import_package.ServiceImport/MutationImportMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceImportServer).MutationImportMethod(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceImport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "import_package.ServiceImport",
	HandlerType: (*ServiceImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getQueryImportMethod",
			Handler:    _ServiceImport_GetQueryImportMethod_Handler,
		},
		{
			MethodName: "mutationImportMethod",
			Handler:    _ServiceImport_MutationImportMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "import/proto/import.proto",
}
