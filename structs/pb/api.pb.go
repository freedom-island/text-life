// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_Done            Status = 0
	Status_ClientParamErr  Status = 1
	Status_ServerException Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Done",
		1: "ClientParamErr",
		2: "ServerException",
	}
	Status_value = map[string]int32{
		"Done":            0,
		"ClientParamErr":  1,
		"ServerException": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type RouteId int32

const (
	RouteId_BodyInfo RouteId = 0
	RouteId_Eat      RouteId = 1
	RouteId_Farm     RouteId = 2
)

// Enum value maps for RouteId.
var (
	RouteId_name = map[int32]string{
		0: "BodyInfo",
		1: "Eat",
		2: "Farm",
	}
	RouteId_value = map[string]int32{
		"BodyInfo": 0,
		"Eat":      1,
		"Farm":     2,
	}
)

func (x RouteId) Enum() *RouteId {
	p := new(RouteId)
	*p = x
	return p
}

func (x RouteId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteId) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (RouteId) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x RouteId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteId.Descriptor instead.
func (RouteId) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type MsgPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId   uint64                 `protobuf:"varint,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`                              // 请求ID
	RouteId RouteId                `protobuf:"varint,2,opt,name=route_id,json=routeId,proto3,enum=text_life.RouteId" json:"route_id,omitempty"` // 路由标记ID
	ReqTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=req_time,json=reqTime,proto3" json:"req_time,omitempty"`                         // 请求时间
	// Types that are assignable to Body:
	//
	//	*MsgPacket_BodyInfoReq
	//	*MsgPacket_BodyInfoResp
	//	*MsgPacket_EatReq
	//	*MsgPacket_EatResp
	//	*MsgPacket_FarmReq
	//	*MsgPacket_FarmResp
	Body isMsgPacket_Body `protobuf_oneof:"body"`
}

func (x *MsgPacket) Reset() {
	*x = MsgPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPacket) ProtoMessage() {}

func (x *MsgPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPacket.ProtoReflect.Descriptor instead.
func (*MsgPacket) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *MsgPacket) GetReqId() uint64 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *MsgPacket) GetRouteId() RouteId {
	if x != nil {
		return x.RouteId
	}
	return RouteId_BodyInfo
}

func (x *MsgPacket) GetReqTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReqTime
	}
	return nil
}

func (m *MsgPacket) GetBody() isMsgPacket_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *MsgPacket) GetBodyInfoReq() *BodyInfoReq {
	if x, ok := x.GetBody().(*MsgPacket_BodyInfoReq); ok {
		return x.BodyInfoReq
	}
	return nil
}

func (x *MsgPacket) GetBodyInfoResp() *BodyInfoResp {
	if x, ok := x.GetBody().(*MsgPacket_BodyInfoResp); ok {
		return x.BodyInfoResp
	}
	return nil
}

func (x *MsgPacket) GetEatReq() *EatReq {
	if x, ok := x.GetBody().(*MsgPacket_EatReq); ok {
		return x.EatReq
	}
	return nil
}

func (x *MsgPacket) GetEatResp() *EatResp {
	if x, ok := x.GetBody().(*MsgPacket_EatResp); ok {
		return x.EatResp
	}
	return nil
}

func (x *MsgPacket) GetFarmReq() *FarmReq {
	if x, ok := x.GetBody().(*MsgPacket_FarmReq); ok {
		return x.FarmReq
	}
	return nil
}

func (x *MsgPacket) GetFarmResp() *FarmResp {
	if x, ok := x.GetBody().(*MsgPacket_FarmResp); ok {
		return x.FarmResp
	}
	return nil
}

type isMsgPacket_Body interface {
	isMsgPacket_Body()
}

type MsgPacket_BodyInfoReq struct {
	BodyInfoReq *BodyInfoReq `protobuf:"bytes,13,opt,name=body_info_req,json=bodyInfoReq,proto3,oneof"` // 身体状况
}

type MsgPacket_BodyInfoResp struct {
	BodyInfoResp *BodyInfoResp `protobuf:"bytes,14,opt,name=body_info_resp,json=bodyInfoResp,proto3,oneof"`
}

type MsgPacket_EatReq struct {
	EatReq *EatReq `protobuf:"bytes,15,opt,name=eat_req,json=eatReq,proto3,oneof"`
}

type MsgPacket_EatResp struct {
	EatResp *EatResp `protobuf:"bytes,16,opt,name=eat_resp,json=eatResp,proto3,oneof"`
}

type MsgPacket_FarmReq struct {
	FarmReq *FarmReq `protobuf:"bytes,17,opt,name=farm_req,json=farmReq,proto3,oneof"`
}

type MsgPacket_FarmResp struct {
	FarmResp *FarmResp `protobuf:"bytes,18,opt,name=farm_resp,json=farmResp,proto3,oneof"`
}

func (*MsgPacket_BodyInfoReq) isMsgPacket_Body() {}

func (*MsgPacket_BodyInfoResp) isMsgPacket_Body() {}

func (*MsgPacket_EatReq) isMsgPacket_Body() {}

func (*MsgPacket_EatResp) isMsgPacket_Body() {}

func (*MsgPacket_FarmReq) isMsgPacket_Body() {}

func (*MsgPacket_FarmResp) isMsgPacket_Body() {}

type BodyInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId uint64 `protobuf:"varint,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
}

func (x *BodyInfoReq) Reset() {
	*x = BodyInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInfoReq) ProtoMessage() {}

func (x *BodyInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInfoReq.ProtoReflect.Descriptor instead.
func (*BodyInfoReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *BodyInfoReq) GetTargetId() uint64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type EatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BagId   uint64 `protobuf:"varint,1,opt,name=bagId,proto3" json:"bagId,omitempty"`
	GoodsId uint64 `protobuf:"varint,2,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
}

