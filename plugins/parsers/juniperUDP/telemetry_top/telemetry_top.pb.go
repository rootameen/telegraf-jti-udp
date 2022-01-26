// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_top.proto

/*
Package telemetry_top is a generated protocol buffer package.

It is generated from these files:
	telemetry_top.proto

It has these top-level messages:
	TelemetryFieldOptions
	TelemetryStream
	IETFSensors
	EnterpriseSensors
	JuniperNetworksSensors
*/
package telemetry_top

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TelemetryFieldOptions struct {
	IsKey            *bool  `protobuf:"varint,1,opt,name=is_key" json:"is_key,omitempty"`
	IsTimestamp      *bool  `protobuf:"varint,2,opt,name=is_timestamp" json:"is_timestamp,omitempty"`
	IsCounter        *bool  `protobuf:"varint,3,opt,name=is_counter" json:"is_counter,omitempty"`
	IsGauge          *bool  `protobuf:"varint,4,opt,name=is_gauge" json:"is_gauge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TelemetryFieldOptions) Reset()                    { *m = TelemetryFieldOptions{} }
func (m *TelemetryFieldOptions) String() string            { return proto.CompactTextString(m) }
func (*TelemetryFieldOptions) ProtoMessage()               {}
func (*TelemetryFieldOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TelemetryFieldOptions) GetIsKey() bool {
	if m != nil && m.IsKey != nil {
		return *m.IsKey
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsTimestamp() bool {
	if m != nil && m.IsTimestamp != nil {
		return *m.IsTimestamp
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsCounter() bool {
	if m != nil && m.IsCounter != nil {
		return *m.IsCounter
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsGauge() bool {
	if m != nil && m.IsGauge != nil {
		return *m.IsGauge
	}
	return false
}

type TelemetryStream struct {
	SystemId         *string            `protobuf:"bytes,1,req,name=system_id" json:"system_id,omitempty"`
	ComponentId      *uint32            `protobuf:"varint,2,opt,name=component_id" json:"component_id,omitempty"`
	SubComponentId   *uint32            `protobuf:"varint,3,opt,name=sub_component_id" json:"sub_component_id,omitempty"`
	SensorName       *string            `protobuf:"bytes,4,opt,name=sensor_name" json:"sensor_name,omitempty"`
	SequenceNumber   *uint32            `protobuf:"varint,5,opt,name=sequence_number" json:"sequence_number,omitempty"`
	Timestamp        *uint64            `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	VersionMajor     *uint32            `protobuf:"varint,7,opt,name=version_major" json:"version_major,omitempty"`
	VersionMinor     *uint32            `protobuf:"varint,8,opt,name=version_minor" json:"version_minor,omitempty"`
	Ietf             *IETFSensors       `protobuf:"bytes,100,opt,name=ietf" json:"ietf,omitempty"`
	Enterprise       *EnterpriseSensors `protobuf:"bytes,101,opt,name=enterprise" json:"enterprise,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *TelemetryStream) Reset()                    { *m = TelemetryStream{} }
func (m *TelemetryStream) String() string            { return proto.CompactTextString(m) }
func (*TelemetryStream) ProtoMessage()               {}
func (*TelemetryStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TelemetryStream) GetSystemId() string {
	if m != nil && m.SystemId != nil {
		return *m.SystemId
	}
	return ""
}

func (m *TelemetryStream) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSubComponentId() uint32 {
	if m != nil && m.SubComponentId != nil {
		return *m.SubComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSensorName() string {
	if m != nil && m.SensorName != nil {
		return *m.SensorName
	}
	return ""
}

func (m *TelemetryStream) GetSequenceNumber() uint32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *TelemetryStream) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryStream) GetVersionMajor() uint32 {
	if m != nil && m.VersionMajor != nil {
		return *m.VersionMajor
	}
	return 0
}

func (m *TelemetryStream) GetVersionMinor() uint32 {
	if m != nil && m.VersionMinor != nil {
		return *m.VersionMinor
	}
	return 0
}

func (m *TelemetryStream) GetIetf() *IETFSensors {
	if m != nil {
		return m.Ietf
	}
	return nil
}

func (m *TelemetryStream) GetEnterprise() *EnterpriseSensors {
	if m != nil {
		return m.Enterprise
	}
	return nil
}

type IETFSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *IETFSensors) Reset()                    { *m = IETFSensors{} }
func (m *IETFSensors) String() string            { return proto.CompactTextString(m) }
func (*IETFSensors) ProtoMessage()               {}
func (*IETFSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

var extRange_IETFSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*IETFSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_IETFSensors
}

type EnterpriseSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *EnterpriseSensors) Reset()                    { *m = EnterpriseSensors{} }
func (m *EnterpriseSensors) String() string            { return proto.CompactTextString(m) }
func (*EnterpriseSensors) ProtoMessage()               {}
func (*EnterpriseSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

var extRange_EnterpriseSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*EnterpriseSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnterpriseSensors
}

type JuniperNetworksSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *JuniperNetworksSensors) Reset()                    { *m = JuniperNetworksSensors{} }
func (m *JuniperNetworksSensors) String() string            { return proto.CompactTextString(m) }
func (*JuniperNetworksSensors) ProtoMessage()               {}
func (*JuniperNetworksSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

var extRange_JuniperNetworksSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*JuniperNetworksSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_JuniperNetworksSensors
}

var E_TelemetryOptions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*TelemetryFieldOptions)(nil),
	Field:         1024,
	Name:          "telemetry_options",
	Tag:           "bytes,1024,opt,name=telemetry_options",
	Filename:      "telemetry_top.proto",
}

var E_JuniperNetworks = &proto.ExtensionDesc{
	ExtendedType:  (*EnterpriseSensors)(nil),
	ExtensionType: (*JuniperNetworksSensors)(nil),
	Field:         2636,
	Name:          "juniperNetworks",
	Tag:           "bytes,2636,opt,name=juniperNetworks",
	Filename:      "telemetry_top.proto",
}

func init() {
	proto.RegisterType((*TelemetryFieldOptions)(nil), "TelemetryFieldOptions")
	proto.RegisterType((*TelemetryStream)(nil), "TelemetryStream")
	proto.RegisterType((*IETFSensors)(nil), "IETFSensors")
	proto.RegisterType((*EnterpriseSensors)(nil), "EnterpriseSensors")
	proto.RegisterType((*JuniperNetworksSensors)(nil), "JuniperNetworksSensors")
	proto.RegisterExtension(E_TelemetryOptions)
	proto.RegisterExtension(E_JuniperNetworks)
}

func init() { proto.RegisterFile("telemetry_top.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x65, 0x37, 0x2d, 0xee, 0x24, 0x25, 0xe9, 0xd2, 0x36, 0xab, 0xa0, 0x0a, 0xcb, 0x07,
	0x14, 0x71, 0x70, 0x25, 0x8e, 0x3d, 0x21, 0xa4, 0x54, 0x82, 0x03, 0x08, 0xb5, 0xf7, 0x95, 0xe3,
	0x4c, 0xad, 0x6d, 0xb3, 0x7f, 0xd8, 0x59, 0x83, 0x72, 0xb3, 0xf8, 0x38, 0x7c, 0x16, 0x3e, 0x14,
	0xf2, 0xa6, 0x49, 0x5d, 0x25, 0x47, 0xbf, 0xf9, 0xbd, 0x9d, 0xb7, 0xde, 0x07, 0x6f, 0x3c, 0x2e,
	0x51, 0xa1, 0x77, 0x2b, 0xe1, 0x8d, 0xcd, 0xad, 0x33, 0xde, 0x4c, 0xd2, 0xca, 0x98, 0x6a, 0x89,
	0x57, 0xe1, 0x6b, 0x5e, 0xdf, 0x5f, 0x2d, 0x90, 0x4a, 0x27, 0xad, 0x37, 0x6e, 0x4d, 0x64, 0x25,
	0x9c, 0xdf, 0x6d, 0x8c, 0x37, 0x12, 0x97, 0x8b, 0xef, 0xd6, 0x4b, 0xa3, 0x89, 0xbd, 0x86, 0x23,
	0x49, 0xe2, 0x11, 0x57, 0x3c, 0x4a, 0xa3, 0x69, 0xc2, 0xce, 0x60, 0x20, 0x49, 0x78, 0xa9, 0x90,
	0x7c, 0xa1, 0x2c, 0x8f, 0x83, 0xca, 0x00, 0x24, 0x89, 0xd2, 0xd4, 0xda, 0xa3, 0xe3, 0x07, 0x41,
	0x1b, 0x41, 0x22, 0x49, 0x54, 0x45, 0x5d, 0x21, 0xef, 0xb5, 0x4a, 0xf6, 0x37, 0x86, 0xe1, 0x76,
	0xcb, 0xad, 0x77, 0x58, 0x28, 0xc6, 0xe1, 0x98, 0x56, 0xe4, 0x51, 0x09, 0xb9, 0xe0, 0x51, 0x1a,
	0x4f, 0x8f, 0x3f, 0x1f, 0xfe, 0xf9, 0x14, 0x27, 0x11, 0x7b, 0x0b, 0x83, 0xd2, 0x28, 0x6b, 0x34,
	0x6a, 0xdf, 0x0e, 0xdb, 0x4d, 0x27, 0x9b, 0xe1, 0x3b, 0x18, 0x51, 0x3d, 0x17, 0x2f, 0x80, 0x83,
	0x2e, 0x30, 0x81, 0x3e, 0xa1, 0x26, 0xe3, 0x84, 0x2e, 0xd4, 0x3a, 0xc0, 0xf6, 0xe4, 0x31, 0x0c,
	0x09, 0x7f, 0xd6, 0xa8, 0x4b, 0x14, 0xba, 0x56, 0x73, 0x74, 0xfc, 0xb0, 0xf5, 0xb6, 0x61, 0x9e,
	0x6f, 0x76, 0x94, 0x46, 0xd3, 0x5e, 0xb0, 0x8c, 0x22, 0x76, 0x0e, 0x27, 0xbf, 0xd0, 0x91, 0x34,
	0x5a, 0xa8, 0xe2, 0xc1, 0x38, 0xfe, 0x2a, 0x18, 0xba, 0xb2, 0xd4, 0xc6, 0xf1, 0x24, 0xc8, 0x13,
	0xe8, 0x49, 0xf4, 0xf7, 0x7c, 0x91, 0x46, 0xd3, 0xfe, 0xc7, 0x41, 0xfe, 0x65, 0x76, 0x77, 0x73,
	0x1b, 0xd2, 0x10, 0x7b, 0x0f, 0x80, 0xed, 0x5f, 0xb2, 0x4e, 0x12, 0x72, 0x0c, 0x04, 0xcb, 0x67,
	0x5b, 0xe9, 0x89, 0xcb, 0xc6, 0xd0, 0xef, 0xd8, 0x3e, 0x24, 0x49, 0x34, 0x6a, 0x9a, 0xa6, 0x89,
	0xb3, 0x4b, 0x38, 0xdd, 0xa1, 0x3b, 0xe3, 0x0c, 0x2e, 0xbe, 0xd6, 0x5a, 0x5a, 0x74, 0xdf, 0xd0,
	0xff, 0x36, 0xee, 0x91, 0x76, 0x98, 0xeb, 0x1f, 0x70, 0xfa, 0x5c, 0x13, 0xf3, 0xf4, 0xd2, 0x97,
	0xf9, 0xba, 0x25, 0xf9, 0xa6, 0x25, 0x79, 0xb7, 0x08, 0xbc, 0x49, 0x42, 0xd4, 0x8b, 0x7c, 0x6f,
	0x4f, 0xae, 0x67, 0x30, 0x7c, 0x78, 0xb9, 0x96, 0xed, 0xb9, 0x15, 0xff, 0x77, 0x16, 0x4e, 0x19,
	0xe7, 0xfb, 0x33, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x88, 0x28, 0x6a, 0xbf, 0x02, 0x00,
	0x00,
}
