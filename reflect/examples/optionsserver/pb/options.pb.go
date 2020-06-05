// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.7.1
// source: reflect/examples/optionsserver/pb/options.proto

package pb

import (
	context "context"
	_ "github.com/danielvladco/go-proto-gql/pb"
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ string    `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`          // must be required
	Foo     *Foo2     `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`                // must be required
	Float   []float32 `protobuf:"fixed32,3,rep,packed,name=float,proto3" json:"float,omitempty"`   // must be required because its greater than 0
	String2 string    `protobuf:"bytes,4,opt,name=string2,proto3" json:"string2,omitempty"`        // simple
	Foo2    *Foo2     `protobuf:"bytes,5,opt,name=foo2,proto3" json:"foo2,omitempty"`              // simple
	Float2  []float32 `protobuf:"fixed32,6,rep,packed,name=float2,proto3" json:"float2,omitempty"` // simple
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflect_examples_optionsserver_pb_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_reflect_examples_optionsserver_pb_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_reflect_examples_optionsserver_pb_options_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Data) GetFoo() *Foo2 {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *Data) GetFloat() []float32 {
	if x != nil {
		return x.Float
	}
	return nil
}

func (x *Data) GetString2() string {
	if x != nil {
		return x.String2
	}
	return ""
}

func (x *Data) GetFoo2() *Foo2 {
	if x != nil {
		return x.Foo2
	}
	return nil
}

func (x *Data) GetFloat2() []float32 {
	if x != nil {
		return x.Float2
	}
	return nil
}

type Foo2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param1 string `protobuf:"bytes,1,opt,name=param1,proto3" json:"param1,omitempty"`
}

func (x *Foo2) Reset() {
	*x = Foo2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflect_examples_optionsserver_pb_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo2) ProtoMessage() {}

func (x *Foo2) ProtoReflect() protoreflect.Message {
	mi := &file_reflect_examples_optionsserver_pb_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo2.ProtoReflect.Descriptor instead.
func (*Foo2) Descriptor() ([]byte, []int) {
	return file_reflect_examples_optionsserver_pb_options_proto_rawDescGZIP(), []int{1}
}

func (x *Foo2) GetParam1() string {
	if x != nil {
		return x.Param1
	}
	return ""
}

var File_reflect_examples_optionsserver_pb_options_proto protoreflect.FileDescriptor

var file_reflect_examples_optionsserver_pb_options_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x70, 0x62, 0x2f, 0x67, 0x71, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xb2, 0xe0,
	0x1f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x03,
	0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x6f, 0x6f, 0x32, 0x42, 0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x52, 0x03, 0x66, 0x6f, 0x6f,
	0x12, 0x1c, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x42,
	0x06, 0xb2, 0xe0, 0x1f, 0x02, 0x08, 0x01, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x32,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6f, 0x32,
	0x52, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x32,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x32, 0x22, 0x1e,
	0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x31, 0x32, 0x90,
	0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x07, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x32, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04, 0xb0, 0xe0, 0x1f, 0x02, 0x12, 0x1f, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x12, 0x21, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01,
	0x12, 0x21, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04, 0xb0, 0xe0,
	0x1f, 0x02, 0x28, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x32, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04, 0xb0,
	0xe0, 0x1f, 0x01, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x33, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04,
	0xb0, 0xe0, 0x1f, 0x02, 0x28, 0x01, 0x30, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x53,
	0x75, 0x62, 0x32, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04, 0xb0, 0xe0, 0x1f, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x32, 0x91, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x06, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x06, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x32, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x07, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x31, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x04, 0xb0, 0xe0, 0x1f, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x1a,
	0x04, 0xb0, 0xe0, 0x1f, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x76, 0x6c, 0x61, 0x64, 0x63, 0x6f,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x71, 0x6c, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reflect_examples_optionsserver_pb_options_proto_rawDescOnce sync.Once
	file_reflect_examples_optionsserver_pb_options_proto_rawDescData = file_reflect_examples_optionsserver_pb_options_proto_rawDesc
)

func file_reflect_examples_optionsserver_pb_options_proto_rawDescGZIP() []byte {
	file_reflect_examples_optionsserver_pb_options_proto_rawDescOnce.Do(func() {
		file_reflect_examples_optionsserver_pb_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflect_examples_optionsserver_pb_options_proto_rawDescData)
	})
	return file_reflect_examples_optionsserver_pb_options_proto_rawDescData
}

