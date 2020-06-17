// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/failover.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//
//Failover configuration for an upstream.
//
//Failover allows for optional fallback endpoints in the case that the primary set of endpoints is deemed
//unhealthy. As failover requires knowledge of the health of each set of endpoints, active or passive
//health checks must be configured on an upstream using failover in order for it to work properly.
//
//Failover closely resembles the Envoy config which this is translated to, with one notable exception.
//The priorities are not defined on the `LocalityLbEndpoints` but rather inferred from the list of
//`PrioritizedLocality`. More information on envoy prioritization can be found
//[here](https://www.envoyproxy.io/docs/envoy/v1.14.1/intro/arch_overview/upstream/load_balancing/priority#arch-overview-load-balancing-priority-levels).
//In practice this means that the priority of a given set of `LocalityLbEndpoints` is determined by its index in
//the list, first being `0` through `n-1`.
//
type Failover struct {
	// Identifies where the parent upstream hosts run.
	Locality *Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// PrioritizedLocality is an implicitly prioritized list of lists of `LocalityLbEndpoints`. The priority of each
	// list of `LocalityLbEndpoints` is determined by its index in the list.
	PrioritizedLocalities []*Failover_PrioritizedLocality `protobuf:"bytes,2,rep,name=prioritized_localities,json=prioritizedLocalities,proto3" json:"prioritized_localities,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *Failover) Reset()         { *m = Failover{} }
func (m *Failover) String() string { return proto.CompactTextString(m) }
func (*Failover) ProtoMessage()    {}
func (*Failover) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{0}
}
func (m *Failover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failover.Unmarshal(m, b)
}
func (m *Failover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failover.Marshal(b, m, deterministic)
}
func (m *Failover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failover.Merge(m, src)
}
func (m *Failover) XXX_Size() int {
	return xxx_messageInfo_Failover.Size(m)
}
func (m *Failover) XXX_DiscardUnknown() {
	xxx_messageInfo_Failover.DiscardUnknown(m)
}

var xxx_messageInfo_Failover proto.InternalMessageInfo

func (m *Failover) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *Failover) GetPrioritizedLocalities() []*Failover_PrioritizedLocality {
	if m != nil {
		return m.PrioritizedLocalities
	}
	return nil
}

type Failover_PrioritizedLocality struct {
	LocalityEndpoints    []*LocalityLbEndpoints `protobuf:"bytes,1,rep,name=locality_endpoints,json=localityEndpoints,proto3" json:"locality_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Failover_PrioritizedLocality) Reset()         { *m = Failover_PrioritizedLocality{} }
func (m *Failover_PrioritizedLocality) String() string { return proto.CompactTextString(m) }
func (*Failover_PrioritizedLocality) ProtoMessage()    {}
func (*Failover_PrioritizedLocality) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{0, 0}
}
func (m *Failover_PrioritizedLocality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failover_PrioritizedLocality.Unmarshal(m, b)
}
func (m *Failover_PrioritizedLocality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failover_PrioritizedLocality.Marshal(b, m, deterministic)
}
func (m *Failover_PrioritizedLocality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failover_PrioritizedLocality.Merge(m, src)
}
func (m *Failover_PrioritizedLocality) XXX_Size() int {
	return xxx_messageInfo_Failover_PrioritizedLocality.Size(m)
}
func (m *Failover_PrioritizedLocality) XXX_DiscardUnknown() {
	xxx_messageInfo_Failover_PrioritizedLocality.DiscardUnknown(m)
}

var xxx_messageInfo_Failover_PrioritizedLocality proto.InternalMessageInfo

func (m *Failover_PrioritizedLocality) GetLocalityEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.LocalityEndpoints
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies where the parent upstream hosts run.
	Locality *Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []*LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	// Optional: Per priority/region/zone/sub_zone weight; at least 1. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities at the same priority level to produce the effective percentage
	// of traffic for the locality.
	LoadBalancingWeight  *types.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{1}
}
func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *types.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	// Address (hostname or IP)
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Port the instance is listening on
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	// This takes into effect only for upstreams with active health checking enabled
	HealthCheckConfig *LbEndpoint_HealthCheckConfig `protobuf:"bytes,3,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	UpstreamSslConfig *UpstreamSslConfig            `protobuf:"bytes,4,opt,name=upstream_ssl_config,json=upstreamSslConfig,proto3" json:"upstream_ssl_config,omitempty"`
	// The optional load balancing weight of the upstream host; at least 1.
	// Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	LoadBalancingWeight  *types.UInt32Value `protobuf:"bytes,5,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{2}
}
func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

