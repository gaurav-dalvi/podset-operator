// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/customer.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A customer.
type Customer struct {
	// The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Whether the customer is a manager.
	Manager *wrappers.BoolValue `protobuf:"bytes,12,opt,name=manager,proto3" json:"manager,omitempty"`
	// Whether the customer is a test account.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,13,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_0ab473c999471a98, []int{0}
}
func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (dst *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(dst, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Customer) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Customer) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *Customer) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Customer) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *Customer) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Customer) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Customer) GetAutoTaggingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.AutoTaggingEnabled
	}
	return nil
}

func (m *Customer) GetHasPartnersBadge() *wrappers.BoolValue {
	if m != nil {
		return m.HasPartnersBadge
	}
	return nil
}

func (m *Customer) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *Customer) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *Customer) GetCallReportingSetting() *CallReportingSetting {
	if m != nil {
		return m.CallReportingSetting
	}
	return nil
}

func (m *Customer) GetConversionTrackingSetting() *ConversionTrackingSetting {
	if m != nil {
		return m.ConversionTrackingSetting
	}
	return nil
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *wrappers.StringValue `protobuf:"bytes,9,opt,name=call_conversion_action,json=callConversionAction,proto3" json:"call_conversion_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallReportingSetting) Reset()         { *m = CallReportingSetting{} }
func (m *CallReportingSetting) String() string { return proto.CompactTextString(m) }
func (*CallReportingSetting) ProtoMessage()    {}
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_0ab473c999471a98, []int{1}
}
func (m *CallReportingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReportingSetting.Unmarshal(m, b)
}
func (m *CallReportingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReportingSetting.Marshal(b, m, deterministic)
}
func (dst *CallReportingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReportingSetting.Merge(dst, src)
}
func (m *CallReportingSetting) XXX_Size() int {
	return xxx_messageInfo_CallReportingSetting.Size(m)
}
func (m *CallReportingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReportingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CallReportingSetting proto.InternalMessageInfo

func (m *CallReportingSetting) GetCallReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallConversionReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.CallConversionAction
	}
	return nil
}

// A collection of customer-wide settings related to Google Ads Conversion
// Tracking.
type ConversionTrackingSetting struct {
	// The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3" json:"conversion_tracking_id,omitempty"`
	// The conversion tracking id of the customer's manager. This is set when the
	// customer is opted into cross account conversion tracking, and it overrides
	// conversion_tracking_id. This field can only be managed through the Google
	// Ads UI. This field is read-only.
	CrossAccountConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cross_account_conversion_tracking_id,json=crossAccountConversionTrackingId,proto3" json:"cross_account_conversion_tracking_id,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *ConversionTrackingSetting) Reset()         { *m = ConversionTrackingSetting{} }
func (m *ConversionTrackingSetting) String() string { return proto.CompactTextString(m) }
func (*ConversionTrackingSetting) ProtoMessage()    {}
func (*ConversionTrackingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_0ab473c999471a98, []int{2}
}
func (m *ConversionTrackingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionTrackingSetting.Unmarshal(m, b)
}
func (m *ConversionTrackingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionTrackingSetting.Marshal(b, m, deterministic)
}
func (dst *ConversionTrackingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionTrackingSetting.Merge(dst, src)
}
func (m *ConversionTrackingSetting) XXX_Size() int {
	return xxx_messageInfo_ConversionTrackingSetting.Size(m)
}
func (m *ConversionTrackingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionTrackingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionTrackingSetting proto.InternalMessageInfo

func (m *ConversionTrackingSetting) GetConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.ConversionTrackingId
	}
	return nil
}

func (m *ConversionTrackingSetting) GetCrossAccountConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.CrossAccountConversionTrackingId
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v0.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v0.resources.CallReportingSetting")
	proto.RegisterType((*ConversionTrackingSetting)(nil), "google.ads.googleads.v0.resources.ConversionTrackingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/customer.proto", fileDescriptor_customer_0ab473c999471a98)
}

