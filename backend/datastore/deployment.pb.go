// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: deployment.proto

package datastore

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Deployment_DeployStatus int32

const (
	Deployment_NOT_STARTED Deployment_DeployStatus = 0
	Deployment_IN_PROGRESS Deployment_DeployStatus = 1
	Deployment_COMPLETE    Deployment_DeployStatus = 2
	Deployment_ERROR       Deployment_DeployStatus = 3
)

// Enum value maps for Deployment_DeployStatus.
var (
	Deployment_DeployStatus_name = map[int32]string{
		0: "NOT_STARTED",
		1: "IN_PROGRESS",
		2: "COMPLETE",
		3: "ERROR",
	}
	Deployment_DeployStatus_value = map[string]int32{
		"NOT_STARTED": 0,
		"IN_PROGRESS": 1,
		"COMPLETE":    2,
		"ERROR":       3,
	}
)

func (x Deployment_DeployStatus) Enum() *Deployment_DeployStatus {
	p := new(Deployment_DeployStatus)
	*p = x
	return p
}

func (x Deployment_DeployStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Deployment_DeployStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_deployment_proto_enumTypes[0].Descriptor()
}

func (Deployment_DeployStatus) Type() protoreflect.EnumType {
	return &file_deployment_proto_enumTypes[0]
}

func (x Deployment_DeployStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Deployment_DeployStatus.Descriptor instead.
func (Deployment_DeployStatus) EnumDescriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0, 0}
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repository      string                  `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Dockerfile      string                  `protobuf:"bytes,3,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
	Commit          string                  `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	LastDeploy      *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=last_deploy,json=lastDeploy,proto3" json:"last_deploy,omitempty"`
	Status          Deployment_DeployStatus `protobuf:"varint,6,opt,name=status,proto3,enum=datastore.Deployment_DeployStatus" json:"status,omitempty"`
	Domain          string                  `protobuf:"bytes,7,opt,name=domain,proto3" json:"domain,omitempty"`
	BuildCommand    string                  `protobuf:"bytes,8,opt,name=build_command,json=buildCommand,proto3" json:"build_command,omitempty"`
	OutputDirectory string                  `protobuf:"bytes,9,opt,name=output_directory,json=outputDirectory,proto3" json:"output_directory,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Deployment) GetDockerfile() string {
	if x != nil {
		return x.Dockerfile
	}
	return ""
}

func (x *Deployment) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *Deployment) GetLastDeploy() *timestamp.Timestamp {
	if x != nil {
		return x.LastDeploy
	}
	return nil
}

func (x *Deployment) GetStatus() Deployment_DeployStatus {
	if x != nil {
		return x.Status
	}
	return Deployment_NOT_STARTED
}

func (x *Deployment) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Deployment) GetBuildCommand() string {
	if x != nil {
		return x.BuildCommand
	}
	return ""
}

func (x *Deployment) GetOutputDirectory() string {
	if x != nil {
		return x.OutputDirectory
	}
	return ""
}

var File_deployment_proto protoreflect.FileDescriptor

var file_deployment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x03, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x49, 0x0a, 0x0c, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_deployment_proto_rawDescOnce sync.Once
	file_deployment_proto_rawDescData = file_deployment_proto_rawDesc
)

func file_deployment_proto_rawDescGZIP() []byte {
	file_deployment_proto_rawDescOnce.Do(func() {
		file_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_proto_rawDescData)
	})
	return file_deployment_proto_rawDescData
}

var file_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_deployment_proto_goTypes = []interface{}{
	(Deployment_DeployStatus)(0), // 0: datastore.Deployment.DeployStatus
	(*Deployment)(nil),           // 1: datastore.Deployment
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_deployment_proto_depIdxs = []int32{
	2, // 0: datastore.Deployment.last_deploy:type_name -> google.protobuf.Timestamp
	0, // 1: datastore.Deployment.status:type_name -> datastore.Deployment.DeployStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_deployment_proto_init() }
func file_deployment_proto_init() {
	if File_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
			RawDescriptor: file_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deployment_proto_goTypes,
		DependencyIndexes: file_deployment_proto_depIdxs,
		EnumInfos:         file_deployment_proto_enumTypes,
		MessageInfos:      file_deployment_proto_msgTypes,
	}.Build()
	File_deployment_proto = out.File
	file_deployment_proto_rawDesc = nil
	file_deployment_proto_goTypes = nil
	file_deployment_proto_depIdxs = nil
}