// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: sodor/sodor.proto

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

type TaskState int32

const (
	TaskState_TS_INIT    TaskState = 0
	TaskState_TS_RUNNING TaskState = 1
	TaskState_TS_STOPPED TaskState = 2
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0: "TS_INIT",
		1: "TS_RUNNING",
		2: "TS_STOPPED",
	}
	TaskState_value = map[string]int32{
		"TS_INIT":    0,
		"TS_RUNNING": 1,
		"TS_STOPPED": 2,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_sodor_sodor_proto_enumTypes[0].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_sodor_sodor_proto_enumTypes[0]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{0}
}

type SchedulerMode int32

const (
	SchedulerMode_SM_Null    SchedulerMode = 0
	SchedulerMode_SM_None    SchedulerMode = 1
	SchedulerMode_SM_Crontab SchedulerMode = 2
)

// Enum value maps for SchedulerMode.
var (
	SchedulerMode_name = map[int32]string{
		0: "SM_Null",
		1: "SM_None",
		2: "SM_Crontab",
	}
	SchedulerMode_value = map[string]int32{
		"SM_Null":    0,
		"SM_None":    1,
		"SM_Crontab": 2,
	}
)

func (x SchedulerMode) Enum() *SchedulerMode {
	p := new(SchedulerMode)
	*p = x
	return p
}

func (x SchedulerMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchedulerMode) Descriptor() protoreflect.EnumDescriptor {
	return file_sodor_sodor_proto_enumTypes[1].Descriptor()
}

func (SchedulerMode) Type() protoreflect.EnumType {
	return &file_sodor_sodor_proto_enumTypes[1]
}

func (x SchedulerMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchedulerMode.Descriptor instead.
func (SchedulerMode) EnumDescriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{1}
}

type HostType int32

const (
	HostType_HT_Null HostType = 0
	HostType_HT_IP   HostType = 1
	HostType_HT_Tag  HostType = 2
)

// Enum value maps for HostType.
var (
	HostType_name = map[int32]string{
		0: "HT_Null",
		1: "HT_IP",
		2: "HT_Tag",
	}
	HostType_value = map[string]int32{
		"HT_Null": 0,
		"HT_IP":   1,
		"HT_Tag":  2,
	}
)

func (x HostType) Enum() *HostType {
	p := new(HostType)
	*p = x
	return p
}

func (x HostType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostType) Descriptor() protoreflect.EnumDescriptor {
	return file_sodor_sodor_proto_enumTypes[2].Descriptor()
}

func (HostType) Type() protoreflect.EnumType {
	return &file_sodor_sodor_proto_enumTypes[2]
}

func (x HostType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostType.Descriptor instead.
func (HostType) EnumDescriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{2}
}

type JobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobReply) Reset() {
	*x = JobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReply) ProtoMessage() {}

func (x *JobReply) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReply.ProtoReflect.Descriptor instead.
func (*JobReply) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{0}
}

func (x *JobReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RoutineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CtSpec string `protobuf:"bytes,1,opt,name=ct_spec,json=ctSpec,proto3" json:"ct_spec,omitempty"`
}

func (x *RoutineSpec) Reset() {
	*x = RoutineSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutineSpec) ProtoMessage() {}

func (x *RoutineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutineSpec.ProtoReflect.Descriptor instead.
func (*RoutineSpec) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{1}
}

