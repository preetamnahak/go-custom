// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/product/product.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	mi := &file_proto_product_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_proto_product_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_proto_product_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductResponse     `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_proto_product_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *SearchResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_proto_product_product_proto protoreflect.FileDescriptor

const file_proto_product_product_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/product/product.proto\x12\aproduct\" \n" +
	"\x0eProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"m\n" +
	"\x0fProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\";\n" +
	"\rSearchRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\"F\n" +
	"\x0eSearchResponse\x124\n" +
	"\bproducts\x18\x01 \x03(\v2\x18.product.ProductResponseR\bproducts2\x98\x01\n" +
	"\x0eProductService\x12A\n" +
	"\n" +
	"GetProduct\x12\x17.product.ProductRequest\x1a\x18.product.ProductResponse\"\x00\x12C\n" +
	"\x0eSearchProducts\x12\x16.product.SearchRequest\x1a\x17.product.SearchResponse\"\x00B2Z0github.com/preetamnahak/mypackages/proto/productb\x06proto3"

var (
	file_proto_product_product_proto_rawDescOnce sync.Once
	file_proto_product_product_proto_rawDescData []byte
)

func file_proto_product_product_proto_rawDescGZIP() []byte {
	file_proto_product_product_proto_rawDescOnce.Do(func() {
		file_proto_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_product_product_proto_rawDesc), len(file_proto_product_product_proto_rawDesc)))
	})
	return file_proto_product_product_proto_rawDescData
}

var file_proto_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_product_product_proto_goTypes = []any{
	(*ProductRequest)(nil),  // 0: product.ProductRequest
	(*ProductResponse)(nil), // 1: product.ProductResponse
	(*SearchRequest)(nil),   // 2: product.SearchRequest
	(*SearchResponse)(nil),  // 3: product.SearchResponse
}
var file_proto_product_product_proto_depIdxs = []int32{
	1, // 0: product.SearchResponse.products:type_name -> product.ProductResponse
	0, // 1: product.ProductService.GetProduct:input_type -> product.ProductRequest
	2, // 2: product.ProductService.SearchProducts:input_type -> product.SearchRequest
	1, // 3: product.ProductService.GetProduct:output_type -> product.ProductResponse
	3, // 4: product.ProductService.SearchProducts:output_type -> product.SearchResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_product_product_proto_init() }
func file_proto_product_product_proto_init() {
	if File_proto_product_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_product_product_proto_rawDesc), len(file_proto_product_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_product_proto_goTypes,
		DependencyIndexes: file_proto_product_product_proto_depIdxs,
		MessageInfos:      file_proto_product_product_proto_msgTypes,
	}.Build()
	File_proto_product_product_proto = out.File
	file_proto_product_product_proto_goTypes = nil
	file_proto_product_product_proto_depIdxs = nil
}