func (m *LbEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LbEndpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *LbEndpoint) GetHealthCheckConfig() *LbEndpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

func (m *LbEndpoint) GetUpstreamSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.UpstreamSslConfig
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *types.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// The optional health check configuration.
type LbEndpoint_HealthCheckConfig struct {
	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue uint32 `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	// By default, the host header for L7 health checks is controlled by cluster level configuration. Setting this
	// to a non-empty value allows overriding the cluster level configuration for a specific endpoint.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LbEndpoint_HealthCheckConfig) Reset()         { *m = LbEndpoint_HealthCheckConfig{} }
func (m *LbEndpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint_HealthCheckConfig) ProtoMessage()    {}
func (*LbEndpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{2, 0}
}
func (m *LbEndpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *LbEndpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *LbEndpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint_HealthCheckConfig.Merge(m, src)
}
func (m *LbEndpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint_HealthCheckConfig.Size(m)
}
func (m *LbEndpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *LbEndpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

func (m *LbEndpoint_HealthCheckConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	// Region this zone belongs to.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. The meaning of zone
	// is context dependent, e.g. `Availability Zone (AZ)
	// <https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html>`_
	// on AWS, `Zone <https://cloud.google.com/compute/docs/regions-zones/>`_ on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone              string   `protobuf:"bytes,3,opt,name=sub_zone,json=subZone,proto3" json:"sub_zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Locality) Reset()         { *m = Locality{} }
func (m *Locality) String() string { return proto.CompactTextString(m) }
func (*Locality) ProtoMessage()    {}
func (*Locality) Descriptor() ([]byte, []int) {
	return fileDescriptor_78ccbfab63a57f32, []int{3}
}
func (m *Locality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Locality.Unmarshal(m, b)
}
func (m *Locality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Locality.Marshal(b, m, deterministic)
}
func (m *Locality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locality.Merge(m, src)
}
func (m *Locality) XXX_Size() int {
	return xxx_messageInfo_Locality.Size(m)
}
func (m *Locality) XXX_DiscardUnknown() {
	xxx_messageInfo_Locality.DiscardUnknown(m)
}

var xxx_messageInfo_Locality proto.InternalMessageInfo

func (m *Locality) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Locality) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Locality) GetSubZone() string {
	if m != nil {
		return m.SubZone
	}
	return ""
}

func init() {
	proto.RegisterType((*Failover)(nil), "gloo.solo.io.Failover")
	proto.RegisterType((*Failover_PrioritizedLocality)(nil), "gloo.solo.io.Failover.PrioritizedLocality")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "gloo.solo.io.LocalityLbEndpoints")
	proto.RegisterType((*LbEndpoint)(nil), "gloo.solo.io.LbEndpoint")
	proto.RegisterType((*LbEndpoint_HealthCheckConfig)(nil), "gloo.solo.io.LbEndpoint.HealthCheckConfig")
	proto.RegisterType((*Locality)(nil), "gloo.solo.io.Locality")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/failover.proto", fileDescriptor_78ccbfab63a57f32)
}

