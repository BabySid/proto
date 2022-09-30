// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sodor/thomas.proto

package sodor

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

type RunTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskInstanceId int32 `protobuf:"varint,1,opt,name=task_instance_id,json=taskInstanceId,proto3" json:"task_instance_id,omitempty"`
	JobId          int32 `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId         int32 `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	JobInstanceId  int32 `protobuf:"varint,4,opt,name=job_instance_id,json=jobInstanceId,proto3" json:"job_instance_id,omitempty"`
	Task           *Task `protobuf:"bytes,5,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *RunTaskRequest) Reset() {
	*x = RunTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_thomas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskRequest) ProtoMessage() {}

func (x *RunTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_thomas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskRequest.ProtoReflect.Descriptor instead.
func (*RunTaskRequest) Descriptor() ([]byte, []int) {
	return file_sodor_thomas_proto_rawDescGZIP(), []int{0}
}

func (x *RunTaskRequest) GetTaskInstanceId() int32 {
	if x != nil {
		return x.TaskInstanceId
	}
	return 0
}

func (x *RunTaskRequest) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *RunTaskRequest) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *RunTaskRequest) GetJobInstanceId() int32 {
	if x != nil {
		return x.JobInstanceId
	}
	return 0
}

func (x *RunTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_sodor_thomas_proto protoreflect.FileDescriptor

var file_sodor_thomas_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2f, 0x74, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2f, 0x73, 0x6f, 0x64, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6a,
	0x6f, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x32, 0x60, 0x0a, 0x06, 0x54, 0x68, 0x6f, 0x6d, 0x61,
	0x73, 0x12, 0x28, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x0b,
	0x2e, 0x54, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x54, 0x68,
	0x6f, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x52,
	0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x62, 0x79, 0x53, 0x69, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sodor_thomas_proto_rawDescOnce sync.Once
	file_sodor_thomas_proto_rawDescData = file_sodor_thomas_proto_rawDesc
)

func file_sodor_thomas_proto_rawDescGZIP() []byte {
	file_sodor_thomas_proto_rawDescOnce.Do(func() {
		file_sodor_thomas_proto_rawDescData = protoimpl.X.CompressGZIP(file_sodor_thomas_proto_rawDescData)
	})
	return file_sodor_thomas_proto_rawDescData
}

var file_sodor_thomas_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sodor_thomas_proto_goTypes = []interface{}{
	(*RunTaskRequest)(nil), // 0: RunTaskRequest
	(*Task)(nil),           // 1: Task
	(*ThomasInfo)(nil),     // 2: ThomasInfo
	(*ThomasReply)(nil),    // 3: ThomasReply
	(*EmptyResponse)(nil),  // 4: EmptyResponse
}
var file_sodor_thomas_proto_depIdxs = []int32{
	1, // 0: RunTaskRequest.task:type_name -> Task
	2, // 1: Thomas.HandShake:input_type -> ThomasInfo
	0, // 2: Thomas.RunTask:input_type -> RunTaskRequest
	3, // 3: Thomas.HandShake:output_type -> ThomasReply
	4, // 4: Thomas.RunTask:output_type -> EmptyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sodor_thomas_proto_init() }
func file_sodor_thomas_proto_init() {
	if File_sodor_thomas_proto != nil {
		return
	}
	file_sodor_sodor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sodor_thomas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskRequest); i {
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
			RawDescriptor: file_sodor_thomas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sodor_thomas_proto_goTypes,
		DependencyIndexes: file_sodor_thomas_proto_depIdxs,
		MessageInfos:      file_sodor_thomas_proto_msgTypes,
	}.Build()
	File_sodor_thomas_proto = out.File
	file_sodor_thomas_proto_rawDesc = nil
	file_sodor_thomas_proto_goTypes = nil
	file_sodor_thomas_proto_depIdxs = nil
}
