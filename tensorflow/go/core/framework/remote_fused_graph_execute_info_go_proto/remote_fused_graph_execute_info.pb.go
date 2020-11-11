// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

package remote_fused_graph_execute_info_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	graph_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/graph_go_proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *graph_go_proto.GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph,proto3" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName,proto3" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName,proto3" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape,proto3" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape,proto3" json:"default_graph_output_tensor_shape,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                                            `json:"-"`
	XXX_unrecognized              []byte                                              `json:"-"`
	XXX_sizecache                 int32                                               `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()         { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()    {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0}
}

func (m *RemoteFusedGraphExecuteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *graph_go_proto.GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype                types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptor_c15f13da5b37f691)
}

var fileDescriptor_c15f13da5b37f691 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xad, 0x5d, 0xd9, 0x69, 0xf5, 0x30, 0xae, 0x12, 0x6a, 0x95, 0xa8, 0x08, 0x41,
	0x24, 0xc5, 0xee, 0xc1, 0x8b, 0x20, 0x2c, 0x5d, 0x97, 0xbd, 0xac, 0x25, 0xee, 0xc9, 0xcb, 0x30,
	0xdb, 0xbc, 0xa4, 0xc1, 0x26, 0x2f, 0x4c, 0x26, 0xae, 0xeb, 0x59, 0xfc, 0x5f, 0x3c, 0xfa, 0xdf,
	0x79, 0x94, 0xbc, 0x89, 0xe9, 0x14, 0xb6, 0x05, 0x61, 0x6f, 0xd3, 0xbe, 0xdf, 0xf7, 0xe6, 0x7b,
	0x5f, 0xde, 0xf0, 0xf7, 0x06, 0x8a, 0x0a, 0x75, 0xb2, 0xc2, 0xab, 0xc9, 0x02, 0x35, 0x4c, 0x12,
	0xad, 0x72, 0xb8, 0x42, 0xfd, 0x65, 0xa2, 0x21, 0x47, 0x03, 0x32, 0xa9, 0x2b, 0x88, 0x65, 0xaa,
	0x55, 0xb9, 0x94, 0xf0, 0x0d, 0x16, 0xb5, 0x01, 0x99, 0x15, 0x09, 0x86, 0xa5, 0x46, 0x83, 0x82,
	0xaf, 0x1b, 0x8c, 0x5e, 0x6e, 0x6f, 0x46, 0x7a, 0x2b, 0x19, 0xbd, 0xde, 0x8e, 0xd9, 0x8a, 0xac,
	0x96, 0xaa, 0x84, 0x96, 0xde, 0xd1, 0xd4, 0x5c, 0x97, 0x50, 0x59, 0xec, 0xf9, 0xef, 0x3e, 0x7f,
	0x1c, 0x91, 0xe3, 0x0f, 0x8d, 0xe1, 0xd3, 0xe6, 0xbe, 0x13, 0x6b, 0xf7, 0xac, 0x48, 0x50, 0xbc,
	0xe5, 0xc3, 0x76, 0x20, 0xb2, 0xe2, 0x31, 0x9f, 0x05, 0x83, 0xe9, 0x61, 0xb8, 0xee, 0x1e, 0x92,
	0x66, 0x06, 0x49, 0x34, 0xb0, 0x24, 0xfd, 0x16, 0x6f, 0xf8, 0x43, 0x3b, 0x7c, 0x56, 0x94, 0xb5,
	0x91, 0x05, 0xc6, 0x20, 0x0b, 0x95, 0x83, 0xb7, 0xe7, 0xf7, 0x82, 0x83, 0x48, 0x50, 0xf1, 0xac,
	0xa9, 0x9d, 0x63, 0x0c, 0xe7, 0x2a, 0x07, 0x71, 0xc4, 0x1f, 0x59, 0x09, 0xd6, 0x66, 0x53, 0xd3,
	0x23, 0xcd, 0x03, 0xaa, 0x7e, 0xa4, 0x62, 0x27, 0x7a, 0xc1, 0xef, 0xd9, 0x78, 0x51, 0x5b, 0xf6,
	0x8e, 0xcf, 0x82, 0x83, 0x68, 0xf8, 0xef, 0x4f, 0x82, 0x66, 0xfc, 0x69, 0x05, 0x3a, 0x53, 0xab,
	0xec, 0x3b, 0xc4, 0xb2, 0xe3, 0x4b, 0xd5, 0x64, 0x62, 0x40, 0x57, 0x5e, 0xdf, 0x67, 0xc1, 0x30,
	0x1a, 0xaf, 0xa9, 0x93, 0x16, 0x9a, 0x77, 0x8c, 0xf8, 0xc1, 0xb8, 0x1f, 0x43, 0xa2, 0xea, 0x95,
	0x91, 0xee, 0x6c, 0x6e, 0xfa, 0xde, 0xbe, 0xdf, 0x0b, 0x06, 0xd3, 0x77, 0x6e, 0x40, 0x3b, 0xf2,
	0x0d, 0x2f, 0x08, 0xfb, 0xd4, 0x48, 0x2f, 0xae, 0x4b, 0x98, 0x37, 0x1f, 0x25, 0x1a, 0xb7, 0xb7,
	0x9c, 0x76, 0x19, 0x39, 0x98, 0xf8, 0xc9, 0xf8, 0xb3, 0x4d, 0x1b, 0x6d, 0x5e, 0x1b, 0x3e, 0xee,
	0xde, 0x82, 0x8f, 0x27, 0xae, 0x0f, 0x9b, 0xbb, 0xc3, 0x8d, 0xbe, 0xf2, 0xc3, 0x9b, 0x64, 0xe2,
	0x15, 0xef, 0xc7, 0xcd, 0x8e, 0xd1, 0xb2, 0xdc, 0xdf, 0x5c, 0x96, 0x99, 0x32, 0xaa, 0x21, 0x23,
	0x8b, 0x88, 0x29, 0xef, 0x5b, 0xbf, 0x7b, 0xb4, 0x58, 0x63, 0x97, 0x75, 0x9a, 0x5b, 0x3f, 0x16,
	0x3d, 0xfe, 0xc5, 0xb8, 0x87, 0x3a, 0x75, 0xd1, 0x6e, 0xb9, 0x8f, 0xfd, 0x1d, 0x53, 0x52, 0x97,
	0x39, 0xfb, 0x9c, 0xa4, 0x99, 0x59, 0xd6, 0x97, 0xe1, 0x02, 0xf3, 0x89, 0xf3, 0x4c, 0x6e, 0x3e,
	0xa6, 0xf8, 0x9f, 0x2f, 0x5c, 0xa6, 0x28, 0xe9, 0x71, 0xfd, 0x61, 0xec, 0x72, 0x9f, 0x4e, 0x47,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0x61, 0xb0, 0x68, 0x31, 0x04, 0x00, 0x00,
}