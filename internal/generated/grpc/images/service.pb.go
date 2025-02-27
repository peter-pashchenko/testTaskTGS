// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.6.1
// source: schemas/proto/grpc/service.proto

package image_service_v1

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{0}
}

type SaveImageReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SaveImageReponse) Reset() {
	*x = SaveImageReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveImageReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveImageReponse) ProtoMessage() {}

func (x *SaveImageReponse) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveImageReponse.ProtoReflect.Descriptor instead.
func (*SaveImageReponse) Descriptor() ([]byte, []int) {
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *SaveImageReponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[3]
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
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Created *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ImageList) Reset() {
	*x = ImageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageList) ProtoMessage() {}

func (x *ImageList) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageList.ProtoReflect.Descriptor instead.
func (*ImageList) Descriptor() ([]byte, []int) {
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *ImageList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageList) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ImageList) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type ListImages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*ImageList `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ListImages) Reset() {
	*x = ListImages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schemas_proto_grpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImages) ProtoMessage() {}

func (x *ListImages) ProtoReflect() protoreflect.Message {
	mi := &file_schemas_proto_grpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImages.ProtoReflect.Descriptor instead.
func (*ListImages) Descriptor() ([]byte, []int) {
	return file_schemas_proto_grpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListImages) GetImages() []*ImageList {
	if x != nil {
		return x.Images
	}
	return nil
}

var File_schemas_proto_grpc_service_proto protoreflect.FileDescriptor

var file_schemas_proto_grpc_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a,
	0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1c, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x32, 0xd8, 0x01, 0x0a, 0x0c, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x53,
	0x61, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x1a, 0x22, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schemas_proto_grpc_service_proto_rawDescOnce sync.Once
	file_schemas_proto_grpc_service_proto_rawDescData = file_schemas_proto_grpc_service_proto_rawDesc
)

func file_schemas_proto_grpc_service_proto_rawDescGZIP() []byte {
	file_schemas_proto_grpc_service_proto_rawDescOnce.Do(func() {
		file_schemas_proto_grpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_schemas_proto_grpc_service_proto_rawDescData)
	})
	return file_schemas_proto_grpc_service_proto_rawDescData
}

var file_schemas_proto_grpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_schemas_proto_grpc_service_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: image_service.v1.Empty
	(*SaveImageReponse)(nil),    // 1: image_service.v1.SaveImageReponse
	(*Name)(nil),                // 2: image_service.v1.Name
	(*Image)(nil),               // 3: image_service.v1.Image
	(*ImageList)(nil),           // 4: image_service.v1.ImageList
	(*ListImages)(nil),          // 5: image_service.v1.ListImages
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_schemas_proto_grpc_service_proto_depIdxs = []int32{
	6, // 0: image_service.v1.ImageList.created:type_name -> google.protobuf.Timestamp
	6, // 1: image_service.v1.ImageList.updated:type_name -> google.protobuf.Timestamp
	4, // 2: image_service.v1.ListImages.images:type_name -> image_service.v1.ImageList
	3, // 3: image_service.v1.imageService.SaveImage:input_type -> image_service.v1.Image
	2, // 4: image_service.v1.imageService.GetByName:input_type -> image_service.v1.Name
	0, // 5: image_service.v1.imageService.ListAll:input_type -> image_service.v1.Empty
	1, // 6: image_service.v1.imageService.SaveImage:output_type -> image_service.v1.SaveImageReponse
	3, // 7: image_service.v1.imageService.GetByName:output_type -> image_service.v1.Image
	5, // 8: image_service.v1.imageService.ListAll:output_type -> image_service.v1.ListImages
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schemas_proto_grpc_service_proto_init() }
func file_schemas_proto_grpc_service_proto_init() {
	if File_schemas_proto_grpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schemas_proto_grpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_schemas_proto_grpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveImageReponse); i {
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
		file_schemas_proto_grpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_schemas_proto_grpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_schemas_proto_grpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageList); i {
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
		file_schemas_proto_grpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImages); i {
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
			RawDescriptor: file_schemas_proto_grpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schemas_proto_grpc_service_proto_goTypes,
		DependencyIndexes: file_schemas_proto_grpc_service_proto_depIdxs,
		MessageInfos:      file_schemas_proto_grpc_service_proto_msgTypes,
	}.Build()
	File_schemas_proto_grpc_service_proto = out.File
	file_schemas_proto_grpc_service_proto_rawDesc = nil
	file_schemas_proto_grpc_service_proto_goTypes = nil
	file_schemas_proto_grpc_service_proto_depIdxs = nil
}