var fileDescriptor_78ccbfab63a57f32 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd6, 0x32, 0x5a, 0x77, 0xbb, 0xa8, 0xcb, 0xaa, 0x50, 0xc1, 0x28, 0xe5, 0xa6, 0x9a,
	0x44, 0x02, 0xdd, 0x1d, 0xbb, 0xdb, 0x04, 0x02, 0x69, 0x82, 0x12, 0x34, 0x90, 0x7a, 0x13, 0x39,
	0xa9, 0xeb, 0x98, 0xb9, 0x39, 0x91, 0xed, 0x74, 0x63, 0x6f, 0x00, 0x4f, 0xc2, 0x23, 0xf0, 0x28,
	0x5c, 0xf3, 0x0e, 0xdc, 0xa3, 0x38, 0x3f, 0xeb, 0xba, 0x56, 0x42, 0xbb, 0xf3, 0xf9, 0x7c, 0xbe,
	0xef, 0x7c, 0xfe, 0x4e, 0x14, 0x74, 0xc4, 0xb8, 0x8e, 0xd2, 0xc0, 0x09, 0x61, 0xee, 0x2a, 0x10,
	0xf0, 0x9c, 0x83, 0xcb, 0x04, 0x80, 0x9b, 0x48, 0xf8, 0x4a, 0x43, 0xad, 0xf2, 0x8a, 0x24, 0xdc,
	0x5d, 0xbc, 0x74, 0x67, 0x84, 0x0b, 0x58, 0x50, 0xe9, 0x24, 0x12, 0x34, 0xe0, 0x9d, 0xec, 0xce,
	0xc9, 0x68, 0x0e, 0x87, 0xde, 0x03, 0x06, 0x0c, 0xcc, 0x85, 0x9b, 0x9d, 0xf2, 0x9e, 0x1e, 0xa6,
	0x97, 0x3a, 0x07, 0xe9, 0xa5, 0x2e, 0xb0, 0x7d, 0x06, 0xc0, 0x04, 0x75, 0x4d, 0x15, 0xa4, 0x33,
	0xf7, 0x42, 0x92, 0x24, 0xa1, 0x52, 0x15, 0xf7, 0xcf, 0x36, 0x3b, 0x50, 0x4a, 0xe4, 0x4d, 0x83,
	0x1f, 0x5b, 0xa8, 0xf1, 0xa6, 0xf0, 0x83, 0x47, 0xa8, 0x21, 0x20, 0x24, 0x82, 0xeb, 0x6f, 0xb6,
	0xd5, 0xb7, 0x86, 0xad, 0x51, 0xd7, 0x59, 0x36, 0xe7, 0x9c, 0x16, 0xb7, 0x5e, 0xd5, 0x87, 0x09,
	0xea, 0x26, 0x92, 0x83, 0xe4, 0x9a, 0x5f, 0xd1, 0xa9, 0x5f, 0xe0, 0x9c, 0x2a, 0x7b, 0xab, 0x5f,
	0x1b, 0xb6, 0x46, 0x07, 0x37, 0x15, 0xca, 0x59, 0xce, 0xf8, 0x9a, 0x54, 0xa9, 0xee, 0x25, 0xb7,
	0x40, 0x4e, 0x55, 0x8f, 0xa1, 0xce, 0x9a, 0x6e, 0x3c, 0x46, 0xb8, 0x74, 0xe1, 0xd3, 0x78, 0x9a,
	0x00, 0x8f, 0xb5, 0xb2, 0x2d, 0x33, 0xf5, 0xe9, 0x7a, 0xdf, 0xa7, 0xc1, 0xeb, 0xb2, 0xd1, 0x6b,
	0x97, 0xe4, 0x0a, 0x1a, 0xfc, 0xb6, 0x50, 0x67, 0x4d, 0xeb, 0x9d, 0x72, 0x39, 0x42, 0x3b, 0x22,
	0x58, 0xf2, 0x95, 0xa7, 0x61, 0xaf, 0xf0, 0xaa, 0x21, 0x5e, 0x4b, 0x2c, 0x0d, 0x1c, 0xa3, 0x3d,
	0x01, 0x64, 0xea, 0x07, 0x44, 0x90, 0x38, 0xe4, 0x31, 0xf3, 0x2f, 0x28, 0x67, 0x91, 0xb6, 0x6b,
	0x66, 0xfa, 0x23, 0x27, 0x5f, 0xbd, 0x53, 0xae, 0xde, 0x39, 0x7b, 0x17, 0xeb, 0xc3, 0xd1, 0x67,
	0x22, 0x52, 0xea, 0x75, 0x32, 0xea, 0x71, 0xc9, 0xfc, 0x62, 0x88, 0x83, 0xef, 0x35, 0x84, 0xae,
	0xa7, 0x61, 0x1b, 0xdd, 0x27, 0xd3, 0xa9, 0xa4, 0x4a, 0x99, 0x07, 0x35, 0xbd, 0xb2, 0xc4, 0x18,
	0xd5, 0x13, 0x90, 0xda, 0xde, 0xea, 0x5b, 0xc3, 0x5d, 0xcf, 0x9c, 0xf1, 0x04, 0x75, 0x22, 0x4a,
	0x84, 0x8e, 0xfc, 0x30, 0xa2, 0xe1, 0xb9, 0x1f, 0x42, 0x3c, 0xe3, 0xac, 0x30, 0x73, 0xb0, 0xe9,
	0x49, 0xce, 0x5b, 0xc3, 0x39, 0xc9, 0x28, 0x27, 0x86, 0xe1, 0xb5, 0xa3, 0x55, 0x08, 0x7f, 0x40,
	0x9d, 0x34, 0x51, 0x5a, 0x52, 0x32, 0xf7, 0x95, 0x12, 0xa5, 0x76, 0xdd, 0x68, 0x3f, 0xb9, 0xa9,
	0x7d, 0x56, 0x34, 0x7e, 0x52, 0xa2, 0x14, 0x4c, 0x57, 0xa1, 0xcd, 0xd9, 0xdd, 0xbb, 0x63, 0x76,
	0xbd, 0xf7, 0xa8, 0x7d, 0xeb, 0x29, 0xf8, 0x31, 0x42, 0x59, 0x36, 0xfe, 0x22, 0xe3, 0x99, 0x10,
	0x77, 0xbd, 0x66, 0x86, 0x18, 0x21, 0xdc, 0x43, 0x8d, 0x08, 0x94, 0x8e, 0xc9, 0x9c, 0x9a, 0x28,
	0x9b, 0x5e, 0x55, 0x0f, 0x3e, 0xa2, 0x46, 0xf5, 0x11, 0x77, 0xd1, 0xb6, 0xa4, 0x8c, 0x43, 0x5c,
	0xec, 0xa1, 0xa8, 0xb2, 0x35, 0x5c, 0x41, 0x5c, 0x72, 0xcd, 0x19, 0x3f, 0x44, 0x0d, 0x95, 0x06,
	0xbe, 0xc1, 0x6b, 0xf9, 0xd6, 0x54, 0x1a, 0x4c, 0x20, 0xa6, 0xc7, 0xaf, 0x7e, 0xfd, 0xad, 0x5b,
	0x3f, 0xff, 0xec, 0x5b, 0x93, 0x17, 0xff, 0xf7, 0x2b, 0x4a, 0xce, 0x59, 0xf1, 0x33, 0x08, 0xb6,
	0x4d, 0x12, 0x87, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0x16, 0x4e, 0x73, 0xc5, 0x04, 0x00,
	0x00,
}

