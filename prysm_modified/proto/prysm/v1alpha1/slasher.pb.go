// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/prysm/v1alpha1/slasher.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	_ "github.com/prysmaticlabs/prysm/v4/proto/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HighestAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorIndex     uint64                                                             `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	HighestSourceEpoch github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch `protobuf:"varint,2,opt,name=highest_source_epoch,json=highestSourceEpoch,proto3" json:"highest_source_epoch,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Epoch"`
	HighestTargetEpoch github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch `protobuf:"varint,3,opt,name=highest_target_epoch,json=highestTargetEpoch,proto3" json:"highest_target_epoch,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Epoch"`
}

func (x *HighestAttestation) Reset() {
	*x = HighestAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v1alpha1_slasher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HighestAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighestAttestation) ProtoMessage() {}

func (x *HighestAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v1alpha1_slasher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighestAttestation.ProtoReflect.Descriptor instead.
func (*HighestAttestation) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v1alpha1_slasher_proto_rawDescGZIP(), []int{0}
}

func (x *HighestAttestation) GetValidatorIndex() uint64 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *HighestAttestation) GetHighestSourceEpoch() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch {
	if x != nil {
		return x.HighestSourceEpoch
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch(0)
}

func (x *HighestAttestation) GetHighestTargetEpoch() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch {
	if x != nil {
		return x.HighestTargetEpoch
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Epoch(0)
}

var File_proto_prysm_v1alpha1_slasher_proto protoreflect.FileDescriptor

var file_proto_prysm_v1alpha1_slasher_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x12, 0x48, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x78, 0x0a, 0x14, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x46, 0x82, 0xb5, 0x18, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x12,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x70, 0x6f,
	0x63, 0x68, 0x12, 0x78, 0x0a, 0x14, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x46, 0x82, 0xb5, 0x18, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x12, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x97, 0x01, 0x0a,
	0x19, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_prysm_v1alpha1_slasher_proto_rawDescOnce sync.Once
	file_proto_prysm_v1alpha1_slasher_proto_rawDescData = file_proto_prysm_v1alpha1_slasher_proto_rawDesc
)

func file_proto_prysm_v1alpha1_slasher_proto_rawDescGZIP() []byte {
	file_proto_prysm_v1alpha1_slasher_proto_rawDescOnce.Do(func() {
		file_proto_prysm_v1alpha1_slasher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prysm_v1alpha1_slasher_proto_rawDescData)
	})
	return file_proto_prysm_v1alpha1_slasher_proto_rawDescData
}

var file_proto_prysm_v1alpha1_slasher_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_prysm_v1alpha1_slasher_proto_goTypes = []interface{}{
	(*HighestAttestation)(nil), // 0: ethereum.eth.v1alpha1.HighestAttestation
}
var file_proto_prysm_v1alpha1_slasher_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_prysm_v1alpha1_slasher_proto_init() }
func file_proto_prysm_v1alpha1_slasher_proto_init() {
	if File_proto_prysm_v1alpha1_slasher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prysm_v1alpha1_slasher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HighestAttestation); i {
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
			RawDescriptor: file_proto_prysm_v1alpha1_slasher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_prysm_v1alpha1_slasher_proto_goTypes,
		DependencyIndexes: file_proto_prysm_v1alpha1_slasher_proto_depIdxs,
		MessageInfos:      file_proto_prysm_v1alpha1_slasher_proto_msgTypes,
	}.Build()
	File_proto_prysm_v1alpha1_slasher_proto = out.File
	file_proto_prysm_v1alpha1_slasher_proto_rawDesc = nil
	file_proto_prysm_v1alpha1_slasher_proto_goTypes = nil
	file_proto_prysm_v1alpha1_slasher_proto_depIdxs = nil
}
