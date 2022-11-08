// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/def.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PostParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	IsAdmin bool   `protobuf:"varint,3,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *PostParam) Reset() {
	*x = PostParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostParam) ProtoMessage() {}

func (x *PostParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostParam.ProtoReflect.Descriptor instead.
func (*PostParam) Descriptor() ([]byte, []int) {
	return file_proto_def_proto_rawDescGZIP(), []int{0}
}

func (x *PostParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostParam) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PostParam) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type GetParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetParam) Reset() {
	*x = GetParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParam) ProtoMessage() {}

func (x *GetParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParam.ProtoReflect.Descriptor instead.
func (*GetParam) Descriptor() ([]byte, []int) {
	return file_proto_def_proto_rawDescGZIP(), []int{1}
}

func (x *GetParam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	IsAdmin bool   `protobuf:"varint,4,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_proto_def_proto_rawDescGZIP(), []int{2}
}

func (x *NewUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewUser) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *NewUser) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*NewUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_proto_def_proto_rawDescGZIP(), []int{3}
}

func (x *UserList) GetUsers() []*NewUser {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_def_proto protoreflect.FileDescriptor

var file_proto_def_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x22, 0x1a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x59, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x31, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa2, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x53, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_def_proto_rawDescOnce sync.Once
	file_proto_def_proto_rawDescData = file_proto_def_proto_rawDesc
)

func file_proto_def_proto_rawDescGZIP() []byte {
	file_proto_def_proto_rawDescOnce.Do(func() {
		file_proto_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_def_proto_rawDescData)
	})
	return file_proto_def_proto_rawDescData
}

var file_proto_def_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_def_proto_goTypes = []interface{}{
	(*PostParam)(nil),   // 0: server.PostParam
	(*GetParam)(nil),    // 1: server.GetParam
	(*NewUser)(nil),     // 2: server.NewUser
	(*UserList)(nil),    // 3: server.UserList
	(*empty.Empty)(nil), // 4: google.protobuf.Empty
}
var file_proto_def_proto_depIdxs = []int32{
	2, // 0: server.UserList.users:type_name -> server.NewUser
	0, // 1: server.User.CreateUSer:input_type -> server.PostParam
	1, // 2: server.User.GetUser:input_type -> server.GetParam
	4, // 3: server.User.GetUsers:input_type -> google.protobuf.Empty
	2, // 4: server.User.CreateUSer:output_type -> server.NewUser
	2, // 5: server.User.GetUser:output_type -> server.NewUser
	3, // 6: server.User.GetUsers:output_type -> server.UserList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_def_proto_init() }
func file_proto_def_proto_init() {
	if File_proto_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostParam); i {
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
		file_proto_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParam); i {
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
		file_proto_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_proto_def_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
			RawDescriptor: file_proto_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_def_proto_goTypes,
		DependencyIndexes: file_proto_def_proto_depIdxs,
		MessageInfos:      file_proto_def_proto_msgTypes,
	}.Build()
	File_proto_def_proto = out.File
	file_proto_def_proto_rawDesc = nil
	file_proto_def_proto_goTypes = nil
	file_proto_def_proto_depIdxs = nil
}
