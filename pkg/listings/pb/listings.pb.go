// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pkg/listings/pb/listings.proto

package pb

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Listing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Uri         string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tags        []*Tag `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Listing) Reset() {
	*x = Listing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listing) ProtoMessage() {}

func (x *Listing) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listing.ProtoReflect.Descriptor instead.
func (*Listing) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{1}
}

func (x *Listing) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Listing) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Listing) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Listing) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Listing) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Listing) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Listing) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Listing) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreateListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Chunk       []byte   `protobuf:"bytes,4,opt,name=chunk,proto3" json:"chunk,omitempty"`
	TagNames    []string `protobuf:"bytes,5,rep,name=tag_names,json=tagNames,proto3" json:"tag_names,omitempty"`
}

func (x *CreateListingRequest) Reset() {
	*x = CreateListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListingRequest) ProtoMessage() {}

func (x *CreateListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListingRequest.ProtoReflect.Descriptor instead.
func (*CreateListingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{2}
}

func (x *CreateListingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateListingRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateListingRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateListingRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *CreateListingRequest) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

type CreateListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CreateListingResponse) Reset() {
	*x = CreateListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListingResponse) ProtoMessage() {}

func (x *CreateListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListingResponse.ProtoReflect.Descriptor instead.
func (*CreateListingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{3}
}

func (x *CreateListingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateListingResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type GetListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetListingRequest) Reset() {
	*x = GetListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingRequest) ProtoMessage() {}

func (x *GetListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingRequest.ProtoReflect.Descriptor instead.
func (*GetListingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{4}
}

func (x *GetListingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Uri         string   `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	TagNames    []string `protobuf:"bytes,6,rep,name=tag_names,json=tagNames,proto3" json:"tag_names,omitempty"`
	CreatedAt   string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetListingResponse) Reset() {
	*x = GetListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingResponse) ProtoMessage() {}