func (this *Failover) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failover)
	if !ok {
		that2, ok := that.(Failover)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Locality.Equal(that1.Locality) {
		return false
	}
	if len(this.PrioritizedLocalities) != len(that1.PrioritizedLocalities) {
		return false
	}
	for i := range this.PrioritizedLocalities {
		if !this.PrioritizedLocalities[i].Equal(that1.PrioritizedLocalities[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Failover_PrioritizedLocality) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failover_PrioritizedLocality)
	if !ok {
		that2, ok := that.(Failover_PrioritizedLocality)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.LocalityEndpoints) != len(that1.LocalityEndpoints) {
		return false
	}
	for i := range this.LocalityEndpoints {
		if !this.LocalityEndpoints[i].Equal(that1.LocalityEndpoints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LocalityLbEndpoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocalityLbEndpoints)
	if !ok {
		that2, ok := that.(LocalityLbEndpoints)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Locality.Equal(that1.Locality) {
		return false
	}
	if len(this.LbEndpoints) != len(that1.LbEndpoints) {
		return false
	}
	for i := range this.LbEndpoints {
		if !this.LbEndpoints[i].Equal(that1.LbEndpoints[i]) {
			return false
		}
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LbEndpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LbEndpoint)
	if !ok {
		that2, ok := that.(LbEndpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.HealthCheckConfig.Equal(that1.HealthCheckConfig) {
		return false
	}
	if !this.UpstreamSslConfig.Equal(that1.UpstreamSslConfig) {
		return false
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LbEndpoint_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LbEndpoint_HealthCheckConfig)
	if !ok {
		that2, ok := that.(LbEndpoint_HealthCheckConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PortValue != that1.PortValue {
		return false
	}
	if this.Hostname != that1.Hostname {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Locality) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Locality)
	if !ok {
		that2, ok := that.(Locality)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	if this.SubZone != that1.SubZone {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