func (x *RoutineSpec) GetCtSpec() string {
	if x != nil {
		return x.CtSpec
	}
	return ""
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  HostType `protobuf:"varint,1,opt,name=type,proto3,enum=HostType" json:"type,omitempty"`
	Nodes []string `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{2}
}

func (x *Host) GetType() HostType {
	if x != nil {
		return x.Type
	}
	return HostType_HT_Null
}

func (x *Host) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt       int32         `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt       int32         `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	JobId          int32         `protobuf:"varint,4,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Name           string        `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	RunningHosts   []*Host       `protobuf:"bytes,6,rep,name=running_hosts,json=runningHosts,proto3" json:"running_hosts,omitempty"`
	SchedulerMode  SchedulerMode `protobuf:"varint,7,opt,name=scheduler_mode,json=schedulerMode,proto3,enum=SchedulerMode" json:"scheduler_mode,omitempty"`
	RoutineSpec    *RoutineSpec  `protobuf:"bytes,8,opt,name=routine_spec,json=routineSpec,proto3" json:"routine_spec,omitempty"`
	Script         string        `protobuf:"bytes,9,opt,name=script,proto3" json:"script,omitempty"`
	RunningTimeout int32         `protobuf:"varint,10,opt,name=running_timeout,json=runningTimeout,proto3" json:"running_timeout,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Task) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Task) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetRunningHosts() []*Host {
	if x != nil {
		return x.RunningHosts
	}
	return nil
}

func (x *Task) GetSchedulerMode() SchedulerMode {
	if x != nil {
		return x.SchedulerMode
	}
	return SchedulerMode_SM_Null
}

func (x *Task) GetRoutineSpec() *RoutineSpec {
	if x != nil {
		return x.RoutineSpec
	}
	return nil
}

func (x *Task) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

func (x *Task) GetRunningTimeout() int32 {
	if x != nil {
		return x.RunningTimeout
	}
	return 0
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt  int32           `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt  int32           `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Name      string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Tasks     []*Task         `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Relations []*TaskRelation `protobuf:"bytes,6,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{4}
}

func (x *Job) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Job) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Job) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *Job) GetRelations() []*TaskRelation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type Jobs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *Jobs) Reset() {
	*x = Jobs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jobs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jobs) ProtoMessage() {}

func (x *Jobs) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jobs.ProtoReflect.Descriptor instead.
func (*Jobs) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{5}
}

func (x *Jobs) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type TaskRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromTask string `protobuf:"bytes,1,opt,name=from_task,json=fromTask,proto3" json:"from_task,omitempty"`
	ToTask   string `protobuf:"bytes,2,opt,name=to_task,json=toTask,proto3" json:"to_task,omitempty"`
}

func (x *TaskRelation) Reset() {
	*x = TaskRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRelation) ProtoMessage() {}

func (x *TaskRelation) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRelation.ProtoReflect.Descriptor instead.
func (*TaskRelation) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{6}
}

func (x *TaskRelation) GetFromTask() string {
	if x != nil {
		return x.FromTask
	}
	return ""
}

func (x *TaskRelation) GetToTask() string {
	if x != nil {
		return x.ToTask
	}
	return ""
}

type AlertGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt  int32   `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt  int32   `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Name      string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PluginIds []int64 `protobuf:"varint,5,rep,packed,name=plugin_ids,json=pluginIds,proto3" json:"plugin_ids,omitempty"`
}

func (x *AlertGroup) Reset() {
	*x = AlertGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertGroup) ProtoMessage() {}

func (x *AlertGroup) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertGroup.ProtoReflect.Descriptor instead.
func (*AlertGroup) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{7}
}

func (x *AlertGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlertGroup) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *AlertGroup) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *AlertGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertGroup) GetPluginIds() []int64 {
	if x != nil {
		return x.PluginIds
	}
	return nil
}

type AlertPluginParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlertPluginParam) Reset() {
	*x = AlertPluginParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPluginParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPluginParam) ProtoMessage() {}

func (x *AlertPluginParam) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPluginParam.ProtoReflect.Descriptor instead.
func (*AlertPluginParam) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{8}
}

type AlertPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt int32             `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt int32             `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Catalog  string            `protobuf:"bytes,4,opt,name=catalog,proto3" json:"catalog,omitempty"`
	Params   *AlertPluginParam `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *AlertPlugin) Reset() {
	*x = AlertPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPlugin) ProtoMessage() {}

func (x *AlertPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPlugin.ProtoReflect.Descriptor instead.
func (*AlertPlugin) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{9}
}

func (x *AlertPlugin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlertPlugin) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *AlertPlugin) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *AlertPlugin) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *AlertPlugin) GetParams() *AlertPluginParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type AlertHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt int32 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt int32 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *AlertHistory) Reset() {
	*x = AlertHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sodor_sodor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertHistory) ProtoMessage() {}

func (x *AlertHistory) ProtoReflect() protoreflect.Message {
	mi := &file_sodor_sodor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertHistory.ProtoReflect.Descriptor instead.
func (*AlertHistory) Descriptor() ([]byte, []int) {
	return file_sodor_sodor_proto_rawDescGZIP(), []int{10}
}

func (x *AlertHistory) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlertHistory) GetCreateAt() int32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *AlertHistory) GetUpdateAt() int32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

var File_sodor_sodor_proto protoreflect.FileDescriptor

