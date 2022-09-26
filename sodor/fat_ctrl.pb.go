// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sodor/fat_ctrl.proto

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

// thomas
type ThomasHandShakeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version   string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Proto     string `protobuf:"bytes,4,opt,name=proto,proto3" json:"proto,omitempty"`
	Host      string `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Port      int32  `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	Pid       int32  `protobuf:"varint,7,opt,name=pid,proto3" json:"pid,omitempty"`
	StartTime int32  `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *ThomasHandShakeReq) Reset() {
	*x = ThomasHandShakeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_fat_ctrl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThomasHandShakeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThomasHandShakeReq) ProtoMessage() {}

func (x *ThomasHandShakeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_fat_ctrl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThomasHandShakeReq.ProtoReflect.Descriptor instead.
func (*ThomasHandShakeReq) Descriptor() ([]byte, []int) {
	return file_sodor_fat_ctrl_proto_rawDescGZIP(), []int{0}
}

func (x *ThomasHandShakeReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ThomasHandShakeReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ThomasHandShakeReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ThomasHandShakeReq) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *ThomasHandShakeReq) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ThomasHandShakeReq) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ThomasHandShakeReq) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ThomasHandShakeReq) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

type ThomasHandShakeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ThomasHandShakeResp) Reset() {
	*x = ThomasHandShakeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_fat_ctrl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThomasHandShakeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThomasHandShakeResp) ProtoMessage() {}

func (x *ThomasHandShakeResp) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_fat_ctrl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThomasHandShakeResp.ProtoReflect.Descriptor instead.
func (*ThomasHandShakeResp) Descriptor() ([]byte, []int) {
	return file_sodor_fat_ctrl_proto_rawDescGZIP(), []int{1}
}

func (x *ThomasHandShakeResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId    int32  `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	JobId         int32  `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId        int32  `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	JobInstanceId int32  `protobuf:"varint,4,opt,name=job_instance_id,json=jobInstanceId,proto3" json:"job_instance_id,omitempty"`
	StartTs       int32  `protobuf:"varint,5,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	StopTs        int32  `protobuf:"varint,6,opt,name=stop_ts,json=stopTs,proto3" json:"stop_ts,omitempty"`
	Pid           int32  `protobuf:"varint,7,opt,name=pid,proto3" json:"pid,omitempty"`
	ExitCode      int32  `protobuf:"varint,8,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	ExitMsg       string `protobuf:"bytes,9,opt,name=exit_msg,json=exitMsg,proto3" json:"exit_msg,omitempty"`
}

func (x *TaskInstance) Reset() {
	*x = TaskInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_fat_ctrl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInstance) ProtoMessage() {}

func (x *TaskInstance) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_fat_ctrl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInstance.ProtoReflect.Descriptor instead.
func (*TaskInstance) Descriptor() ([]byte, []int) {
	return file_sodor_fat_ctrl_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInstance) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *TaskInstance) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *TaskInstance) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskInstance) GetJobInstanceId() int32 {
	if x != nil {
		return x.JobInstanceId
	}
	return 0
}

func (x *TaskInstance) GetStartTs() int32 {
	if x != nil {
		return x.StartTs
	}
	return 0
}

func (x *TaskInstance) GetStopTs() int32 {
	if x != nil {
		return x.StopTs
	}
	return 0
}

func (x *TaskInstance) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *TaskInstance) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *TaskInstance) GetExitMsg() string {
	if x != nil {
		return x.ExitMsg
	}
	return ""
}

var File_sodor_fat_ctrl_proto protoreflect.FileDescriptor

var file_sodor_fat_ctrl_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2f, 0x66, 0x61, 0x74, 0x5f, 0x63, 0x74, 0x72, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2f, 0x73, 0x6f,
	0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x12, 0x54, 0x68,
	0x6f, 0x6d, 0x61, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x25, 0x0a,
	0x13, 0x54, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x70, 0x5f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x70,
	0x54, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x69, 0x74, 0x4d, 0x73, 0x67, 0x32, 0x88, 0x01, 0x0a,
	0x0d, 0x46, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x13, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54,
	0x68, 0x6f, 0x6d, 0x61, 0x73, 0x12, 0x13, 0x2e, 0x54, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x54, 0x68, 0x6f,
	0x6d, 0x61, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x62, 0x79, 0x53, 0x69, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sodor_fat_ctrl_proto_rawDescOnce sync.Once
	file_sodor_fat_ctrl_proto_rawDescData = file_sodor_fat_ctrl_proto_rawDesc
)

func file_sodor_fat_ctrl_proto_rawDescGZIP() []byte {
	file_sodor_fat_ctrl_proto_rawDescOnce.Do(func() {
		file_sodor_fat_ctrl_proto_rawDescData = protoimpl.X.CompressGZIP(file_sodor_fat_ctrl_proto_rawDescData)
	})
	return file_sodor_fat_ctrl_proto_rawDescData
}

var file_sodor_fat_ctrl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sodor_fat_ctrl_proto_goTypes = []interface{}{
	(*ThomasHandShakeReq)(nil),  // 0: ThomasHandShakeReq
	(*ThomasHandShakeResp)(nil), // 1: ThomasHandShakeResp
	(*TaskInstance)(nil),        // 2: TaskInstance
	(*EmptyResponse)(nil),       // 3: EmptyResponse
}
var file_sodor_fat_ctrl_proto_depIdxs = []int32{
	0, // 0: FatController.HandShakeWithThomas:input_type -> ThomasHandShakeReq
	2, // 1: FatController.UpdateTaskStatus:input_type -> TaskInstance
	1, // 2: FatController.HandShakeWithThomas:output_type -> ThomasHandShakeResp
	3, // 3: FatController.UpdateTaskStatus:output_type -> EmptyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sodor_fat_ctrl_proto_init() }
func file_sodor_fat_ctrl_proto_init() {
	if File_sodor_fat_ctrl_proto != nil {
		return
	}
	file_sodor_sodor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sodor_fat_ctrl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThomasHandShakeReq); i {
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
		file_sodor_fat_ctrl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThomasHandShakeResp); i {
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
		file_sodor_fat_ctrl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInstance); i {
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
			RawDescriptor: file_sodor_fat_ctrl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sodor_fat_ctrl_proto_goTypes,
		DependencyIndexes: file_sodor_fat_ctrl_proto_depIdxs,
		MessageInfos:      file_sodor_fat_ctrl_proto_msgTypes,
	}.Build()
	File_sodor_fat_ctrl_proto = out.File
	file_sodor_fat_ctrl_proto_rawDesc = nil
	file_sodor_fat_ctrl_proto_goTypes = nil
	file_sodor_fat_ctrl_proto_depIdxs = nil
}
