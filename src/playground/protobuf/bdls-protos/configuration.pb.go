//
//Copyright  All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: configuration.proto

package bdls

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

// Metadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "BDLS".
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consenters []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options    *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetConsenters() []*Consenter {
	if x != nil {
		return x.Consenters
	}
	return nil
}

func (x *Metadata) GetOptions() *Options {
	if x != nil {
		return x.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ClientTlsCert []byte `protobuf:"bytes,3,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert []byte `protobuf:"bytes,4,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
}

func (x *Consenter) Reset() {
	*x = Consenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consenter) ProtoMessage() {}

func (x *Consenter) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consenter.ProtoReflect.Descriptor instead.
func (*Consenter) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *Consenter) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Consenter) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Consenter) GetClientTlsCert() []byte {
	if x != nil {
		return x.ClientTlsCert
	}
	return nil
}

func (x *Consenter) GetServerTlsCert() []byte {
	if x != nil {
		return x.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the BDLS nodes. These can be modified on a
// per-channel basis.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch         uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"` // specified in miliseconds
	CurrentHeight uint64 `protobuf:"varint,2,opt,name=currentHeight,proto3" json:"currentHeight,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *Options) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Options) GetCurrentHeight() uint64 {
	if x != nil {
		return x.CurrentHeight
	}
	return 0
}

var File_configuration_proto protoreflect.FileDescriptor

var file_configuration_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x64, 0x6c, 0x73, 0x22, 0x64, 0x0a, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x64,
	0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x64, 0x6c, 0x73,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x22, 0x45, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x5f,
	0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x62, 0x64, 0x6c, 0x73, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x62, 0x64, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configuration_proto_rawDescOnce sync.Once
	file_configuration_proto_rawDescData = file_configuration_proto_rawDesc
)

func file_configuration_proto_rawDescGZIP() []byte {
	file_configuration_proto_rawDescOnce.Do(func() {
		file_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_configuration_proto_rawDescData)
	})
	return file_configuration_proto_rawDescData
}

var file_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_configuration_proto_goTypes = []interface{}{
	(*Metadata)(nil),  // 0: bdls.Metadata
	(*Consenter)(nil), // 1: bdls.Consenter
	(*Options)(nil),   // 2: bdls.Options
}
var file_configuration_proto_depIdxs = []int32{
	1, // 0: bdls.Metadata.consenters:type_name -> bdls.Consenter
	2, // 1: bdls.Metadata.options:type_name -> bdls.Options
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_configuration_proto_init() }
func file_configuration_proto_init() {
	if File_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consenter); i {
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
		file_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configuration_proto_goTypes,
		DependencyIndexes: file_configuration_proto_depIdxs,
		MessageInfos:      file_configuration_proto_msgTypes,
	}.Build()
	File_configuration_proto = out.File
	file_configuration_proto_rawDesc = nil
	file_configuration_proto_goTypes = nil
	file_configuration_proto_depIdxs = nil
}