func (x *GetListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingResponse.ProtoReflect.Descriptor instead.
func (*GetListingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{5}
}

func (x *GetListingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetListingResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetListingResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetListingResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GetListingResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *GetListingResponse) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

func (x *GetListingResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetListingResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetListingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listings []*Listing `protobuf:"bytes,1,rep,name=listings,proto3" json:"listings,omitempty"`
}

func (x *GetListingsResponse) Reset() {
	*x = GetListingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingsResponse) ProtoMessage() {}

func (x *GetListingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingsResponse.ProtoReflect.Descriptor instead.
func (*GetListingsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{6}
}

func (x *GetListingsResponse) GetListings() []*Listing {
	if x != nil {
		return x.Listings
	}
	return nil
}

type UpdateListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	TagNames    []string `protobuf:"bytes,5,rep,name=tag_names,json=tagNames,proto3" json:"tag_names,omitempty"`
}

func (x *UpdateListingRequest) Reset() {
	*x = UpdateListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListingRequest) ProtoMessage() {}

func (x *UpdateListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListingRequest.ProtoReflect.Descriptor instead.
func (*UpdateListingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateListingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateListingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateListingRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateListingRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *UpdateListingRequest) GetTagNames() []string {
	if x != nil {
		return x.TagNames
	}
	return nil
}

type UpdateListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateListingResponse) Reset() {
	*x = UpdateListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListingResponse) ProtoMessage() {}

func (x *UpdateListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListingResponse.ProtoReflect.Descriptor instead.
func (*UpdateListingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateListingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateListingStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateListingStatusRequest) Reset() {
	*x = UpdateListingStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListingStatusRequest) ProtoMessage() {}

func (x *UpdateListingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListingStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateListingStatusRequest) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateListingStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateListingStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateListingStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateListingStatusResponse) Reset() {
	*x = UpdateListingStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListingStatusResponse) ProtoMessage() {}

func (x *UpdateListingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListingStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateListingStatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateListingStatusResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteListingRequest) Reset() {
	*x = DeleteListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_listings_pb_listings_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteListingRequest) ProtoMessage() {}

func (x *DeleteListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_listings_pb_listings_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteListingRequest.ProtoReflect.Descriptor instead.
func (*DeleteListingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_listings_pb_listings_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteListingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pkg_listings_pb_listings_proto protoreflect.FileDescriptor

var file_pkg_listings_pb_listings_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x95, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x27, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x1b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xfd, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1e, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_listings_pb_listings_proto_rawDescOnce sync.Once
	file_pkg_listings_pb_listings_proto_rawDescData = file_pkg_listings_pb_listings_proto_rawDesc
)

func file_pkg_listings_pb_listings_proto_rawDescGZIP() []byte {
	file_pkg_listings_pb_listings_proto_rawDescOnce.Do(func() {
		file_pkg_listings_pb_listings_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_listings_pb_listings_proto_rawDescData)
	})
	return file_pkg_listings_pb_listings_proto_rawDescData
}

var file_pkg_listings_pb_listings_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_listings_pb_listings_proto_goTypes = []interface{}{
	(*Tag)(nil),                         // 0: listings.Tag
	(*Listing)(nil),                     // 1: listings.Listing
	(*CreateListingRequest)(nil),        // 2: listings.CreateListingRequest
	(*CreateListingResponse)(nil),       // 3: listings.CreateListingResponse
	(*GetListingRequest)(nil),           // 4: listings.GetListingRequest
	(*GetListingResponse)(nil),          // 5: listings.GetListingResponse
	(*GetListingsResponse)(nil),         // 6: listings.GetListingsResponse
	(*UpdateListingRequest)(nil),        // 7: listings.UpdateListingRequest
	(*UpdateListingResponse)(nil),       // 8: listings.UpdateListingResponse
	(*UpdateListingStatusRequest)(nil),  // 9: listings.UpdateListingStatusRequest
	(*UpdateListingStatusResponse)(nil), // 10: listings.UpdateListingStatusResponse
	(*DeleteListingRequest)(nil),        // 11: listings.DeleteListingRequest
	(*emptypb.Empty)(nil),               // 12: google.protobuf.Empty
}
var file_pkg_listings_pb_listings_proto_depIdxs = []int32{
	0,  // 0: listings.Listing.tags:type_name -> listings.Tag
	1,  // 1: listings.GetListingsResponse.listings:type_name -> listings.Listing
	2,  // 2: listings.ListingsService.CreateListing:input_type -> listings.CreateListingRequest
	4,  // 3: listings.ListingsService.GetListing:input_type -> listings.GetListingRequest
	12, // 4: listings.ListingsService.GetListings:input_type -> google.protobuf.Empty
	7,  // 5: listings.ListingsService.UpdateListing:input_type -> listings.UpdateListingRequest
	9,  // 6: listings.ListingsService.UpdateListingStatus:input_type -> listings.UpdateListingStatusRequest
	11, // 7: listings.ListingsService.DeleteListing:input_type -> listings.DeleteListingRequest
	3,  // 8: listings.ListingsService.CreateListing:output_type -> listings.CreateListingResponse
	5,  // 9: listings.ListingsService.GetListing:output_type -> listings.GetListingResponse
	6,  // 10: listings.ListingsService.GetListings:output_type -> listings.GetListingsResponse
	8,  // 11: listings.ListingsService.UpdateListing:output_type -> listings.UpdateListingResponse
	10, // 12: listings.ListingsService.UpdateListingStatus:output_type -> listings.UpdateListingStatusResponse
	12, // 13: listings.ListingsService.DeleteListing:output_type -> google.protobuf.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_listings_pb_listings_proto_init() }
func file_pkg_listings_pb_listings_proto_init() {
	if File_pkg_listings_pb_listings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_listings_pb_listings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listing); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListingRequest); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListingResponse); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingRequest); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingResponse); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListingsResponse); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListingRequest); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListingResponse); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListingStatusRequest); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListingStatusResponse); i {
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
		file_pkg_listings_pb_listings_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteListingRequest); i {
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
			RawDescriptor: file_pkg_listings_pb_listings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_listings_pb_listings_proto_goTypes,
		DependencyIndexes: file_pkg_listings_pb_listings_proto_depIdxs,
		MessageInfos:      file_pkg_listings_pb_listings_proto_msgTypes,
	}.Build()
	File_pkg_listings_pb_listings_proto = out.File
	file_pkg_listings_pb_listings_proto_rawDesc = nil
	file_pkg_listings_pb_listings_proto_goTypes = nil
	file_pkg_listings_pb_listings_proto_depIdxs = nil
}
