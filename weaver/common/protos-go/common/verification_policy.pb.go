// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: common/verification_policy.proto

package common

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

// VerificationPolicy stores the rules around which parties from a foreign
// network need to provide proof of a view in order for it to be deemed valid by
// the Fabric network
type VerificationPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain string        `protobuf:"bytes,1,opt,name=securityDomain,proto3" json:"securityDomain,omitempty"`
	Identifiers    []*Identifier `protobuf:"bytes,2,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
}

func (x *VerificationPolicy) Reset() {
	*x = VerificationPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_verification_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationPolicy) ProtoMessage() {}

func (x *VerificationPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_common_verification_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationPolicy.ProtoReflect.Descriptor instead.
func (*VerificationPolicy) Descriptor() ([]byte, []int) {
	return file_common_verification_policy_proto_rawDescGZIP(), []int{0}
}

func (x *VerificationPolicy) GetSecurityDomain() string {
	if x != nil {
		return x.SecurityDomain
	}
	return ""
}

func (x *VerificationPolicy) GetIdentifiers() []*Identifier {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

// The Policy captures the list of parties that are required to provide proofs
// of a view in order for the Fabric network to accept the view as valid.
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Criteria []string `protobuf:"bytes,2,rep,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_verification_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_common_verification_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_common_verification_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Policy) GetCriteria() []string {
	if x != nil {
		return x.Criteria
	}
	return nil
}

// List of identifiers for the VerificationPolicy
type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pattern defines the view/views that this rule applies to
	// A rule may contain a "*" at the end of the pattern
	Pattern string  `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Policy  *Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_verification_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_common_verification_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_common_verification_policy_proto_rawDescGZIP(), []int{2}
}

func (x *Identifier) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *Identifier) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_common_verification_policy_proto protoreflect.FileDescriptor

var file_common_verification_policy_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x86,
	0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x48, 0x0a,
	0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x38, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x22, 0x62, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x80, 0x01, 0x0a, 0x3e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79,
	0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_verification_policy_proto_rawDescOnce sync.Once
	file_common_verification_policy_proto_rawDescData = file_common_verification_policy_proto_rawDesc
)

func file_common_verification_policy_proto_rawDescGZIP() []byte {
	file_common_verification_policy_proto_rawDescOnce.Do(func() {
		file_common_verification_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_verification_policy_proto_rawDescData)
	})
	return file_common_verification_policy_proto_rawDescData
}

var file_common_verification_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_verification_policy_proto_goTypes = []interface{}{
	(*VerificationPolicy)(nil), // 0: common.verification_policy.VerificationPolicy
	(*Policy)(nil),             // 1: common.verification_policy.Policy
	(*Identifier)(nil),         // 2: common.verification_policy.Identifier
}
var file_common_verification_policy_proto_depIdxs = []int32{
	2, // 0: common.verification_policy.VerificationPolicy.identifiers:type_name -> common.verification_policy.Identifier
	1, // 1: common.verification_policy.Identifier.policy:type_name -> common.verification_policy.Policy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_verification_policy_proto_init() }
func file_common_verification_policy_proto_init() {
	if File_common_verification_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_verification_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationPolicy); i {
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
		file_common_verification_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_common_verification_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
			RawDescriptor: file_common_verification_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_verification_policy_proto_goTypes,
		DependencyIndexes: file_common_verification_policy_proto_depIdxs,
		MessageInfos:      file_common_verification_policy_proto_msgTypes,
	}.Build()
	File_common_verification_policy_proto = out.File
	file_common_verification_policy_proto_rawDesc = nil
	file_common_verification_policy_proto_goTypes = nil
	file_common_verification_policy_proto_depIdxs = nil
}
