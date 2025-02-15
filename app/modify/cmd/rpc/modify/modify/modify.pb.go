// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: modify.proto

package modify

import (
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

type ModifytheUsertypeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	UserType string `protobuf:"bytes,2,opt,name=UserType,proto3" json:"UserType,omitempty"`
}

func (x *ModifytheUsertypeReq) Reset() {
	*x = ModifytheUsertypeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUsertypeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUsertypeReq) ProtoMessage() {}

func (x *ModifytheUsertypeReq) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUsertypeReq.ProtoReflect.Descriptor instead.
func (*ModifytheUsertypeReq) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{0}
}

func (x *ModifytheUsertypeReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifytheUsertypeReq) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type ModifytheUsertypeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (x *ModifytheUsertypeResp) Reset() {
	*x = ModifytheUsertypeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUsertypeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUsertypeResp) ProtoMessage() {}

func (x *ModifytheUsertypeResp) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUsertypeResp.ProtoReflect.Descriptor instead.
func (*ModifytheUsertypeResp) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{1}
}

func (x *ModifytheUsertypeResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type ModifytheUseravatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar string `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *ModifytheUseravatarReq) Reset() {
	*x = ModifytheUseravatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUseravatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUseravatarReq) ProtoMessage() {}

func (x *ModifytheUseravatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUseravatarReq.ProtoReflect.Descriptor instead.
func (*ModifytheUseravatarReq) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{2}
}

func (x *ModifytheUseravatarReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type ModifytheUseravatarResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (x *ModifytheUseravatarResp) Reset() {
	*x = ModifytheUseravatarResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUseravatarResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUseravatarResp) ProtoMessage() {}

func (x *ModifytheUseravatarResp) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUseravatarResp.ProtoReflect.Descriptor instead.
func (*ModifytheUseravatarResp) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{3}
}

func (x *ModifytheUseravatarResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type ModifytheUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
}

func (x *ModifytheUsernameReq) Reset() {
	*x = ModifytheUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUsernameReq) ProtoMessage() {}

func (x *ModifytheUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUsernameReq.ProtoReflect.Descriptor instead.
func (*ModifytheUsernameReq) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{4}
}

func (x *ModifytheUsernameReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type ModifytheUsernameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
}

func (x *ModifytheUsernameResp) Reset() {
	*x = ModifytheUsernameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifytheUsernameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifytheUsernameResp) ProtoMessage() {}

func (x *ModifytheUsernameResp) ProtoReflect() protoreflect.Message {
	mi := &file_modify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifytheUsernameResp.ProtoReflect.Descriptor instead.
func (*ModifytheUsernameResp) Descriptor() ([]byte, []int) {
	return file_modify_proto_rawDescGZIP(), []int{5}
}

func (x *ModifytheUsernameResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

var File_modify_proto protoreflect.FileDescriptor

var file_modify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x22, 0x48, 0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x2b, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x30, 0x0a,
	0x16, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22,
	0x2d, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x32,
	0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x32,
	0x84, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x50, 0x0a, 0x11, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74,
	0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x56, 0x0a, 0x13,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x50, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x74, 0x68, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modify_proto_rawDescOnce sync.Once
	file_modify_proto_rawDescData = file_modify_proto_rawDesc
)

func file_modify_proto_rawDescGZIP() []byte {
	file_modify_proto_rawDescOnce.Do(func() {
		file_modify_proto_rawDescData = protoimpl.X.CompressGZIP(file_modify_proto_rawDescData)
	})
	return file_modify_proto_rawDescData
}

var file_modify_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_modify_proto_goTypes = []interface{}{
	(*ModifytheUsertypeReq)(nil),    // 0: modify.ModifytheUsertypeReq
	(*ModifytheUsertypeResp)(nil),   // 1: modify.ModifytheUsertypeResp
	(*ModifytheUseravatarReq)(nil),  // 2: modify.ModifytheUseravatarReq
	(*ModifytheUseravatarResp)(nil), // 3: modify.ModifytheUseravatarResp
	(*ModifytheUsernameReq)(nil),    // 4: modify.ModifytheUsernameReq
	(*ModifytheUsernameResp)(nil),   // 5: modify.ModifytheUsernameResp
}
var file_modify_proto_depIdxs = []int32{
	0, // 0: modify.Modify.ModifytheUsertype:input_type -> modify.ModifytheUsertypeReq
	2, // 1: modify.Modify.ModifytheUseravatar:input_type -> modify.ModifytheUseravatarReq
	4, // 2: modify.Modify.ModifytheUsername:input_type -> modify.ModifytheUsernameReq
	1, // 3: modify.Modify.ModifytheUsertype:output_type -> modify.ModifytheUsertypeResp
	3, // 4: modify.Modify.ModifytheUseravatar:output_type -> modify.ModifytheUseravatarResp
	5, // 5: modify.Modify.ModifytheUsername:output_type -> modify.ModifytheUsernameResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modify_proto_init() }
func file_modify_proto_init() {
	if File_modify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUsertypeReq); i {
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
		file_modify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUsertypeResp); i {
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
		file_modify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUseravatarReq); i {
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
		file_modify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUseravatarResp); i {
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
		file_modify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUsernameReq); i {
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
		file_modify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifytheUsernameResp); i {
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
			RawDescriptor: file_modify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modify_proto_goTypes,
		DependencyIndexes: file_modify_proto_depIdxs,
		MessageInfos:      file_modify_proto_msgTypes,
	}.Build()
	File_modify_proto = out.File
	file_modify_proto_rawDesc = nil
	file_modify_proto_goTypes = nil
	file_modify_proto_depIdxs = nil
}
