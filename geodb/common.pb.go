// Copied from https://github.com/v2fly/v2ray-core/blob/42b166760b2ba8d984e514b830fcd44e23728e43/app/router/routercommon

package geodb

import (
	"reflect"
	"sync"

	_ "github.com/taodev/pkg/geodb/protoext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of domain value.
type Domain_Type int32

const (
	// The value is used as is.
	Domain_Plain Domain_Type = 0
	// The value is used as a regular expression.
	Domain_Regex Domain_Type = 1
	// The value is a root domain.
	Domain_RootDomain Domain_Type = 2
	// The value is a domain.
	Domain_Full Domain_Type = 3
)

// Enum value maps for Domain_Type.
var (
	Domain_Type_name = map[int32]string{
		0: "Plain",
		1: "Regex",
		2: "RootDomain",
		3: "Full",
	}
	Domain_Type_value = map[string]int32{
		"Plain":      0,
		"Regex":      1,
		"RootDomain": 2,
		"Full":       3,
	}
)

func (x Domain_Type) Enum() *Domain_Type {
	p := new(Domain_Type)
	*p = x
	return p
}

func (x Domain_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Domain_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_app_router_routercommon_common_proto_enumTypes[0].Descriptor()
}

func (Domain_Type) Type() protoreflect.EnumType {
	return &file_app_router_routercommon_common_proto_enumTypes[0]
}

func (x Domain_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Domain_Type.Descriptor instead.
func (Domain_Type) EnumDescriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0, 0}
}

// Domain for routing decision.
type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain matching type.
	Type Domain_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.router.routercommon.Domain_Type" json:"type,omitempty"`
	// Domain value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Attributes of this domain. May be used for filtering.
	Attribute []*Domain_Attribute `protobuf:"bytes,3,rep,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetType() Domain_Type {
	if x != nil {
		return x.Type
	}
	return Domain_Plain
}

func (x *Domain) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Domain) GetAttribute() []*Domain_Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// IP for routing decision, in CIDR form.
type CIDR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP address, should be either 4 or 16 bytes.
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Number of leading ones in the network mask.
	Prefix uint32 `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	IpAddr string `protobuf:"bytes,68000,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
}

func (x *CIDR) Reset() {
	*x = CIDR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIDR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIDR) ProtoMessage() {}

func (x *CIDR) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIDR.ProtoReflect.Descriptor instead.
func (*CIDR) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{1}
}

func (x *CIDR) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *CIDR) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

func (x *CIDR) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

type GeoIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode  string  `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Cidr         []*CIDR `protobuf:"bytes,2,rep,name=cidr,proto3" json:"cidr,omitempty"`
	InverseMatch bool    `protobuf:"varint,3,opt,name=inverse_match,json=inverseMatch,proto3" json:"inverse_match,omitempty"`
	// resource_hash instruct simplified config converter to load domain from geo file.
	ResourceHash []byte `protobuf:"bytes,4,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash,omitempty"`
	Code         string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	FilePath     string `protobuf:"bytes,68000,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *GeoIP) Reset() {
	*x = GeoIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIP) ProtoMessage() {}

func (x *GeoIP) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIP.ProtoReflect.Descriptor instead.
func (*GeoIP) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{2}
}

func (x *GeoIP) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GeoIP) GetCidr() []*CIDR {
	if x != nil {
		return x.Cidr
	}
	return nil
}

func (x *GeoIP) GetInverseMatch() bool {
	if x != nil {
		return x.InverseMatch
	}
	return false
}

func (x *GeoIP) GetResourceHash() []byte {
	if x != nil {
		return x.ResourceHash
	}
	return nil
}

func (x *GeoIP) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GeoIP) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GeoIPList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry []*GeoIP `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GeoIPList) Reset() {
	*x = GeoIPList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIPList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIPList) ProtoMessage() {}

func (x *GeoIPList) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIPList.ProtoReflect.Descriptor instead.
func (*GeoIPList) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{3}
}

func (x *GeoIPList) GetEntry() []*GeoIP {
	if x != nil {
		return x.Entry
	}
	return nil
}

type GeoSite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode string    `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Domain      []*Domain `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	// resource_hash instruct simplified config converter to load domain from geo file.
	ResourceHash []byte `protobuf:"bytes,3,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash,omitempty"`
	Code         string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	FilePath     string `protobuf:"bytes,68000,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *GeoSite) Reset() {
	*x = GeoSite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoSite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSite) ProtoMessage() {}

func (x *GeoSite) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoSite.ProtoReflect.Descriptor instead.
func (*GeoSite) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{4}
}

func (x *GeoSite) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GeoSite) GetDomain() []*Domain {
	if x != nil {
		return x.Domain
	}
	return nil
}

func (x *GeoSite) GetResourceHash() []byte {
	if x != nil {
		return x.ResourceHash
	}
	return nil
}

func (x *GeoSite) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GeoSite) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GeoSiteList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry []*GeoSite `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GeoSiteList) Reset() {
	*x = GeoSiteList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoSiteList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSiteList) ProtoMessage() {}

func (x *GeoSiteList) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoSiteList.ProtoReflect.Descriptor instead.
func (*GeoSiteList) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{5}
}

func (x *GeoSiteList) GetEntry() []*GeoSite {
	if x != nil {
		return x.Entry
	}
	return nil
}

type Domain_Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to TypedValue:
	//
	//	*Domain_Attribute_BoolValue
	//	*Domain_Attribute_IntValue
	TypedValue isDomain_Attribute_TypedValue `protobuf_oneof:"typed_value"`
}

func (x *Domain_Attribute) Reset() {
	*x = Domain_Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_router_routercommon_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain_Attribute) ProtoMessage() {}

