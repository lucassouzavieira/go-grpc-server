// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/protobuf/schema/fleet.proto

package fleet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FleetNumber       string `protobuf:"bytes,1,opt,name=fleet_number,json=fleetNumber,proto3" json:"fleet_number,omitempty"`
	OperationalStatus string `protobuf:"bytes,2,opt,name=operational_status,json=operationalStatus,proto3" json:"operational_status,omitempty"`
	Lfb               string `protobuf:"bytes,3,opt,name=lfb,proto3" json:"lfb,omitempty"`
	Make              string `protobuf:"bytes,4,opt,name=make,proto3" json:"make,omitempty"`
	Model             string `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	Type              string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Category          string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	RegistrationYear  int32  `protobuf:"varint,8,opt,name=registration_year,json=registrationYear,proto3" json:"registration_year,omitempty"`
	Life              int32  `protobuf:"varint,9,opt,name=life,proto3" json:"life,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_schema_fleet_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetFleetNumber() string {
	if x != nil {
		return x.FleetNumber
	}
	return ""
}

func (x *Vehicle) GetOperationalStatus() string {
	if x != nil {
		return x.OperationalStatus
	}
	return ""
}

func (x *Vehicle) GetLfb() string {
	if x != nil {
		return x.Lfb
	}
	return ""
}

func (x *Vehicle) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *Vehicle) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Vehicle) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Vehicle) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Vehicle) GetRegistrationYear() int32 {
	if x != nil {
		return x.RegistrationYear
	}
	return 0
}

func (x *Vehicle) GetLife() int32 {
	if x != nil {
		return x.Life
	}
	return 0
}

type VehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationalStatus string `protobuf:"bytes,2,opt,name=operational_status,json=operationalStatus,proto3" json:"operational_status,omitempty"`
	Lfb               string `protobuf:"bytes,3,opt,name=lfb,proto3" json:"lfb,omitempty"`
	Make              string `protobuf:"bytes,4,opt,name=make,proto3" json:"make,omitempty"`
	Model             string `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	Type              string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Category          string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	RegistrationYear  int32  `protobuf:"varint,8,opt,name=registration_year,json=registrationYear,proto3" json:"registration_year,omitempty"`
	Life              int32  `protobuf:"varint,9,opt,name=life,proto3" json:"life,omitempty"`
}

func (x *VehicleRequest) Reset() {
	*x = VehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleRequest) ProtoMessage() {}

func (x *VehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleRequest.ProtoReflect.Descriptor instead.
func (*VehicleRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_schema_fleet_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleRequest) GetOperationalStatus() string {
	if x != nil {
		return x.OperationalStatus
	}
	return ""
}

func (x *VehicleRequest) GetLfb() string {
	if x != nil {
		return x.Lfb
	}
	return ""
}

func (x *VehicleRequest) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *VehicleRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *VehicleRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VehicleRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *VehicleRequest) GetRegistrationYear() int32 {
	if x != nil {
		return x.RegistrationYear
	}
	return 0
}

func (x *VehicleRequest) GetLife() int32 {
	if x != nil {
		return x.Life
	}
	return 0
}

type VehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	Created bool     `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *VehicleResponse) Reset() {
	*x = VehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleResponse) ProtoMessage() {}

func (x *VehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleResponse.ProtoReflect.Descriptor instead.
func (*VehicleResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_schema_fleet_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleResponse) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

func (x *VehicleResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

type VehicleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicles []*Vehicle `protobuf:"bytes,1,rep,name=vehicles,proto3" json:"vehicles,omitempty"`
}

func (x *VehicleList) Reset() {
	*x = VehicleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleList) ProtoMessage() {}

func (x *VehicleList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_schema_fleet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleList.ProtoReflect.Descriptor instead.
func (*VehicleList) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_schema_fleet_proto_rawDescGZIP(), []int{3}
}

func (x *VehicleList) GetVehicles() []*Vehicle {
	if x != nil {
		return x.Vehicles
	}
	return nil
}

var File_internal_protobuf_schema_fleet_proto protoreflect.FileDescriptor

var file_internal_protobuf_schema_fleet_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x07, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x66, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x66, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61,
	0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6c, 0x69, 0x66, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x66, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x66, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x6c, 0x69, 0x66, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x39, 0x0a, 0x0b, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32, 0xe0, 0x02, 0x0a, 0x0c, 0x46, 0x6c, 0x65, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x12, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x15, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_schema_fleet_proto_rawDescOnce sync.Once
	file_internal_protobuf_schema_fleet_proto_rawDescData = file_internal_protobuf_schema_fleet_proto_rawDesc
)

func file_internal_protobuf_schema_fleet_proto_rawDescGZIP() []byte {
	file_internal_protobuf_schema_fleet_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_schema_fleet_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_schema_fleet_proto_rawDescData)
	})
	return file_internal_protobuf_schema_fleet_proto_rawDescData
}

var file_internal_protobuf_schema_fleet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_protobuf_schema_fleet_proto_goTypes = []interface{}{
	(*Vehicle)(nil),         // 0: fleet.Vehicle
	(*VehicleRequest)(nil),  // 1: fleet.VehicleRequest
	(*VehicleResponse)(nil), // 2: fleet.VehicleResponse
	(*VehicleList)(nil),     // 3: fleet.VehicleList
	(*emptypb.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_internal_protobuf_schema_fleet_proto_depIdxs = []int32{
	0, // 0: fleet.VehicleResponse.vehicle:type_name -> fleet.Vehicle
	0, // 1: fleet.VehicleList.vehicles:type_name -> fleet.Vehicle
	4, // 2: fleet.FleetService.ListVehicles:input_type -> google.protobuf.Empty
	1, // 3: fleet.FleetService.AddVehicle:input_type -> fleet.VehicleRequest
	1, // 4: fleet.FleetService.GetTrainingVehicles:input_type -> fleet.VehicleRequest
	1, // 5: fleet.FleetService.GetReserveVehicles:input_type -> fleet.VehicleRequest
	1, // 6: fleet.FleetService.GetActiveVehicles:input_type -> fleet.VehicleRequest
	3, // 7: fleet.FleetService.ListVehicles:output_type -> fleet.VehicleList
	2, // 8: fleet.FleetService.AddVehicle:output_type -> fleet.VehicleResponse
	2, // 9: fleet.FleetService.GetTrainingVehicles:output_type -> fleet.VehicleResponse
	2, // 10: fleet.FleetService.GetReserveVehicles:output_type -> fleet.VehicleResponse
	2, // 11: fleet.FleetService.GetActiveVehicles:output_type -> fleet.VehicleResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_protobuf_schema_fleet_proto_init() }
func file_internal_protobuf_schema_fleet_proto_init() {
	if File_internal_protobuf_schema_fleet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_schema_fleet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
		file_internal_protobuf_schema_fleet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleRequest); i {
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
		file_internal_protobuf_schema_fleet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleResponse); i {
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
		file_internal_protobuf_schema_fleet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleList); i {
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
			RawDescriptor: file_internal_protobuf_schema_fleet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_schema_fleet_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_schema_fleet_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_schema_fleet_proto_msgTypes,
	}.Build()
	File_internal_protobuf_schema_fleet_proto = out.File
	file_internal_protobuf_schema_fleet_proto_rawDesc = nil
	file_internal_protobuf_schema_fleet_proto_goTypes = nil
	file_internal_protobuf_schema_fleet_proto_depIdxs = nil
}
