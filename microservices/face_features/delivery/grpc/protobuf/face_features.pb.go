// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: face_features.proto

package face_features

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

type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mask string `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_features_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_face_features_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_face_features_proto_rawDescGZIP(), []int{0}
}

func (x *Photo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Photo) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

type Face struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Have bool `protobuf:"varint,1,opt,name=have,proto3" json:"have,omitempty"`
}

func (x *Face) Reset() {
	*x = Face{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_features_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Face) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Face) ProtoMessage() {}

func (x *Face) ProtoReflect() protoreflect.Message {
	mi := &file_face_features_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Face.ProtoReflect.Descriptor instead.
func (*Face) Descriptor() ([]byte, []int) {
	return file_face_features_proto_rawDescGZIP(), []int{1}
}

func (x *Face) GetHave() bool {
	if x != nil {
		return x.Have
	}
	return false
}

var File_face_features_proto protoreflect.FileDescriptor

var file_face_features_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x1a, 0x0a, 0x04, 0x46, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x61, 0x76,
	0x65, 0x32, 0x83, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x65, 0x47, 0x52, 0x50, 0x43, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x48, 0x61, 0x76, 0x65, 0x46, 0x61, 0x63,
	0x65, 0x12, 0x14, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_face_features_proto_rawDescOnce sync.Once
	file_face_features_proto_rawDescData = file_face_features_proto_rawDesc
)

func file_face_features_proto_rawDescGZIP() []byte {
	file_face_features_proto_rawDescOnce.Do(func() {
		file_face_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_face_features_proto_rawDescData)
	})
	return file_face_features_proto_rawDescData
}

var file_face_features_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_face_features_proto_goTypes = []interface{}{
	(*Photo)(nil), // 0: face_features.Photo
	(*Face)(nil),  // 1: face_features.Face
}
var file_face_features_proto_depIdxs = []int32{
	0, // 0: face_features.FaceGRPCHandler.HaveFace:input_type -> face_features.Photo
	0, // 1: face_features.FaceGRPCHandler.AddMask:input_type -> face_features.Photo
	1, // 2: face_features.FaceGRPCHandler.HaveFace:output_type -> face_features.Face
	0, // 3: face_features.FaceGRPCHandler.AddMask:output_type -> face_features.Photo
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_face_features_proto_init() }
func file_face_features_proto_init() {
	if File_face_features_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_face_features_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
		file_face_features_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Face); i {
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
			RawDescriptor: file_face_features_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_face_features_proto_goTypes,
		DependencyIndexes: file_face_features_proto_depIdxs,
		MessageInfos:      file_face_features_proto_msgTypes,
	}.Build()
	File_face_features_proto = out.File
	file_face_features_proto_rawDesc = nil
	file_face_features_proto_goTypes = nil
	file_face_features_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FaceGRPCHandlerClient is the client API for FaceGRPCHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaceGRPCHandlerClient interface {
	HaveFace(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Face, error)
	AddMask(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error)
}

type faceGRPCHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewFaceGRPCHandlerClient(cc grpc.ClientConnInterface) FaceGRPCHandlerClient {
	return &faceGRPCHandlerClient{cc}
}

func (c *faceGRPCHandlerClient) HaveFace(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Face, error) {
	out := new(Face)
	err := c.cc.Invoke(ctx, "/face_features.FaceGRPCHandler/HaveFace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceGRPCHandlerClient) AddMask(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/face_features.FaceGRPCHandler/AddMask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceGRPCHandlerServer is the server API for FaceGRPCHandler service.
type FaceGRPCHandlerServer interface {
	HaveFace(context.Context, *Photo) (*Face, error)
	AddMask(context.Context, *Photo) (*Photo, error)
}

// UnimplementedFaceGRPCHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedFaceGRPCHandlerServer struct {
}

func (*UnimplementedFaceGRPCHandlerServer) HaveFace(context.Context, *Photo) (*Face, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HaveFace not implemented")
}
func (*UnimplementedFaceGRPCHandlerServer) AddMask(context.Context, *Photo) (*Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMask not implemented")
}

func RegisterFaceGRPCHandlerServer(s *grpc.Server, srv FaceGRPCHandlerServer) {
	s.RegisterService(&_FaceGRPCHandler_serviceDesc, srv)
}

func _FaceGRPCHandler_HaveFace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Photo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceGRPCHandlerServer).HaveFace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/face_features.FaceGRPCHandler/HaveFace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceGRPCHandlerServer).HaveFace(ctx, req.(*Photo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceGRPCHandler_AddMask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Photo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceGRPCHandlerServer).AddMask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/face_features.FaceGRPCHandler/AddMask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceGRPCHandlerServer).AddMask(ctx, req.(*Photo))
	}
	return interceptor(ctx, in, info, handler)
}

var _FaceGRPCHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "face_features.FaceGRPCHandler",
	HandlerType: (*FaceGRPCHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HaveFace",
			Handler:    _FaceGRPCHandler_HaveFace_Handler,
		},
		{
			MethodName: "AddMask",
			Handler:    _FaceGRPCHandler_AddMask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "face_features.proto",
}