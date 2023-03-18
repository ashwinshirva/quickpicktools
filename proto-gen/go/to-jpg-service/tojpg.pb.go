// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: to-jpg-service/tojpg.proto

package to_jpg_worker

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Status int32

const (
	Status_NOT_DEFINED Status = 0
	Status_SUCCESS     Status = 1
	Status_FAILURE     Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "NOT_DEFINED",
		1: "SUCCESS",
		2: "FAILURE",
	}
	Status_value = map[string]int32{
		"NOT_DEFINED": 0,
		"SUCCESS":     1,
		"FAILURE":     2,
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
	return file_to_jpg_service_tojpg_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_to_jpg_service_tojpg_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{0}
}

type PngToJpgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Image    *Image `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *PngToJpgReq) Reset() {
	*x = PngToJpgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_jpg_service_tojpg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PngToJpgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PngToJpgReq) ProtoMessage() {}

func (x *PngToJpgReq) ProtoReflect() protoreflect.Message {
	mi := &file_to_jpg_service_tojpg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PngToJpgReq.ProtoReflect.Descriptor instead.
func (*PngToJpgReq) Descriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{0}
}

func (x *PngToJpgReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *PngToJpgReq) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *ImageMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Data     []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_jpg_service_tojpg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_to_jpg_service_tojpg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetMetadata() *ImageMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Image) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImageMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ImageMetadata) Reset() {
	*x = ImageMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_jpg_service_tojpg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMetadata) ProtoMessage() {}

func (x *ImageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_to_jpg_service_tojpg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMetadata.ProtoReflect.Descriptor instead.
func (*ImageMetadata) Descriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{2}
}

func (x *ImageMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PngToJpgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            *ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ConvertedImageUrl string          `protobuf:"bytes,2,opt,name=ConvertedImageUrl,proto3" json:"ConvertedImageUrl,omitempty"`
	ConvertedImage    []byte          `protobuf:"bytes,3,opt,name=converted_image,json=convertedImage,proto3" json:"converted_image,omitempty"`
}

func (x *PngToJpgResp) Reset() {
	*x = PngToJpgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_jpg_service_tojpg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PngToJpgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PngToJpgResp) ProtoMessage() {}

func (x *PngToJpgResp) ProtoReflect() protoreflect.Message {
	mi := &file_to_jpg_service_tojpg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PngToJpgResp.ProtoReflect.Descriptor instead.
func (*PngToJpgResp) Descriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{3}
}

func (x *PngToJpgResp) GetStatus() *ResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *PngToJpgResp) GetConvertedImageUrl() string {
	if x != nil {
		return x.ConvertedImageUrl
	}
	return ""
}

func (x *PngToJpgResp) GetConvertedImage() []byte {
	if x != nil {
		return x.ConvertedImage
	}
	return nil
}

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_to_jpg_service_tojpg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_to_jpg_service_tojpg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_to_jpg_service_tojpg_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_NOT_DEFINED
}

func (x *ResponseStatus) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_to_jpg_service_tojpg_proto protoreflect.FileDescriptor

var file_to_jpg_service_tojpg_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6f, 0x2d, 0x6a, 0x70, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x74, 0x6f, 0x6a, 0x70, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x0b, 0x50, 0x6e, 0x67, 0x54, 0x6f, 0x4a, 0x70, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x4d, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x23, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x50, 0x6e, 0x67, 0x54, 0x6f, 0x4a,
	0x70, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x33, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x32, 0x62, 0x0a, 0x0c,
	0x54, 0x6f, 0x4a, 0x70, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x08,
	0x50, 0x6e, 0x67, 0x54, 0x6f, 0x4a, 0x70, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6e, 0x67, 0x54, 0x6f, 0x4a, 0x70, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6e, 0x67, 0x54, 0x6f, 0x4a, 0x70, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x74,
	0x6f, 0x2d, 0x6a, 0x70, 0x67, 0x2f, 0x70, 0x6e, 0x67, 0x2d, 0x74, 0x6f, 0x2d, 0x6a, 0x70, 0x67,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71,
	0x75, 0x69, 0x63, 0x6b, 0x70, 0x69, 0x63, 0x6b, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x74, 0x6f,
	0x2d, 0x6a, 0x70, 0x67, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_to_jpg_service_tojpg_proto_rawDescOnce sync.Once
	file_to_jpg_service_tojpg_proto_rawDescData = file_to_jpg_service_tojpg_proto_rawDesc
)

func file_to_jpg_service_tojpg_proto_rawDescGZIP() []byte {
	file_to_jpg_service_tojpg_proto_rawDescOnce.Do(func() {
		file_to_jpg_service_tojpg_proto_rawDescData = protoimpl.X.CompressGZIP(file_to_jpg_service_tojpg_proto_rawDescData)
	})
	return file_to_jpg_service_tojpg_proto_rawDescData
}

var file_to_jpg_service_tojpg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_to_jpg_service_tojpg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_to_jpg_service_tojpg_proto_goTypes = []interface{}{
	(Status)(0),            // 0: proto.Status
	(*PngToJpgReq)(nil),    // 1: proto.PngToJpgReq
	(*Image)(nil),          // 2: proto.Image
	(*ImageMetadata)(nil),  // 3: proto.ImageMetadata
	(*PngToJpgResp)(nil),   // 4: proto.PngToJpgResp
	(*ResponseStatus)(nil), // 5: proto.ResponseStatus
}
var file_to_jpg_service_tojpg_proto_depIdxs = []int32{
	2, // 0: proto.PngToJpgReq.image:type_name -> proto.Image
	3, // 1: proto.Image.metadata:type_name -> proto.ImageMetadata
	5, // 2: proto.PngToJpgResp.status:type_name -> proto.ResponseStatus
	0, // 3: proto.ResponseStatus.status:type_name -> proto.Status
	1, // 4: proto.ToJpgService.PngToJpg:input_type -> proto.PngToJpgReq
	4, // 5: proto.ToJpgService.PngToJpg:output_type -> proto.PngToJpgResp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_to_jpg_service_tojpg_proto_init() }
func file_to_jpg_service_tojpg_proto_init() {
	if File_to_jpg_service_tojpg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_to_jpg_service_tojpg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PngToJpgReq); i {
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
		file_to_jpg_service_tojpg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_to_jpg_service_tojpg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageMetadata); i {
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
		file_to_jpg_service_tojpg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PngToJpgResp); i {
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
		file_to_jpg_service_tojpg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
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
			RawDescriptor: file_to_jpg_service_tojpg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_to_jpg_service_tojpg_proto_goTypes,
		DependencyIndexes: file_to_jpg_service_tojpg_proto_depIdxs,
		EnumInfos:         file_to_jpg_service_tojpg_proto_enumTypes,
		MessageInfos:      file_to_jpg_service_tojpg_proto_msgTypes,
	}.Build()
	File_to_jpg_service_tojpg_proto = out.File
	file_to_jpg_service_tojpg_proto_rawDesc = nil
	file_to_jpg_service_tojpg_proto_goTypes = nil
	file_to_jpg_service_tojpg_proto_depIdxs = nil
}