func (x *EatReq) Reset() {
	*x = EatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatReq) ProtoMessage() {}

func (x *EatReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatReq.ProtoReflect.Descriptor instead.
func (*EatReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *EatReq) GetBagId() uint64 {
	if x != nil {
		return x.BagId
	}
	return 0
}

func (x *EatReq) GetGoodsId() uint64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

type BodyInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Satiation uint32 `protobuf:"varint,2,opt,name=satiation,proto3" json:"satiation,omitempty"` // 饱腹
	Energy    uint32 `protobuf:"varint,3,opt,name=energy,proto3" json:"energy,omitempty"`       // 精力
	Knowledge uint32 `protobuf:"varint,4,opt,name=knowledge,proto3" json:"knowledge,omitempty"` // 知识水准
	Age       uint32 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`             // 年龄
}

func (x *BodyInfoResp) Reset() {
	*x = BodyInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInfoResp) ProtoMessage() {}

func (x *BodyInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInfoResp.ProtoReflect.Descriptor instead.
func (*BodyInfoResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *BodyInfoResp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BodyInfoResp) GetSatiation() uint32 {
	if x != nil {
		return x.Satiation
	}
	return 0
}

func (x *BodyInfoResp) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *BodyInfoResp) GetKnowledge() uint32 {
	if x != nil {
		return x.Knowledge
	}
	return 0
}

func (x *BodyInfoResp) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type EatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=text_life.Status" json:"status,omitempty"`
}

func (x *EatResp) Reset() {
	*x = EatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatResp) ProtoMessage() {}

func (x *EatResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatResp.ProtoReflect.Descriptor instead.
func (*EatResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *EatResp) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Done
}

type FarmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=text_life.Status" json:"status,omitempty"`
}

func (x *FarmReq) Reset() {
	*x = FarmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FarmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FarmReq) ProtoMessage() {}

func (x *FarmReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FarmReq.ProtoReflect.Descriptor instead.
func (*FarmReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *FarmReq) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Done
}

type FarmResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=text_life.Status" json:"status,omitempty"`
}

func (x *FarmResp) Reset() {
	*x = FarmResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FarmResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FarmResp) ProtoMessage() {}

func (x *FarmResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FarmResp.ProtoReflect.Descriptor instead.
func (*FarmResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *FarmResp) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Done
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x08,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x64, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x72, 0x65, 0x71, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x3f, 0x0a, 0x0e, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x6c, 0x69, 0x66, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x45,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x06, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x2f, 0x0a, 0x08, 0x65, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x45, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x07, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x2f, 0x0a, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x46,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x07, 0x66, 0x61, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x32, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65,
	0x2e, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x08, 0x66, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x29, 0x0a,
	0x0b, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x06, 0x45, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x62, 0x61, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x61, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x73, 0x61, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x22, 0x34, 0x0a, 0x07, 0x45, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x34, 0x0a, 0x07, 0x46, 0x61, 0x72, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35,
	0x0a, 0x08, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x3b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x02, 0x2a, 0x2a, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x0c, 0x0a,
	0x08, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45,
	0x61, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x10, 0x02, 0x42, 0x0c,
	0x5a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: text_life.Status
	(RouteId)(0),                  // 1: text_life.RouteId
	(*MsgPacket)(nil),             // 2: text_life.MsgPacket
	(*BodyInfoReq)(nil),           // 3: text_life.BodyInfoReq
	(*EatReq)(nil),                // 4: text_life.EatReq
	(*BodyInfoResp)(nil),          // 5: text_life.BodyInfoResp
	(*EatResp)(nil),               // 6: text_life.EatResp
	(*FarmReq)(nil),               // 7: text_life.FarmReq
	(*FarmResp)(nil),              // 8: text_life.FarmResp
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_api_proto_depIdxs = []int32{
	1,  // 0: text_life.MsgPacket.route_id:type_name -> text_life.RouteId
	9,  // 1: text_life.MsgPacket.req_time:type_name -> google.protobuf.Timestamp
	3,  // 2: text_life.MsgPacket.body_info_req:type_name -> text_life.BodyInfoReq
	5,  // 3: text_life.MsgPacket.body_info_resp:type_name -> text_life.BodyInfoResp
	4,  // 4: text_life.MsgPacket.eat_req:type_name -> text_life.EatReq
	6,  // 5: text_life.MsgPacket.eat_resp:type_name -> text_life.EatResp
	7,  // 6: text_life.MsgPacket.farm_req:type_name -> text_life.FarmReq
	8,  // 7: text_life.MsgPacket.farm_resp:type_name -> text_life.FarmResp
	0,  // 8: text_life.EatResp.status:type_name -> text_life.Status
	0,  // 9: text_life.FarmReq.status:type_name -> text_life.Status
	0,  // 10: text_life.FarmResp.status:type_name -> text_life.Status
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPacket); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInfoReq); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatReq); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInfoResp); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatResp); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FarmReq); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FarmResp); i {
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
	file_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MsgPacket_BodyInfoReq)(nil),
		(*MsgPacket_BodyInfoResp)(nil),
		(*MsgPacket_EatReq)(nil),
		(*MsgPacket_EatResp)(nil),
		(*MsgPacket_FarmReq)(nil),
		(*MsgPacket_FarmResp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