var file_reflect_examples_optionsserver_pb_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reflect_examples_optionsserver_pb_options_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: pb.Data
	(*Foo2)(nil), // 1: pb.Foo2
}
var file_reflect_examples_optionsserver_pb_options_proto_depIdxs = []int32{
	1,  // 0: pb.Data.foo:type_name -> pb.Foo2
	1,  // 1: pb.Data.foo2:type_name -> pb.Foo2
	0,  // 2: pb.Service.Mutate1:input_type -> pb.Data
	0,  // 3: pb.Service.Mutate2:input_type -> pb.Data
	0,  // 4: pb.Service.Query1:input_type -> pb.Data
	0,  // 5: pb.Service.Publish:input_type -> pb.Data
	0,  // 6: pb.Service.Subscribe:input_type -> pb.Data
	0,  // 7: pb.Service.PubSub1:input_type -> pb.Data
	0,  // 8: pb.Service.InvalidSubscribe1:input_type -> pb.Data
	0,  // 9: pb.Service.InvalidSubscribe2:input_type -> pb.Data
	0,  // 10: pb.Service.InvalidSubscribe3:input_type -> pb.Data
	0,  // 11: pb.Service.PubSub2:input_type -> pb.Data
	0,  // 12: pb.Query.Query1:input_type -> pb.Data
	0,  // 13: pb.Query.Query2:input_type -> pb.Data
	0,  // 14: pb.Query.Mutate1:input_type -> pb.Data
	0,  // 15: pb.Query.Subscribe:input_type -> pb.Data
	0,  // 16: pb.Service.Mutate1:output_type -> pb.Data
	0,  // 17: pb.Service.Mutate2:output_type -> pb.Data
	0,  // 18: pb.Service.Query1:output_type -> pb.Data
	0,  // 19: pb.Service.Publish:output_type -> pb.Data
	0,  // 20: pb.Service.Subscribe:output_type -> pb.Data
	0,  // 21: pb.Service.PubSub1:output_type -> pb.Data
	0,  // 22: pb.Service.InvalidSubscribe1:output_type -> pb.Data
	0,  // 23: pb.Service.InvalidSubscribe2:output_type -> pb.Data
	0,  // 24: pb.Service.InvalidSubscribe3:output_type -> pb.Data
	0,  // 25: pb.Service.PubSub2:output_type -> pb.Data
	0,  // 26: pb.Query.Query1:output_type -> pb.Data
	0,  // 27: pb.Query.Query2:output_type -> pb.Data
	0,  // 28: pb.Query.Mutate1:output_type -> pb.Data
	0,  // 29: pb.Query.Subscribe:output_type -> pb.Data
	16, // [16:30] is the sub-list for method output_type
	2,  // [2:16] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_reflect_examples_optionsserver_pb_options_proto_init() }
func file_reflect_examples_optionsserver_pb_options_proto_init() {
	if File_reflect_examples_optionsserver_pb_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reflect_examples_optionsserver_pb_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_reflect_examples_optionsserver_pb_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo2); i {
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
			RawDescriptor: file_reflect_examples_optionsserver_pb_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_reflect_examples_optionsserver_pb_options_proto_goTypes,
		DependencyIndexes: file_reflect_examples_optionsserver_pb_options_proto_depIdxs,
		MessageInfos:      file_reflect_examples_optionsserver_pb_options_proto_msgTypes,
	}.Build()
	File_reflect_examples_optionsserver_pb_options_proto = out.File
	file_reflect_examples_optionsserver_pb_options_proto_rawDesc = nil
	file_reflect_examples_optionsserver_pb_options_proto_goTypes = nil
	file_reflect_examples_optionsserver_pb_options_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Mutate1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Mutate2(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Query1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (Service_PublishClient, error)
	Subscribe(ctx context.Context, in *Data, opts ...grpc.CallOption) (Service_SubscribeClient, error)
	PubSub1(ctx context.Context, opts ...grpc.CallOption) (Service_PubSub1Client, error)
	InvalidSubscribe1(ctx context.Context, opts ...grpc.CallOption) (Service_InvalidSubscribe1Client, error)
	InvalidSubscribe2(ctx context.Context, in *Data, opts ...grpc.CallOption) (Service_InvalidSubscribe2Client, error)
	InvalidSubscribe3(ctx context.Context, opts ...grpc.CallOption) (Service_InvalidSubscribe3Client, error)
	PubSub2(ctx context.Context, opts ...grpc.CallOption) (Service_PubSub2Client, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Mutate1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Service/Mutate1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Mutate2(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Service/Mutate2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Query1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Service/Query1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (Service_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/pb.Service/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicePublishClient{stream}
	return x, nil
}

type Service_PublishClient interface {
	Send(*Data) error
	CloseAndRecv() (*Data, error)
	grpc.ClientStream
}

type servicePublishClient struct {
	grpc.ClientStream
}

func (x *servicePublishClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicePublishClient) CloseAndRecv() (*Data, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Subscribe(ctx context.Context, in *Data, opts ...grpc.CallOption) (Service_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[1], "/pb.Service/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_SubscribeClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type serviceSubscribeClient struct {
	grpc.ClientStream
}

func (x *serviceSubscribeClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) PubSub1(ctx context.Context, opts ...grpc.CallOption) (Service_PubSub1Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[2], "/pb.Service/PubSub1", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicePubSub1Client{stream}
	return x, nil
}

type Service_PubSub1Client interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type servicePubSub1Client struct {
	grpc.ClientStream
}

func (x *servicePubSub1Client) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicePubSub1Client) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) InvalidSubscribe1(ctx context.Context, opts ...grpc.CallOption) (Service_InvalidSubscribe1Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[3], "/pb.Service/InvalidSubscribe1", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceInvalidSubscribe1Client{stream}
	return x, nil
}

type Service_InvalidSubscribe1Client interface {
	Send(*Data) error
	CloseAndRecv() (*Data, error)
	grpc.ClientStream
}

type serviceInvalidSubscribe1Client struct {
	grpc.ClientStream
}

func (x *serviceInvalidSubscribe1Client) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceInvalidSubscribe1Client) CloseAndRecv() (*Data, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) InvalidSubscribe2(ctx context.Context, in *Data, opts ...grpc.CallOption) (Service_InvalidSubscribe2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[4], "/pb.Service/InvalidSubscribe2", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceInvalidSubscribe2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_InvalidSubscribe2Client interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type serviceInvalidSubscribe2Client struct {
	grpc.ClientStream
}

func (x *serviceInvalidSubscribe2Client) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) InvalidSubscribe3(ctx context.Context, opts ...grpc.CallOption) (Service_InvalidSubscribe3Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[5], "/pb.Service/InvalidSubscribe3", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceInvalidSubscribe3Client{stream}
	return x, nil
}

type Service_InvalidSubscribe3Client interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type serviceInvalidSubscribe3Client struct {
	grpc.ClientStream
}

func (x *serviceInvalidSubscribe3Client) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceInvalidSubscribe3Client) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) PubSub2(ctx context.Context, opts ...grpc.CallOption) (Service_PubSub2Client, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[6], "/pb.Service/PubSub2", opts...)
	if err != nil {
		return nil, err
	}
	x := &servicePubSub2Client{stream}
	return x, nil
}

