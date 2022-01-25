// Code generated by protoc-gen-go. DO NOT EDIT.
// source: port.proto

/*
Package port is a generated protocol buffer package.

It is generated from these files:
	port.proto

It has these top-level messages:
	Port
	InterfaceInfos
	QueueStats
	InterfaceStats
	IngressInterfaceErrors
	EgressInterfaceErrors
*/
package port

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import telemetry_top "github.com/influxdata/telegraf/plugins/parsers/juniperUDP/telemetry_top"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Port struct {
	InterfaceStats   []*InterfaceInfos `protobuf:"bytes,1,rep,name=interface_stats" json:"interface_stats,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Port) GetInterfaceStats() []*InterfaceInfos {
	if m != nil {
		return m.InterfaceStats
	}
	return nil
}

type InterfaceInfos struct {
	IfName                 *string                 `protobuf:"bytes,1,req,name=if_name" json:"if_name,omitempty"`
	InitTime               *uint64                 `protobuf:"varint,2,opt,name=init_time" json:"init_time,omitempty"`
	SnmpIfIndex            *uint32                 `protobuf:"varint,3,opt,name=snmp_if_index" json:"snmp_if_index,omitempty"`
	ParentAeName           *string                 `protobuf:"bytes,4,opt,name=parent_ae_name" json:"parent_ae_name,omitempty"`
	EgressQueueInfo        []*QueueStats           `protobuf:"bytes,5,rep,name=egress_queue_info" json:"egress_queue_info,omitempty"`
	IngressQueueInfo       []*QueueStats           `protobuf:"bytes,6,rep,name=ingress_queue_info" json:"ingress_queue_info,omitempty"`
	IngressStats           *InterfaceStats         `protobuf:"bytes,7,opt,name=ingress_stats" json:"ingress_stats,omitempty"`
	EgressStats            *InterfaceStats         `protobuf:"bytes,8,opt,name=egress_stats" json:"egress_stats,omitempty"`
	IngressErrors          *IngressInterfaceErrors `protobuf:"bytes,9,opt,name=ingress_errors" json:"ingress_errors,omitempty"`
	IfAdministrationStatus *string                 `protobuf:"bytes,10,opt,name=if_administration_status" json:"if_administration_status,omitempty"`
	IfOperationalStatus    *string                 `protobuf:"bytes,11,opt,name=if_operational_status" json:"if_operational_status,omitempty"`
	IfDescription          *string                 `protobuf:"bytes,12,opt,name=if_description" json:"if_description,omitempty"`
	IfTransitions          *uint64                 `protobuf:"varint,13,opt,name=if_transitions" json:"if_transitions,omitempty"`
	IfLastChange           *uint32                 `protobuf:"varint,14,opt,name=ifLastChange" json:"ifLastChange,omitempty"`
	IfHighSpeed            *uint32                 `protobuf:"varint,15,opt,name=ifHighSpeed" json:"ifHighSpeed,omitempty"`
	EgressErrors           *EgressInterfaceErrors  `protobuf:"bytes,16,opt,name=egress_errors" json:"egress_errors,omitempty"`
	XXX_unrecognized       []byte                  `json:"-"`
}

func (m *InterfaceInfos) Reset()                    { *m = InterfaceInfos{} }
func (m *InterfaceInfos) String() string            { return proto.CompactTextString(m) }
func (*InterfaceInfos) ProtoMessage()               {}
func (*InterfaceInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InterfaceInfos) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *InterfaceInfos) GetInitTime() uint64 {
	if m != nil && m.InitTime != nil {
		return *m.InitTime
	}
	return 0
}

func (m *InterfaceInfos) GetSnmpIfIndex() uint32 {
	if m != nil && m.SnmpIfIndex != nil {
		return *m.SnmpIfIndex
	}
	return 0
}

func (m *InterfaceInfos) GetParentAeName() string {
	if m != nil && m.ParentAeName != nil {
		return *m.ParentAeName
	}
	return ""
}

func (m *InterfaceInfos) GetEgressQueueInfo() []*QueueStats {
	if m != nil {
		return m.EgressQueueInfo
	}
	return nil
}

func (m *InterfaceInfos) GetIngressQueueInfo() []*QueueStats {
	if m != nil {
		return m.IngressQueueInfo
	}
	return nil
}

func (m *InterfaceInfos) GetIngressStats() *InterfaceStats {
	if m != nil {
		return m.IngressStats
	}
	return nil
}

func (m *InterfaceInfos) GetEgressStats() *InterfaceStats {
	if m != nil {
		return m.EgressStats
	}
	return nil
}

func (m *InterfaceInfos) GetIngressErrors() *IngressInterfaceErrors {
	if m != nil {
		return m.IngressErrors
	}
	return nil
}

func (m *InterfaceInfos) GetIfAdministrationStatus() string {
	if m != nil && m.IfAdministrationStatus != nil {
		return *m.IfAdministrationStatus
	}
	return ""
}

func (m *InterfaceInfos) GetIfOperationalStatus() string {
	if m != nil && m.IfOperationalStatus != nil {
		return *m.IfOperationalStatus
	}
	return ""
}

func (m *InterfaceInfos) GetIfDescription() string {
	if m != nil && m.IfDescription != nil {
		return *m.IfDescription
	}
	return ""
}

func (m *InterfaceInfos) GetIfTransitions() uint64 {
	if m != nil && m.IfTransitions != nil {
		return *m.IfTransitions
	}
	return 0
}

func (m *InterfaceInfos) GetIfLastChange() uint32 {
	if m != nil && m.IfLastChange != nil {
		return *m.IfLastChange
	}
	return 0
}

func (m *InterfaceInfos) GetIfHighSpeed() uint32 {
	if m != nil && m.IfHighSpeed != nil {
		return *m.IfHighSpeed
	}
	return 0
}

func (m *InterfaceInfos) GetEgressErrors() *EgressInterfaceErrors {
	if m != nil {
		return m.EgressErrors
	}
	return nil
}

type QueueStats struct {
	QueueNumber         *uint32 `protobuf:"varint,1,opt,name=queue_number" json:"queue_number,omitempty"`
	Packets             *uint64 `protobuf:"varint,2,opt,name=packets" json:"packets,omitempty"`
	Bytes               *uint64 `protobuf:"varint,3,opt,name=bytes" json:"bytes,omitempty"`
	TailDropPackets     *uint64 `protobuf:"varint,4,opt,name=tail_drop_packets" json:"tail_drop_packets,omitempty"`
	RlDropPackets       *uint64 `protobuf:"varint,5,opt,name=rl_drop_packets" json:"rl_drop_packets,omitempty"`
	RlDropBytes         *uint64 `protobuf:"varint,6,opt,name=rl_drop_bytes" json:"rl_drop_bytes,omitempty"`
	RedDropPackets      *uint64 `protobuf:"varint,7,opt,name=red_drop_packets" json:"red_drop_packets,omitempty"`
	RedDropBytes        *uint64 `protobuf:"varint,8,opt,name=red_drop_bytes" json:"red_drop_bytes,omitempty"`
	AvgBufferOccupancy  *uint64 `protobuf:"varint,9,opt,name=avg_buffer_occupancy" json:"avg_buffer_occupancy,omitempty"`
	CurBufferOccupancy  *uint64 `protobuf:"varint,10,opt,name=cur_buffer_occupancy" json:"cur_buffer_occupancy,omitempty"`
	PeakBufferOccupancy *uint64 `protobuf:"varint,11,opt,name=peak_buffer_occupancy" json:"peak_buffer_occupancy,omitempty"`
	AllocatedBufferSize *uint64 `protobuf:"varint,12,opt,name=allocated_buffer_size" json:"allocated_buffer_size,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *QueueStats) Reset()                    { *m = QueueStats{} }
func (m *QueueStats) String() string            { return proto.CompactTextString(m) }
func (*QueueStats) ProtoMessage()               {}
func (*QueueStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueueStats) GetQueueNumber() uint32 {
	if m != nil && m.QueueNumber != nil {
		return *m.QueueNumber
	}
	return 0
}

func (m *QueueStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *QueueStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *QueueStats) GetTailDropPackets() uint64 {
	if m != nil && m.TailDropPackets != nil {
		return *m.TailDropPackets
	}
	return 0
}

func (m *QueueStats) GetRlDropPackets() uint64 {
	if m != nil && m.RlDropPackets != nil {
		return *m.RlDropPackets
	}
	return 0
}

func (m *QueueStats) GetRlDropBytes() uint64 {
	if m != nil && m.RlDropBytes != nil {
		return *m.RlDropBytes
	}
	return 0
}

func (m *QueueStats) GetRedDropPackets() uint64 {
	if m != nil && m.RedDropPackets != nil {
		return *m.RedDropPackets
	}
	return 0
}

func (m *QueueStats) GetRedDropBytes() uint64 {
	if m != nil && m.RedDropBytes != nil {
		return *m.RedDropBytes
	}
	return 0
}

func (m *QueueStats) GetAvgBufferOccupancy() uint64 {
	if m != nil && m.AvgBufferOccupancy != nil {
		return *m.AvgBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetCurBufferOccupancy() uint64 {
	if m != nil && m.CurBufferOccupancy != nil {
		return *m.CurBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetPeakBufferOccupancy() uint64 {
	if m != nil && m.PeakBufferOccupancy != nil {
		return *m.PeakBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetAllocatedBufferSize() uint64 {
	if m != nil && m.AllocatedBufferSize != nil {
		return *m.AllocatedBufferSize
	}
	return 0
}

type InterfaceStats struct {
	IfPkts             *uint64 `protobuf:"varint,1,opt,name=if_pkts" json:"if_pkts,omitempty"`
	IfOctets           *uint64 `protobuf:"varint,2,opt,name=if_octets" json:"if_octets,omitempty"`
	If_1SecPkts        *uint64 `protobuf:"varint,3,opt,name=if_1sec_pkts" json:"if_1sec_pkts,omitempty"`
	If_1SecOctets      *uint64 `protobuf:"varint,4,opt,name=if_1sec_octets" json:"if_1sec_octets,omitempty"`
	IfUcPkts           *uint64 `protobuf:"varint,5,opt,name=if_uc_pkts" json:"if_uc_pkts,omitempty"`
	IfMcPkts           *uint64 `protobuf:"varint,6,opt,name=if_mc_pkts" json:"if_mc_pkts,omitempty"`
	IfBcPkts           *uint64 `protobuf:"varint,7,opt,name=if_bc_pkts" json:"if_bc_pkts,omitempty"`
	IfError            *uint64 `protobuf:"varint,8,opt,name=if_error" json:"if_error,omitempty"`
	IfPausePkts        *uint64 `protobuf:"varint,9,opt,name=if_pause_pkts" json:"if_pause_pkts,omitempty"`
	IfUnknownProtoPkts *uint64 `protobuf:"varint,10,opt,name=if_unknown_proto_pkts" json:"if_unknown_proto_pkts,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *InterfaceStats) Reset()                    { *m = InterfaceStats{} }
func (m *InterfaceStats) String() string            { return proto.CompactTextString(m) }
func (*InterfaceStats) ProtoMessage()               {}
func (*InterfaceStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InterfaceStats) GetIfPkts() uint64 {
	if m != nil && m.IfPkts != nil {
		return *m.IfPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfOctets() uint64 {
	if m != nil && m.IfOctets != nil {
		return *m.IfOctets
	}
	return 0
}

func (m *InterfaceStats) GetIf_1SecPkts() uint64 {
	if m != nil && m.If_1SecPkts != nil {
		return *m.If_1SecPkts
	}
	return 0
}

func (m *InterfaceStats) GetIf_1SecOctets() uint64 {
	if m != nil && m.If_1SecOctets != nil {
		return *m.If_1SecOctets
	}
	return 0
}

func (m *InterfaceStats) GetIfUcPkts() uint64 {
	if m != nil && m.IfUcPkts != nil {
		return *m.IfUcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfMcPkts() uint64 {
	if m != nil && m.IfMcPkts != nil {
		return *m.IfMcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfBcPkts() uint64 {
	if m != nil && m.IfBcPkts != nil {
		return *m.IfBcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfError() uint64 {
	if m != nil && m.IfError != nil {
		return *m.IfError
	}
	return 0
}

func (m *InterfaceStats) GetIfPausePkts() uint64 {
	if m != nil && m.IfPausePkts != nil {
		return *m.IfPausePkts
	}
	return 0
}

func (m *InterfaceStats) GetIfUnknownProtoPkts() uint64 {
	if m != nil && m.IfUnknownProtoPkts != nil {
		return *m.IfUnknownProtoPkts
	}
	return 0
}

type IngressInterfaceErrors struct {
	IfErrors               *uint64 `protobuf:"varint,1,opt,name=if_errors" json:"if_errors,omitempty"`
	IfInQdrops             *uint64 `protobuf:"varint,2,opt,name=if_in_qdrops" json:"if_in_qdrops,omitempty"`
	IfInFrameErrors        *uint64 `protobuf:"varint,3,opt,name=if_in_frame_errors" json:"if_in_frame_errors,omitempty"`
	IfDiscards             *uint64 `protobuf:"varint,4,opt,name=if_discards" json:"if_discards,omitempty"`
	IfInRunts              *uint64 `protobuf:"varint,5,opt,name=if_in_runts" json:"if_in_runts,omitempty"`
	IfInL3Incompletes      *uint64 `protobuf:"varint,6,opt,name=if_in_l3_incompletes" json:"if_in_l3_incompletes,omitempty"`
	IfInL2ChanErrors       *uint64 `protobuf:"varint,7,opt,name=if_in_l2chan_errors" json:"if_in_l2chan_errors,omitempty"`
	IfInL2MismatchTimeouts *uint64 `protobuf:"varint,8,opt,name=if_in_l2_mismatch_timeouts" json:"if_in_l2_mismatch_timeouts,omitempty"`
	IfInFifoErrors         *uint64 `protobuf:"varint,9,opt,name=if_in_fifo_errors" json:"if_in_fifo_errors,omitempty"`
	IfInResourceErrors     *uint64 `protobuf:"varint,10,opt,name=if_in_resource_errors" json:"if_in_resource_errors,omitempty"`
	XXX_unrecognized       []byte  `json:"-"`
}

func (m *IngressInterfaceErrors) Reset()                    { *m = IngressInterfaceErrors{} }
func (m *IngressInterfaceErrors) String() string            { return proto.CompactTextString(m) }
func (*IngressInterfaceErrors) ProtoMessage()               {}
func (*IngressInterfaceErrors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IngressInterfaceErrors) GetIfErrors() uint64 {
	if m != nil && m.IfErrors != nil {
		return *m.IfErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInQdrops() uint64 {
	if m != nil && m.IfInQdrops != nil {
		return *m.IfInQdrops
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInFrameErrors() uint64 {
	if m != nil && m.IfInFrameErrors != nil {
		return *m.IfInFrameErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfDiscards() uint64 {
	if m != nil && m.IfDiscards != nil {
		return *m.IfDiscards
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInRunts() uint64 {
	if m != nil && m.IfInRunts != nil {
		return *m.IfInRunts
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL3Incompletes() uint64 {
	if m != nil && m.IfInL3Incompletes != nil {
		return *m.IfInL3Incompletes
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL2ChanErrors() uint64 {
	if m != nil && m.IfInL2ChanErrors != nil {
		return *m.IfInL2ChanErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL2MismatchTimeouts() uint64 {
	if m != nil && m.IfInL2MismatchTimeouts != nil {
		return *m.IfInL2MismatchTimeouts
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInFifoErrors() uint64 {
	if m != nil && m.IfInFifoErrors != nil {
		return *m.IfInFifoErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInResourceErrors() uint64 {
	if m != nil && m.IfInResourceErrors != nil {
		return *m.IfInResourceErrors
	}
	return 0
}

type EgressInterfaceErrors struct {
	IfErrors         *uint64 `protobuf:"varint,1,opt,name=if_errors" json:"if_errors,omitempty"`
	IfDiscards       *uint64 `protobuf:"varint,2,opt,name=if_discards" json:"if_discards,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EgressInterfaceErrors) Reset()                    { *m = EgressInterfaceErrors{} }
func (m *EgressInterfaceErrors) String() string            { return proto.CompactTextString(m) }
func (*EgressInterfaceErrors) ProtoMessage()               {}
func (*EgressInterfaceErrors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EgressInterfaceErrors) GetIfErrors() uint64 {
	if m != nil && m.IfErrors != nil {
		return *m.IfErrors
	}
	return 0
}

func (m *EgressInterfaceErrors) GetIfDiscards() uint64 {
	if m != nil && m.IfDiscards != nil {
		return *m.IfDiscards
	}
	return 0
}

var E_JnprInterfaceExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*Port)(nil),
	Field:         3,
	Name:          "jnpr_interface_ext",
	Tag:           "bytes,3,opt,name=jnpr_interface_ext",
	Filename:      "port.proto",
}

func init() {
	proto.RegisterType((*Port)(nil), "Port")
	proto.RegisterType((*InterfaceInfos)(nil), "InterfaceInfos")
	proto.RegisterType((*QueueStats)(nil), "QueueStats")
	proto.RegisterType((*InterfaceStats)(nil), "InterfaceStats")
	proto.RegisterType((*IngressInterfaceErrors)(nil), "IngressInterfaceErrors")
	proto.RegisterType((*EgressInterfaceErrors)(nil), "EgressInterfaceErrors")
	proto.RegisterExtension(E_JnprInterfaceExt)
}

func init() { proto.RegisterFile("port.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xd1, 0x72, 0x1b, 0x35,
	0x14, 0x86, 0x67, 0x1d, 0x3b, 0x89, 0x8f, 0x63, 0xa7, 0xd9, 0x24, 0xce, 0x92, 0xb6, 0xb0, 0xb8,
	0xa5, 0x98, 0x0b, 0x0c, 0x84, 0xbb, 0x5e, 0x31, 0x30, 0x99, 0x21, 0x0c, 0x30, 0x30, 0x79, 0x00,
	0x8d, 0xb2, 0x7b, 0x14, 0x0b, 0x7b, 0x25, 0x55, 0xd2, 0xd2, 0x86, 0x4b, 0xee, 0xfb, 0x26, 0x3c,
	0x17, 0xcf, 0xc1, 0xac, 0xb4, 0xeb, 0x6a, 0x77, 0xcb, 0x45, 0x6f, 0xcf, 0xff, 0x1d, 0x59, 0xe7,
	0xd7, 0x7f, 0xd6, 0x00, 0x4a, 0x6a, 0xbb, 0x52, 0x5a, 0x5a, 0x79, 0x79, 0x6a, 0x71, 0x8b, 0x05,
	0x5a, 0xfd, 0x40, 0xac, 0x54, 0xbe, 0xb8, 0xf8, 0x1a, 0x86, 0xbf, 0x49, 0x6d, 0xe3, 0x25, 0x1c,
	0x73, 0x61, 0x51, 0x33, 0x9a, 0x21, 0x31, 0x96, 0x5a, 0x93, 0x44, 0xe9, 0xde, 0x72, 0x72, 0x75,
	0xbc, 0xba, 0x69, 0xea, 0x37, 0x82, 0x49, 0xb3, 0x78, 0x3b, 0x84, 0x59, 0xbb, 0x14, 0xcf, 0xe1,
	0x80, 0x33, 0x22, 0x68, 0x81, 0x49, 0x94, 0x0e, 0x96, 0xe3, 0xef, 0x47, 0x7f, 0x7f, 0x37, 0x38,
	0x8c, 0xe2, 0x13, 0x18, 0x73, 0xc1, 0x2d, 0xb1, 0xbc, 0xc0, 0x64, 0x90, 0x46, 0xcb, 0x61, 0x7c,
	0x0e, 0x53, 0x23, 0x0a, 0x45, 0x38, 0x23, 0x5c, 0xe4, 0xf8, 0x26, 0xd9, 0x4b, 0xa3, 0xe5, 0x34,
	0x9e, 0xc3, 0x4c, 0x51, 0x8d, 0xc2, 0x12, 0x8a, 0xfe, 0xa0, 0x61, 0x1a, 0x2d, 0xc7, 0xf1, 0x0b,
	0x38, 0xc1, 0x7b, 0x8d, 0xc6, 0x90, 0x57, 0x25, 0x96, 0x48, 0xb8, 0x60, 0x32, 0x19, 0xb9, 0x8b,
	0x4d, 0x56, 0xbf, 0x57, 0xa5, 0xdb, 0xea, 0xae, 0xf1, 0xe7, 0x10, 0x73, 0xd1, 0x03, 0xf7, 0xfb,
	0xe0, 0x0b, 0x98, 0x36, 0xa0, 0x9f, 0xf2, 0x20, 0x8d, 0xda, 0x53, 0x7a, 0xee, 0x33, 0x38, 0xc2,
	0x10, 0x3b, 0x7c, 0x3f, 0xf6, 0x15, 0xcc, 0x9a, 0xe3, 0x50, 0x6b, 0xa9, 0x4d, 0x32, 0x76, 0xe0,
	0xc5, 0xea, 0xc6, 0x97, 0x77, 0xfc, 0xb5, 0x93, 0xe3, 0x14, 0x12, 0xce, 0x08, 0xcd, 0x0b, 0x2e,
	0xb8, 0xb1, 0x9a, 0x5a, 0x2e, 0x85, 0xfb, 0x89, 0xd2, 0x24, 0xe0, 0x46, 0x7e, 0x0a, 0xe7, 0x9c,
	0x11, 0xa9, 0xd0, 0x8b, 0x74, 0xdb, 0xc8, 0x13, 0x27, 0xcf, 0x61, 0xc6, 0x19, 0xc9, 0xd1, 0x64,
	0x9a, 0xab, 0x0a, 0x48, 0x8e, 0xea, 0xb6, 0xaa, 0x6e, 0x35, 0x15, 0x86, 0x57, 0x65, 0x93, 0x4c,
	0x2b, 0xc3, 0xdd, 0x53, 0x24, 0x51, 0x7c, 0x06, 0x47, 0x9c, 0xfd, 0x4c, 0x8d, 0xfd, 0x61, 0x4d,
	0xc5, 0x3d, 0x26, 0x33, 0x67, 0xfb, 0x29, 0x4c, 0x38, 0xfb, 0x91, 0xdf, 0xaf, 0x6f, 0x15, 0x62,
	0x9e, 0x1c, 0xbb, 0xe2, 0x97, 0x30, 0xc5, 0xd6, 0x48, 0x8f, 0xdc, 0x48, 0xf3, 0xd5, 0xf5, 0xfb,
	0x26, 0x5a, 0xbc, 0xdd, 0x03, 0x08, 0x0c, 0x7e, 0x0c, 0x47, 0xfe, 0x05, 0x44, 0x59, 0xdc, 0xa1,
	0x4e, 0xa2, 0xea, 0xcc, 0x26, 0x10, 0x73, 0x38, 0x50, 0x34, 0xdb, 0xa0, 0x35, 0x3e, 0x0e, 0xef,
	0x6e, 0x37, 0xba, 0x7b, 0xb0, 0x68, 0x5c, 0x1a, 0x76, 0xd5, 0x14, 0x4e, 0x2c, 0xe5, 0x5b, 0x92,
	0x6b, 0xa9, 0x48, 0xd3, 0x37, 0x0c, 0x89, 0x8f, 0xe1, 0x58, 0x77, 0xf4, 0x51, 0xa8, 0x3f, 0x81,
	0x69, 0xa3, 0xfb, 0xf3, 0xf7, 0x43, 0xf5, 0x13, 0x78, 0xa4, 0x31, 0x6f, 0xb7, 0x1f, 0x84, 0xc0,
	0x53, 0x98, 0xed, 0x00, 0xdf, 0x7f, 0x18, 0xca, 0xcf, 0xe0, 0x8c, 0xfe, 0x79, 0x4f, 0xee, 0x4a,
	0xc6, 0x50, 0x13, 0x99, 0x65, 0xa5, 0xa2, 0x22, 0x7b, 0x70, 0x11, 0xf0, 0x50, 0xea, 0xa0, 0xac,
	0xd4, 0x7d, 0x08, 0x42, 0xe8, 0x39, 0x9c, 0x2b, 0xa4, 0x9b, 0x3e, 0x35, 0xe9, 0x50, 0x74, 0xbb,
	0x95, 0x19, 0xb5, 0x98, 0x37, 0xa8, 0xe1, 0x7f, 0xa1, 0x4b, 0x40, 0x43, 0x2d, 0xfe, 0x19, 0x04,
	0xfb, 0xe9, 0xdf, 0xc4, 0xef, 0xa7, 0xda, 0xb8, 0xa5, 0x0e, 0x06, 0x48, 0x60, 0x5c, 0x45, 0x2d,
	0xb3, 0xbd, 0x07, 0x79, 0x5c, 0xc5, 0x85, 0x7c, 0x63, 0x30, 0xf3, 0x6d, 0x7b, 0xe1, 0x3d, 0x7c,
	0xd4, 0x9c, 0x58, 0xf7, 0x0e, 0x43, 0xf9, 0x23, 0x00, 0xce, 0x48, 0x59, 0x77, 0xb6, 0xde, 0xc3,
	0x4b, 0x45, 0x2d, 0xed, 0xf7, 0xa5, 0xbb, 0x5a, 0x6a, 0x3d, 0xc3, 0x05, 0x1c, 0x72, 0xe6, 0xc3,
	0xd8, 0x7e, 0x80, 0x27, 0x30, 0xad, 0xe6, 0xa2, 0xa5, 0x41, 0xdf, 0x36, 0x0e, 0xd5, 0xe7, 0x6e,
	0x91, 0x4a, 0xb1, 0x11, 0xf2, 0xb5, 0x20, 0xee, 0x73, 0xe7, 0x29, 0x08, 0xa8, 0xc5, 0xbf, 0x03,
	0x98, 0xff, 0xcf, 0xae, 0x7a, 0x7b, 0xea, 0x25, 0x88, 0xfa, 0xf6, 0x70, 0x41, 0x5e, 0x55, 0xd9,
	0xe8, 0x78, 0xf7, 0x29, 0xc4, 0x5e, 0x64, 0x9a, 0x16, 0xd8, 0xf4, 0xb7, 0x92, 0x7d, 0x59, 0xed,
	0x1d, 0xc9, 0xb9, 0xc9, 0xa8, 0xce, 0x3b, 0x99, 0xf6, 0x1a, 0x17, 0x44, 0x97, 0xa2, 0xeb, 0xdf,
	0x33, 0x38, 0xf3, 0xda, 0xf6, 0x5b, 0xc2, 0x45, 0x26, 0x0b, 0xb5, 0xc5, 0x5e, 0xac, 0x17, 0x70,
	0x5a, 0x43, 0x57, 0xd9, 0x9a, 0x8a, 0xe6, 0x02, 0x2d, 0x4b, 0xbf, 0x80, 0xcb, 0x86, 0x21, 0x05,
	0x37, 0x05, 0xb5, 0xd9, 0xda, 0x7d, 0xa6, 0x65, 0x69, 0x3b, 0x29, 0x4f, 0xe1, 0xa4, 0x1e, 0x87,
	0x33, 0x19, 0x7e, 0xe5, 0x3a, 0x46, 0x57, 0x37, 0x46, 0x23, 0x4b, 0x9d, 0xed, 0x66, 0x6e, 0x19,
	0xfd, 0x0b, 0x9c, 0x5f, 0x7f, 0xa0, 0xcd, 0x1d, 0x9b, 0x42, 0x97, 0x5f, 0xbe, 0x84, 0xf8, 0x0f,
	0xa1, 0x34, 0x79, 0xf7, 0xaf, 0x85, 0x6f, 0x6c, 0x7c, 0xb1, 0xfa, 0xa9, 0x14, 0x5c, 0xa1, 0xfe,
	0x15, 0xed, 0x6b, 0xa9, 0x37, 0xe6, 0x16, 0x85, 0x69, 0xec, 0x9f, 0x5c, 0x8d, 0x56, 0xd5, 0x9f,
	0xdd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xba, 0x5d, 0x1c, 0x16, 0x07, 0x00, 0x00,
}