func (x *Domain_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_routercommon_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain_Attribute.ProtoReflect.Descriptor instead.
func (*Domain_Attribute) Descriptor() ([]byte, []int) {
	return file_app_router_routercommon_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Domain_Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *Domain_Attribute) GetTypedValue() isDomain_Attribute_TypedValue {
	if m != nil {
		return m.TypedValue
	}
	return nil
}

func (x *Domain_Attribute) GetBoolValue() bool {
	if x, ok := x.GetTypedValue().(*Domain_Attribute_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Domain_Attribute) GetIntValue() int64 {
	if x, ok := x.GetTypedValue().(*Domain_Attribute_IntValue); ok {
		return x.IntValue
	}
	return 0
}

type isDomain_Attribute_TypedValue interface {
	isDomain_Attribute_TypedValue()
}

type Domain_Attribute_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Domain_Attribute_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

func (*Domain_Attribute_BoolValue) isDomain_Attribute_TypedValue() {}

func (*Domain_Attribute_IntValue) isDomain_Attribute_TypedValue() {}

var File_app_router_routercommon_common_proto protoreflect.FileDescriptor

var file_app_router_routercommon_common_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a,
	0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x6c, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x6c, 0x61, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x75, 0x6c, 0x6c, 0x10, 0x03, 0x22, 0x53, 0x0a, 0x04,
	0x43, 0x49, 0x44, 0x52, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x23, 0x0a, 0x07,
	0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0xa0, 0x93, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0x82, 0xb5, 0x18, 0x04, 0x3a, 0x02, 0x69, 0x70, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x22, 0xfa, 0x01, 0x0a, 0x05, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3c,
	0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0xa0, 0x93, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0x82, 0xb5, 0x18, 0x0f, 0x32, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4c,
	0x0a, 0x09, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x6f, 0x49, 0x50, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xdd, 0x01, 0x0a,
	0x07, 0x47, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0xa0, 0x93, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0x82,
	0xb5, 0x18, 0x0f, 0x32, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x50, 0x0a, 0x0b,
	0x47, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x87,
	0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xaa, 0x02, 0x22, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_router_routercommon_common_proto_rawDescOnce sync.Once
	file_app_router_routercommon_common_proto_rawDescData = file_app_router_routercommon_common_proto_rawDesc
)

func file_app_router_routercommon_common_proto_rawDescGZIP() []byte {
	file_app_router_routercommon_common_proto_rawDescOnce.Do(func() {
		file_app_router_routercommon_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_router_routercommon_common_proto_rawDescData)
	})
	return file_app_router_routercommon_common_proto_rawDescData
}

var file_app_router_routercommon_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_router_routercommon_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_router_routercommon_common_proto_goTypes = []interface{}{
	(Domain_Type)(0),         // 0: v2ray.core.app.router.routercommon.Domain.Type
	(*Domain)(nil),           // 1: v2ray.core.app.router.routercommon.Domain
	(*CIDR)(nil),             // 2: v2ray.core.app.router.routercommon.CIDR
	(*GeoIP)(nil),            // 3: v2ray.core.app.router.routercommon.GeoIP
	(*GeoIPList)(nil),        // 4: v2ray.core.app.router.routercommon.GeoIPList
	(*GeoSite)(nil),          // 5: v2ray.core.app.router.routercommon.GeoSite
	(*GeoSiteList)(nil),      // 6: v2ray.core.app.router.routercommon.GeoSiteList
	(*Domain_Attribute)(nil), // 7: v2ray.core.app.router.routercommon.Domain.Attribute
}
var file_app_router_routercommon_common_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.router.routercommon.Domain.type:type_name -> v2ray.core.app.router.routercommon.Domain.Type
	7, // 1: v2ray.core.app.router.routercommon.Domain.attribute:type_name -> v2ray.core.app.router.routercommon.Domain.Attribute
	2, // 2: v2ray.core.app.router.routercommon.GeoIP.cidr:type_name -> v2ray.core.app.router.routercommon.CIDR
	3, // 3: v2ray.core.app.router.routercommon.GeoIPList.entry:type_name -> v2ray.core.app.router.routercommon.GeoIP
	1, // 4: v2ray.core.app.router.routercommon.GeoSite.domain:type_name -> v2ray.core.app.router.routercommon.Domain
	5, // 5: v2ray.core.app.router.routercommon.GeoSiteList.entry:type_name -> v2ray.core.app.router.routercommon.GeoSite
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_app_router_routercommon_common_proto_init() }
func file_app_router_routercommon_common_proto_init() {
	if File_app_router_routercommon_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_router_routercommon_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_app_router_routercommon_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIDR); i {
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
		file_app_router_routercommon_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIP); i {
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
		file_app_router_routercommon_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIPList); i {
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
		file_app_router_routercommon_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoSite); i {
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
		file_app_router_routercommon_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoSiteList); i {
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
		file_app_router_routercommon_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain_Attribute); i {
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
	file_app_router_routercommon_common_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Domain_Attribute_BoolValue)(nil),
		(*Domain_Attribute_IntValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_router_routercommon_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_router_routercommon_common_proto_goTypes,
		DependencyIndexes: file_app_router_routercommon_common_proto_depIdxs,
		EnumInfos:         file_app_router_routercommon_common_proto_enumTypes,
		MessageInfos:      file_app_router_routercommon_common_proto_msgTypes,
	}.Build()
	File_app_router_routercommon_common_proto = out.File
	file_app_router_routercommon_common_proto_rawDesc = nil
	file_app_router_routercommon_common_proto_goTypes = nil
	file_app_router_routercommon_common_proto_depIdxs = nil
}