type Service_PubSub2Client interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type servicePubSub2Client struct {
	grpc.ClientStream
}

func (x *servicePubSub2Client) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicePubSub2Client) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Mutate1(context.Context, *Data) (*Data, error)
	Mutate2(context.Context, *Data) (*Data, error)
	Query1(context.Context, *Data) (*Data, error)
	Publish(Service_PublishServer) error
	Subscribe(*Data, Service_SubscribeServer) error
	PubSub1(Service_PubSub1Server) error
	InvalidSubscribe1(Service_InvalidSubscribe1Server) error
	InvalidSubscribe2(*Data, Service_InvalidSubscribe2Server) error
	InvalidSubscribe3(Service_InvalidSubscribe3Server) error
	PubSub2(Service_PubSub2Server) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Mutate1(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate1 not implemented")
}
func (*UnimplementedServiceServer) Mutate2(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate2 not implemented")
}
func (*UnimplementedServiceServer) Query1(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query1 not implemented")
}
func (*UnimplementedServiceServer) Publish(Service_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedServiceServer) Subscribe(*Data, Service_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedServiceServer) PubSub1(Service_PubSub1Server) error {
	return status.Errorf(codes.Unimplemented, "method PubSub1 not implemented")
}
func (*UnimplementedServiceServer) InvalidSubscribe1(Service_InvalidSubscribe1Server) error {
	return status.Errorf(codes.Unimplemented, "method InvalidSubscribe1 not implemented")
}
func (*UnimplementedServiceServer) InvalidSubscribe2(*Data, Service_InvalidSubscribe2Server) error {
	return status.Errorf(codes.Unimplemented, "method InvalidSubscribe2 not implemented")
}
func (*UnimplementedServiceServer) InvalidSubscribe3(Service_InvalidSubscribe3Server) error {
	return status.Errorf(codes.Unimplemented, "method InvalidSubscribe3 not implemented")
}
func (*UnimplementedServiceServer) PubSub2(Service_PubSub2Server) error {
	return status.Errorf(codes.Unimplemented, "method PubSub2 not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Mutate1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Mutate1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Mutate1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Mutate1(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Mutate2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Mutate2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Mutate2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Mutate2(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Query1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Query1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Query1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Query1(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Publish(&servicePublishServer{stream})
}

type Service_PublishServer interface {
	SendAndClose(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type servicePublishServer struct {
	grpc.ServerStream
}

func (x *servicePublishServer) SendAndClose(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicePublishServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Data)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Subscribe(m, &serviceSubscribeServer{stream})
}

type Service_SubscribeServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type serviceSubscribeServer struct {
	grpc.ServerStream
}

func (x *serviceSubscribeServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_PubSub1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).PubSub1(&servicePubSub1Server{stream})
}

type Service_PubSub1Server interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type servicePubSub1Server struct {
	grpc.ServerStream
}

func (x *servicePubSub1Server) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicePubSub1Server) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_InvalidSubscribe1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).InvalidSubscribe1(&serviceInvalidSubscribe1Server{stream})
}

