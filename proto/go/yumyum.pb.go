// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: yumyum.proto

// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
// go install google.golang.org/protobuf/cmd/protoc-gen-go
// PATH="${PATH}:${HOME}/go/bin"
// protoc --go_out=. --go-grpc_out=. yumyum.proto                                                                             ✔  14:01:41 

package _go

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

type EmojiReaction int32

const (
	EmojiReaction_UNKNOWN EmojiReaction = 0
	EmojiReaction_LIKE    EmojiReaction = 1
	EmojiReaction_LOVE    EmojiReaction = 2
	EmojiReaction_HAHA    EmojiReaction = 3
	EmojiReaction_WOW     EmojiReaction = 4
	EmojiReaction_SAD     EmojiReaction = 5
	EmojiReaction_ANGRY   EmojiReaction = 6
)

// Enum value maps for EmojiReaction.
var (
	EmojiReaction_name = map[int32]string{
		0: "UNKNOWN",
		1: "LIKE",
		2: "LOVE",
		3: "HAHA",
		4: "WOW",
		5: "SAD",
		6: "ANGRY",
	}
	EmojiReaction_value = map[string]int32{
		"UNKNOWN": 0,
		"LIKE":    1,
		"LOVE":    2,
		"HAHA":    3,
		"WOW":     4,
		"SAD":     5,
		"ANGRY":   6,
	}
)

func (x EmojiReaction) Enum() *EmojiReaction {
	p := new(EmojiReaction)
	*p = x
	return p
}

func (x EmojiReaction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmojiReaction) Descriptor() protoreflect.EnumDescriptor {
	return file_yumyum_proto_enumTypes[0].Descriptor()
}

func (EmojiReaction) Type() protoreflect.EnumType {
	return &file_yumyum_proto_enumTypes[0]
}

func (x EmojiReaction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmojiReaction.Descriptor instead.
func (EmojiReaction) EnumDescriptor() ([]byte, []int) {
	return file_yumyum_proto_rawDescGZIP(), []int{0}
}

type Emoji struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reaction EmojiReaction `protobuf:"varint,1,opt,name=reaction,proto3,enum=reactions.EmojiReaction" json:"reaction,omitempty"`
}

func (x *Emoji) Reset() {
	*x = Emoji{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yumyum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emoji) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emoji) ProtoMessage() {}

func (x *Emoji) ProtoReflect() protoreflect.Message {
	mi := &file_yumyum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emoji.ProtoReflect.Descriptor instead.
func (*Emoji) Descriptor() ([]byte, []int) {
	return file_yumyum_proto_rawDescGZIP(), []int{0}
}

func (x *Emoji) GetReaction() EmojiReaction {
	if x != nil {
		return x.Reaction
	}
	return EmojiReaction_UNKNOWN
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yumyum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_yumyum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_yumyum_proto_rawDescGZIP(), []int{1}
}

var File_yumyum_proto protoreflect.FileDescriptor

var file_yumyum_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x79, 0x75, 0x6d, 0x79, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3d, 0x0a, 0x05, 0x45, 0x6d, 0x6f,
	0x6a, 0x69, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x2a, 0x57, 0x0a, 0x0d, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x48, 0x41, 0x10, 0x03, 0x12, 0x07, 0x0a,
	0x03, 0x57, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x41, 0x44, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x4e, 0x47, 0x52, 0x59, 0x10, 0x06, 0x32, 0x46, 0x0a, 0x0d, 0x59, 0x75,
	0x6d, 0x59, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x45,
	0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yumyum_proto_rawDescOnce sync.Once
	file_yumyum_proto_rawDescData = file_yumyum_proto_rawDesc
)

func file_yumyum_proto_rawDescGZIP() []byte {
	file_yumyum_proto_rawDescOnce.Do(func() {
		file_yumyum_proto_rawDescData = protoimpl.X.CompressGZIP(file_yumyum_proto_rawDescData)
	})
	return file_yumyum_proto_rawDescData
}

var file_yumyum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yumyum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yumyum_proto_goTypes = []interface{}{
	(EmojiReaction)(0), // 0: reactions.EmojiReaction
	(*Emoji)(nil),      // 1: reactions.Emoji
	(*Empty)(nil),      // 2: reactions.Empty
}
var file_yumyum_proto_depIdxs = []int32{
	0, // 0: reactions.Emoji.reaction:type_name -> reactions.EmojiReaction
	1, // 1: reactions.YumYumService.EmojiChat:input_type -> reactions.Emoji
	1, // 2: reactions.YumYumService.EmojiChat:output_type -> reactions.Emoji
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yumyum_proto_init() }
func file_yumyum_proto_init() {
	if File_yumyum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yumyum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emoji); i {
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
		file_yumyum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_yumyum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yumyum_proto_goTypes,
		DependencyIndexes: file_yumyum_proto_depIdxs,
		EnumInfos:         file_yumyum_proto_enumTypes,
		MessageInfos:      file_yumyum_proto_msgTypes,
	}.Build()
	File_yumyum_proto = out.File
	file_yumyum_proto_rawDesc = nil
	file_yumyum_proto_goTypes = nil
	file_yumyum_proto_depIdxs = nil
}