// Bloodmage Engine
// Copyright (C) 2024 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: engine/level/level.proto

package level

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

type Level struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width   int32  `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height  int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Floor   uint32 `protobuf:"varint,3,opt,name=floor,proto3" json:"floor,omitempty"`
	Ceiling uint32 `protobuf:"varint,4,opt,name=ceiling,proto3" json:"ceiling,omitempty"`
	Walls   []byte `protobuf:"bytes,5,opt,name=walls,proto3" json:"walls,omitempty"`
}

func (x *Level) Reset() {
	*x = Level{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_level_level_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Level) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Level) ProtoMessage() {}

func (x *Level) ProtoReflect() protoreflect.Message {
	mi := &file_engine_level_level_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Level.ProtoReflect.Descriptor instead.
func (*Level) Descriptor() ([]byte, []int) {
	return file_engine_level_level_proto_rawDescGZIP(), []int{0}
}

func (x *Level) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Level) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Level) GetFloor() uint32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

func (x *Level) GetCeiling() uint32 {
	if x != nil {
		return x.Ceiling
	}
	return 0
}

func (x *Level) GetWalls() []byte {
	if x != nil {
		return x.Walls
	}
	return nil
}

var File_engine_level_level_proto protoreflect.FileDescriptor

var file_engine_level_level_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x65, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x65, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x77, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x64, 0x6d, 0x61, 0x67,
	0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_engine_level_level_proto_rawDescOnce sync.Once
	file_engine_level_level_proto_rawDescData = file_engine_level_level_proto_rawDesc
)

func file_engine_level_level_proto_rawDescGZIP() []byte {
	file_engine_level_level_proto_rawDescOnce.Do(func() {
		file_engine_level_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_level_level_proto_rawDescData)
	})
	return file_engine_level_level_proto_rawDescData
}

var file_engine_level_level_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_engine_level_level_proto_goTypes = []interface{}{
	(*Level)(nil), // 0: Level
}
var file_engine_level_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_engine_level_level_proto_init() }
func file_engine_level_level_proto_init() {
	if File_engine_level_level_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_level_level_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Level); i {
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
			RawDescriptor: file_engine_level_level_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_level_level_proto_goTypes,
		DependencyIndexes: file_engine_level_level_proto_depIdxs,
		MessageInfos:      file_engine_level_level_proto_msgTypes,
	}.Build()
	File_engine_level_level_proto = out.File
	file_engine_level_level_proto_rawDesc = nil
	file_engine_level_level_proto_goTypes = nil
	file_engine_level_level_proto_depIdxs = nil
}