type Service_InvalidSubscribe1Server interface {
	SendAndClose(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type serviceInvalidSubscribe1Server struct {
	grpc.ServerStream
}

func (x *serviceInvalidSubscribe1Server) SendAndClose(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceInvalidSubscribe1Server) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_InvalidSubscribe2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Data)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).InvalidSubscribe2(m, &serviceInvalidSubscribe2Server{stream})
}

type Service_InvalidSubscribe2Server interface {
	Send(*Data) error
	grpc.ServerStream
}

type serviceInvalidSubscribe2Server struct {
	grpc.ServerStream
}

func (x *serviceInvalidSubscribe2Server) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_InvalidSubscribe3_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).InvalidSubscribe3(&serviceInvalidSubscribe3Server{stream})
}

type Service_InvalidSubscribe3Server interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type serviceInvalidSubscribe3Server struct {
	grpc.ServerStream
}

func (x *serviceInvalidSubscribe3Server) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceInvalidSubscribe3Server) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_PubSub2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).PubSub2(&servicePubSub2Server{stream})
}

type Service_PubSub2Server interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type servicePubSub2Server struct {
	grpc.ServerStream
}

func (x *servicePubSub2Server) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicePubSub2Server) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mutate1",
			Handler:    _Service_Mutate1_Handler,
		},
		{
			MethodName: "Mutate2",
			Handler:    _Service_Mutate2_Handler,
		},
		{
			MethodName: "Query1",
			Handler:    _Service_Query1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _Service_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Service_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PubSub1",
			Handler:       _Service_PubSub1_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "InvalidSubscribe1",
			Handler:       _Service_InvalidSubscribe1_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "InvalidSubscribe2",
			Handler:       _Service_InvalidSubscribe2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "InvalidSubscribe3",
			Handler:       _Service_InvalidSubscribe3_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PubSub2",
			Handler:       _Service_PubSub2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "reflect/examples/optionsserver/pb/options.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Query1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Query2(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Mutate1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	Subscribe(ctx context.Context, in *Data, opts ...grpc.CallOption) (Query_SubscribeClient, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Query1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Query/Query1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Query2(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Query/Query2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Mutate1(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/pb.Query/Mutate1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Subscribe(ctx context.Context, in *Data, opts ...grpc.CallOption) (Query_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[0], "/pb.Query/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &querySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_SubscribeClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type querySubscribeClient struct {
	grpc.ClientStream
}

func (x *querySubscribeClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Query1(context.Context, *Data) (*Data, error)
	Query2(context.Context, *Data) (*Data, error)
	Mutate1(context.Context, *Data) (*Data, error)
	Subscribe(*Data, Query_SubscribeServer) error
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Query1(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query1 not implemented")
}
func (*UnimplementedQueryServer) Query2(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query2 not implemented")
}
func (*UnimplementedQueryServer) Mutate1(context.Context, *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate1 not implemented")
}
func (*UnimplementedQueryServer) Subscribe(*Data, Query_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Query1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/Query1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query1(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Query2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/Query2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query2(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Mutate1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Mutate1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/Mutate1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Mutate1(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Data)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).Subscribe(m, &querySubscribeServer{stream})
}

type Query_SubscribeServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type querySubscribeServer struct {
	grpc.ServerStream
}

func (x *querySubscribeServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query1",
			Handler:    _Query_Query1_Handler,
		},
		{
			MethodName: "Query2",
			Handler:    _Query_Query2_Handler,
		},
		{
			MethodName: "Mutate1",
			Handler:    _Query_Mutate1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Query_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reflect/examples/optionsserver/pb/options.proto",
}