var file_sodor_sodor_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2f, 0x73, 0x6f, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x3b, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0xd0, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x35,
	0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x18, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x44, 0x0a, 0x0c, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x22,
	0x89, 0x01, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22,
	0x9c, 0x01, 0x0a, 0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x58,
	0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x2a, 0x38, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x53, 0x5f, 0x49, 0x4e, 0x49, 0x54,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x02, 0x2a, 0x39, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4d, 0x5f, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4d, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x4d, 0x5f, 0x43, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x10, 0x02, 0x2a, 0x2e, 0x0a,
	0x08, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x54, 0x5f,
	0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x5f, 0x49, 0x50, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x54, 0x5f, 0x54, 0x61, 0x67, 0x10, 0x02, 0x32, 0x2f, 0x0a,
	0x0d, 0x46, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x04, 0x2e, 0x4a, 0x6f,
	0x62, 0x1a, 0x09, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x62,
	0x79, 0x53, 0x69, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x64, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sodor_sodor_proto_rawDescOnce sync.Once
	file_sodor_sodor_proto_rawDescData = file_sodor_sodor_proto_rawDesc
)

func file_sodor_sodor_proto_rawDescGZIP() []byte {
	file_sodor_sodor_proto_rawDescOnce.Do(func() {
		file_sodor_sodor_proto_rawDescData = protoimpl.X.CompressGZIP(file_sodor_sodor_proto_rawDescData)
	})
	return file_sodor_sodor_proto_rawDescData
}

var file_sodor_sodor_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sodor_sodor_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sodor_sodor_proto_goTypes = []interface{}{
	(TaskState)(0),           // 0: TaskState
	(SchedulerMode)(0),       // 1: SchedulerMode
	(HostType)(0),            // 2: HostType
	(*JobReply)(nil),         // 3: JobReply
	(*RoutineSpec)(nil),      // 4: RoutineSpec
	(*Host)(nil),             // 5: Host
	(*Task)(nil),             // 6: Task
	(*Job)(nil),              // 7: Job
	(*Jobs)(nil),             // 8: Jobs
	(*TaskRelation)(nil),     // 9: TaskRelation
	(*AlertGroup)(nil),       // 10: AlertGroup
	(*AlertPluginParam)(nil), // 11: AlertPluginParam
	(*AlertPlugin)(nil),      // 12: AlertPlugin
	(*AlertHistory)(nil),     // 13: AlertHistory
}
var file_sodor_sodor_proto_depIdxs = []int32{
	2,  // 0: Host.type:type_name -> HostType
	5,  // 1: Task.running_hosts:type_name -> Host
	1,  // 2: Task.scheduler_mode:type_name -> SchedulerMode
	4,  // 3: Task.routine_spec:type_name -> RoutineSpec
	6,  // 4: Job.tasks:type_name -> Task
	9,  // 5: Job.relations:type_name -> TaskRelation
	7,  // 6: Jobs.jobs:type_name -> Job
	11, // 7: AlertPlugin.params:type_name -> AlertPluginParam
	7,  // 8: FatController.CreateJob:input_type -> Job
	3,  // 9: FatController.CreateJob:output_type -> JobReply
	9,  // [9:10] is the sub-list for method output_type
	8,  // [8:9] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_sodor_sodor_proto_init() }
func file_sodor_sodor_proto_init() {
	if File_sodor_sodor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sodor_sodor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReply); i {
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
		file_sodor_sodor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutineSpec); i {
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
		file_sodor_sodor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_sodor_sodor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_sodor_sodor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_sodor_sodor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jobs); i {
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
		file_sodor_sodor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRelation); i {
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
		file_sodor_sodor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertGroup); i {
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
		file_sodor_sodor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPluginParam); i {
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
		file_sodor_sodor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPlugin); i {
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
		file_sodor_sodor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertHistory); i {
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
			RawDescriptor: file_sodor_sodor_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sodor_sodor_proto_goTypes,
		DependencyIndexes: file_sodor_sodor_proto_depIdxs,
		EnumInfos:         file_sodor_sodor_proto_enumTypes,
		MessageInfos:      file_sodor_sodor_proto_msgTypes,
	}.Build()
	File_sodor_sodor_proto = out.File
	file_sodor_sodor_proto_rawDesc = nil
	file_sodor_sodor_proto_goTypes = nil
	file_sodor_sodor_proto_depIdxs = nil
}
