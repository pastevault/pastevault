// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: paste.proto

package proto

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

type PasteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content    string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Expiration string `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Encrypted  bool   `protobuf:"varint,3,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *PasteRequest) Reset() {
	*x = PasteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paste_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasteRequest) ProtoMessage() {}

func (x *PasteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paste_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasteRequest.ProtoReflect.Descriptor instead.
func (*PasteRequest) Descriptor() ([]byte, []int) {
	return file_paste_proto_rawDescGZIP(), []int{0}
}

func (x *PasteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PasteRequest) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *PasteRequest) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

type Paste struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Expiration string `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Encrypted  bool   `protobuf:"varint,4,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *Paste) Reset() {
	*x = Paste{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paste_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paste) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paste) ProtoMessage() {}

func (x *Paste) ProtoReflect() protoreflect.Message {
	mi := &file_paste_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paste.ProtoReflect.Descriptor instead.
func (*Paste) Descriptor() ([]byte, []int) {
	return file_paste_proto_rawDescGZIP(), []int{1}
}

func (x *Paste) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Paste) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Paste) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *Paste) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

var File_paste_proto protoreflect.FileDescriptor

var file_paste_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x73, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x22, 0x6f, 0x0a, 0x05,
	0x50, 0x61, 0x73, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a,
	0x1d, 0x70, 0x61, 0x73, 0x74, 0x65, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paste_proto_rawDescOnce sync.Once
	file_paste_proto_rawDescData = file_paste_proto_rawDesc
)

func file_paste_proto_rawDescGZIP() []byte {
	file_paste_proto_rawDescOnce.Do(func() {
		file_paste_proto_rawDescData = protoimpl.X.CompressGZIP(file_paste_proto_rawDescData)
	})
	return file_paste_proto_rawDescData
}

var file_paste_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paste_proto_goTypes = []interface{}{
	(*PasteRequest)(nil), // 0: proto.PasteRequest
	(*Paste)(nil),        // 1: proto.Paste
}
var file_paste_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paste_proto_init() }
func file_paste_proto_init() {
	if File_paste_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paste_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasteRequest); i {
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
		file_paste_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paste); i {
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
			RawDescriptor: file_paste_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paste_proto_goTypes,
		DependencyIndexes: file_paste_proto_depIdxs,
		MessageInfos:      file_paste_proto_msgTypes,
	}.Build()
	File_paste_proto = out.File
	file_paste_proto_rawDesc = nil
	file_paste_proto_goTypes = nil
	file_paste_proto_depIdxs = nil
}