var fileDescriptor_customer_0ab473c999471a98 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x6e, 0xd3, 0x3c,
	0x18, 0xc7, 0x95, 0xec, 0x7d, 0xf7, 0xe1, 0xb5, 0x63, 0x32, 0x65, 0xca, 0x36, 0x34, 0x6d, 0x83,
	0x49, 0x93, 0x90, 0xd2, 0x0a, 0x26, 0x10, 0x01, 0x0e, 0xd2, 0x0a, 0xc6, 0x10, 0x42, 0xa5, 0xdb,
	0x7a, 0x30, 0x55, 0x8a, 0x5c, 0xc7, 0xcd, 0xa2, 0x25, 0x76, 0x64, 0x3b, 0xe5, 0x43, 0x5c, 0x0d,
	0x9c, 0x71, 0x03, 0xdc, 0x03, 0xb7, 0xc0, 0x1d, 0x70, 0x09, 0x1c, 0xa1, 0x38, 0x71, 0xd8, 0xd4,
	0x6e, 0xe9, 0x51, 0xad, 0xe6, 0xf9, 0xfd, 0xfe, 0xf6, 0x93, 0x27, 0x09, 0x68, 0x05, 0x8c, 0x05,
	0x11, 0x69, 0x22, 0x5f, 0x34, 0xf3, 0x65, 0xb6, 0x1a, 0xb7, 0x9a, 0x9c, 0x08, 0x96, 0x72, 0x4c,
	0x44, 0x13, 0xa7, 0x42, 0xb2, 0x98, 0x70, 0x3b, 0xe1, 0x4c, 0x32, 0xb8, 0x93, 0x97, 0xd9, 0xc8,
	0x17, 0x76, 0x49, 0xd8, 0xe3, 0x96, 0x5d, 0x12, 0x1b, 0x5b, 0x85, 0x54, 0x01, 0xc3, 0x74, 0xd4,
	0xfc, 0xc0, 0x51, 0x92, 0x10, 0x2e, 0x72, 0xc5, 0xee, 0x8f, 0x05, 0xb0, 0xd8, 0x29, 0xac, 0xf0,
	0x1e, 0xa8, 0x6b, 0xd2, 0xa3, 0x28, 0x26, 0x96, 0xb1, 0x6d, 0xec, 0x2f, 0xf5, 0x6a, 0xfa, 0xcf,
	0x77, 0x28, 0x26, 0xf0, 0x01, 0x30, 0x43, 0xdf, 0x9a, 0xdb, 0x36, 0xf6, 0x97, 0x1f, 0x6e, 0x16,
	0xb1, 0xb6, 0xd6, 0xdb, 0x47, 0x54, 0x3e, 0x3e, 0xe8, 0xa3, 0x28, 0x25, 0x3d, 0x33, 0xf4, 0xe1,
	0x21, 0x58, 0xf5, 0x89, 0xc0, 0x3c, 0x4c, 0x64, 0x38, 0x2e, 0xa4, 0xff, 0x29, 0xf4, 0xee, 0x04,
	0x7a, 0x2c, 0x79, 0x48, 0x83, 0x9c, 0xbd, 0x75, 0x89, 0x52, 0xa9, 0x2e, 0xa8, 0xe3, 0x94, 0x73,
	0x42, 0xf1, 0x27, 0x0f, 0x33, 0x9f, 0x58, 0xff, 0xcf, 0x60, 0xa9, 0x69, 0xa4, 0xc3, 0x7c, 0x02,
	0x9f, 0x82, 0x25, 0x19, 0xc6, 0xc4, 0xfb, 0xcc, 0x28, 0xb1, 0xe6, 0x67, 0xc0, 0x17, 0xb3, 0xf2,
	0x33, 0x46, 0x09, 0xec, 0x82, 0x3b, 0x92, 0x23, 0x7c, 0x11, 0xd2, 0xc0, 0x4b, 0x79, 0xe4, 0x49,
	0x12, 0x27, 0x11, 0x92, 0xc4, 0x5a, 0x98, 0x41, 0x73, 0x5b, 0xa3, 0xa7, 0x3c, 0x3a, 0x29, 0x40,
	0xf8, 0x0a, 0xac, 0x8e, 0x42, 0x8a, 0x22, 0xa5, 0x13, 0xe9, 0x68, 0x14, 0x7e, 0xb4, 0x96, 0x67,
	0x90, 0xad, 0x28, 0xea, 0x94, 0x47, 0xc7, 0x8a, 0x81, 0x6f, 0x41, 0x03, 0xa5, 0x92, 0x79, 0x12,
	0x05, 0x41, 0xb6, 0x3b, 0x42, 0xd1, 0x30, 0x22, 0xbe, 0xb5, 0xa8, 0x5c, 0x1b, 0x13, 0xae, 0x36,
	0x63, 0x51, 0x6e, 0x82, 0x19, 0x77, 0x92, 0x63, 0x2f, 0x73, 0x0a, 0xbe, 0x06, 0xf0, 0x1c, 0x09,
	0x2f, 0x41, 0x5c, 0x52, 0xc2, 0x85, 0x37, 0x44, 0x7e, 0x40, 0xac, 0xa5, 0x4a, 0xd7, 0xea, 0x39,
	0x12, 0xdd, 0x02, 0x6a, 0x67, 0x0c, 0x3c, 0x00, 0x0b, 0x31, 0xa2, 0x28, 0x20, 0xdc, 0xaa, 0x55,
	0xe2, 0xba, 0x14, 0xbe, 0x00, 0x35, 0x49, 0x84, 0xf4, 0x10, 0xc6, 0x2c, 0xa5, 0xd2, 0xaa, 0x57,
	0xa2, 0xcb, 0x59, 0xbd, 0x9b, 0x97, 0xc3, 0x18, 0xac, 0x61, 0x14, 0x45, 0x1e, 0x27, 0x09, 0xe3,
	0x32, 0x6b, 0x87, 0x20, 0x32, 0xfb, 0xb5, 0x80, 0x12, 0x3d, 0xb1, 0x2b, 0x1f, 0x18, 0xbb, 0x83,
	0xa2, 0xa8, 0xa7, 0xf9, 0xe3, 0x1c, 0xef, 0x35, 0xf0, 0x94, 0x7f, 0xe1, 0x17, 0xb0, 0x89, 0x19,
	0x1d, 0x13, 0x2e, 0x42, 0x46, 0xbd, 0x72, 0x40, 0x74, 0xe6, 0x8a, 0xca, 0x7c, 0x3e, 0x4b, 0x66,
	0x69, 0x39, 0x29, 0x24, 0x3a, 0x78, 0x1d, 0x5f, 0x77, 0x69, 0xf7, 0x9b, 0x09, 0x1a, 0xd3, 0x36,
	0x0b, 0xbb, 0x13, 0x5d, 0xd0, 0x43, 0x61, 0x54, 0xb6, 0xf3, 0xea, 0x41, 0xf5, 0x58, 0x10, 0xb0,
	0xa3, 0x8c, 0x97, 0x4e, 0x3b, 0x29, 0x37, 0x2b, 0xe5, 0x5b, 0x99, 0xe4, 0xdf, 0x59, 0x27, 0x62,
	0x7a, 0xc5, 0xc6, 0x2f, 0xc5, 0x20, 0x2c, 0x43, 0x46, 0x8b, 0x09, 0xbc, 0xf9, 0xc9, 0x68, 0x5c,
	0xb5, 0xbb, 0x8a, 0xdc, 0xfd, 0x65, 0x80, 0xf5, 0x6b, 0xdb, 0x0b, 0xdf, 0x83, 0xb5, 0x69, 0x77,
	0x30, 0xd4, 0xad, 0xba, 0xf1, 0xfd, 0xd6, 0x98, 0xbc, 0x37, 0x47, 0x3e, 0xbc, 0x00, 0xf7, 0x31,
	0x67, 0x42, 0xe8, 0x19, 0xf6, 0xae, 0x09, 0x30, 0xab, 0x03, 0xb6, 0x95, 0xa8, 0x18, 0xee, 0xce,
	0x94, 0xb0, 0xf6, 0x1f, 0x03, 0xec, 0x61, 0x16, 0x57, 0x8f, 0x58, 0xbb, 0xae, 0x5f, 0xf2, 0xdd,
	0x2c, 0xa8, 0x6b, 0x9c, 0xbd, 0x29, 0x98, 0x80, 0x45, 0x88, 0x06, 0x36, 0xe3, 0x41, 0x33, 0x20,
	0x54, 0x6d, 0x43, 0x7f, 0x7d, 0x92, 0x50, 0xdc, 0xf0, 0x31, 0x7a, 0x56, 0xae, 0xbe, 0x9a, 0x73,
	0x87, 0xae, 0xfb, 0xdd, 0xdc, 0x39, 0xcc, 0x95, 0xae, 0x2f, 0xec, 0x7c, 0x99, 0xad, 0xfa, 0x2d,
	0xbb, 0xa7, 0x2b, 0x7f, 0xea, 0x9a, 0x81, 0xeb, 0x8b, 0x41, 0x59, 0x33, 0xe8, 0xb7, 0x06, 0x65,
	0xcd, 0x6f, 0x73, 0x2f, 0xbf, 0xe0, 0x38, 0xae, 0x2f, 0x1c, 0xa7, 0xac, 0x72, 0x9c, 0x7e, 0xcb,
	0x71, 0xca, 0xba, 0xe1, 0xbc, 0xda, 0xec, 0xa3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x94,
	0x25, 0x6e, 0x38, 0x07, 0x00, 0x00,
}
