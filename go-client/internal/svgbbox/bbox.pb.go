// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.3
// source: bbox.proto

package svgbbox

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

type Svg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Svg) Reset() {
	*x = Svg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Svg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Svg) ProtoMessage() {}

func (x *Svg) ProtoReflect() protoreflect.Message {
	mi := &file_bbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Svg.ProtoReflect.Descriptor instead.
func (*Svg) Descriptor() ([]byte, []int) {
	return file_bbox_proto_rawDescGZIP(), []int{0}
}

func (x *Svg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type BBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y      float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Width  float32 `protobuf:"fixed32,3,opt,name=width,proto3" json:"width,omitempty"`
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *BBox) Reset() {
	*x = BBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBox) ProtoMessage() {}

func (x *BBox) ProtoReflect() protoreflect.Message {
	mi := &file_bbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBox.ProtoReflect.Descriptor instead.
func (*BBox) Descriptor() ([]byte, []int) {
	return file_bbox_proto_rawDescGZIP(), []int{1}
}

func (x *BBox) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BBox) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BBox) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BBox) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_bbox_proto protoreflect.FileDescriptor

var file_bbox_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x76,
	0x67, 0x62, 0x62, 0x6f, 0x78, 0x22, 0x1f, 0x0a, 0x03, 0x53, 0x76, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x04, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0x39, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x42, 0x6f, 0x78, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x42, 0x6f, 0x78, 0x12, 0x0c, 0x2e, 0x73, 0x76, 0x67, 0x62, 0x62, 0x6f, 0x78, 0x2e, 0x53,
	0x76, 0x67, 0x1a, 0x0d, 0x2e, 0x73, 0x76, 0x67, 0x62, 0x62, 0x6f, 0x78, 0x2e, 0x42, 0x42, 0x6f,
	0x78, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x76, 0x67, 0x62, 0x62, 0x6f, 0x78,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bbox_proto_rawDescOnce sync.Once
	file_bbox_proto_rawDescData = file_bbox_proto_rawDesc
)

func file_bbox_proto_rawDescGZIP() []byte {
	file_bbox_proto_rawDescOnce.Do(func() {
		file_bbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_bbox_proto_rawDescData)
	})
	return file_bbox_proto_rawDescData
}

var file_bbox_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bbox_proto_goTypes = []interface{}{
	(*Svg)(nil),  // 0: svgbbox.Svg
	(*BBox)(nil), // 1: svgbbox.BBox
}
var file_bbox_proto_depIdxs = []int32{
	0, // 0: svgbbox.CalculateBBox.GetBBox:input_type -> svgbbox.Svg
	1, // 1: svgbbox.CalculateBBox.GetBBox:output_type -> svgbbox.BBox
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bbox_proto_init() }
func file_bbox_proto_init() {
	if File_bbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Svg); i {
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
		file_bbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBox); i {
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
			RawDescriptor: file_bbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bbox_proto_goTypes,
		DependencyIndexes: file_bbox_proto_depIdxs,
		MessageInfos:      file_bbox_proto_msgTypes,
	}.Build()
	File_bbox_proto = out.File
	file_bbox_proto_rawDesc = nil
	file_bbox_proto_goTypes = nil
	file_bbox_proto_depIdxs = nil
}