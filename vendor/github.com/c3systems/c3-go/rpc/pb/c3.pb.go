// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/pb/c3.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Jsonrpc              string   `protobuf:"bytes,1,opt,name=jsonrpc,proto3" json:"jsonrpc,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Params               []string `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetJsonrpc() string {
	if m != nil {
		return m.Jsonrpc
	}
	return ""
}

func (m *Request) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Response struct {
	Jsonrpc              string   `protobuf:"bytes,1,opt,name=jsonrpc,proto3" json:"jsonrpc,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Result               *any.Any `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetJsonrpc() string {
	if m != nil {
		return m.Jsonrpc
	}
	return ""
}

func (m *Response) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

type ErrorResponse struct {
	Code                 uint64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{2}
}
func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (dst *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(dst, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{3}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type LatestBlockResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatestBlockResponse) Reset()         { *m = LatestBlockResponse{} }
func (m *LatestBlockResponse) String() string { return proto.CompactTextString(m) }
func (*LatestBlockResponse) ProtoMessage()    {}
func (*LatestBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{4}
}
func (m *LatestBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestBlockResponse.Unmarshal(m, b)
}
func (m *LatestBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestBlockResponse.Marshal(b, m, deterministic)
}
func (dst *LatestBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestBlockResponse.Merge(dst, src)
}
func (m *LatestBlockResponse) XXX_Size() int {
	return xxx_messageInfo_LatestBlockResponse.Size(m)
}
func (m *LatestBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestBlockResponse proto.InternalMessageInfo

func (m *LatestBlockResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type BlockResponse struct {
	BlockHash             string     `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber           string     `protobuf:"bytes,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockTime             string     `protobuf:"bytes,3,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	ImageHash             string     `protobuf:"bytes,4,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
	StateBlocksMerkleHash string     `protobuf:"bytes,5,opt,name=stateBlocksMerkleHash,proto3" json:"stateBlocksMerkleHash,omitempty"`
	PrevBlockHash         string     `protobuf:"bytes,6,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Nonce                 string     `protobuf:"bytes,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Difficulty            string     `protobuf:"bytes,8,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	MinerAddress          string     `protobuf:"bytes,9,opt,name=minerAddress,proto3" json:"minerAddress,omitempty"`
	MinerSig              *Signature `protobuf:"bytes,10,opt,name=minerSig,proto3" json:"minerSig,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-"`
	XXX_unrecognized      []byte     `json:"-"`
	XXX_sizecache         int32      `json:"-"`
}

func (m *BlockResponse) Reset()         { *m = BlockResponse{} }
func (m *BlockResponse) String() string { return proto.CompactTextString(m) }
func (*BlockResponse) ProtoMessage()    {}
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{5}
}
func (m *BlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockResponse.Unmarshal(m, b)
}
func (m *BlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockResponse.Marshal(b, m, deterministic)
}
func (dst *BlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponse.Merge(dst, src)
}
func (m *BlockResponse) XXX_Size() int {
	return xxx_messageInfo_BlockResponse.Size(m)
}
func (m *BlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponse proto.InternalMessageInfo

func (m *BlockResponse) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *BlockResponse) GetBlockNumber() string {
	if m != nil {
		return m.BlockNumber
	}
	return ""
}

func (m *BlockResponse) GetBlockTime() string {
	if m != nil {
		return m.BlockTime
	}
	return ""
}

func (m *BlockResponse) GetImageHash() string {
	if m != nil {
		return m.ImageHash
	}
	return ""
}

func (m *BlockResponse) GetStateBlocksMerkleHash() string {
	if m != nil {
		return m.StateBlocksMerkleHash
	}
	return ""
}

func (m *BlockResponse) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *BlockResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *BlockResponse) GetDifficulty() string {
	if m != nil {
		return m.Difficulty
	}
	return ""
}

func (m *BlockResponse) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *BlockResponse) GetMinerSig() *Signature {
	if m != nil {
		return m.MinerSig
	}
	return nil
}

type Signature struct {
	R                    string   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	S                    string   `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{6}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *Signature) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type TransactionResponse struct {
	TxHash               string     `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	ImageHash            string     `protobuf:"bytes,2,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
	Method               string     `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Payload              []string   `protobuf:"bytes,4,rep,name=payload,proto3" json:"payload,omitempty"`
	From                 string     `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Sig                  *Signature `protobuf:"bytes,6,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{7}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (dst *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(dst, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *TransactionResponse) GetImageHash() string {
	if m != nil {
		return m.ImageHash
	}
	return ""
}

func (m *TransactionResponse) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *TransactionResponse) GetPayload() []string {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransactionResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionResponse) GetSig() *Signature {
	if m != nil {
		return m.Sig
	}
	return nil
}

type StateBlockResponse struct {
	BlockHash            string   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber          string   `protobuf:"bytes,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockTime            string   `protobuf:"bytes,3,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	ImageHash            string   `protobuf:"bytes,4,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
	TxHash               string   `protobuf:"bytes,5,opt,name=txHash,proto3" json:"txHash,omitempty"`
	PrevBlockHash        string   `protobuf:"bytes,6,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	StatePrevDiffHash    string   `protobuf:"bytes,7,opt,name=statePrevDiffHash,proto3" json:"statePrevDiffHash,omitempty"`
	StateCurrentHash     string   `protobuf:"bytes,8,opt,name=stateCurrentHash,proto3" json:"stateCurrentHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateBlockResponse) Reset()         { *m = StateBlockResponse{} }
func (m *StateBlockResponse) String() string { return proto.CompactTextString(m) }
func (*StateBlockResponse) ProtoMessage()    {}
func (*StateBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{8}
}
func (m *StateBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateBlockResponse.Unmarshal(m, b)
}
func (m *StateBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateBlockResponse.Marshal(b, m, deterministic)
}
func (dst *StateBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateBlockResponse.Merge(dst, src)
}
func (m *StateBlockResponse) XXX_Size() int {
	return xxx_messageInfo_StateBlockResponse.Size(m)
}
func (m *StateBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateBlockResponse proto.InternalMessageInfo

func (m *StateBlockResponse) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *StateBlockResponse) GetBlockNumber() string {
	if m != nil {
		return m.BlockNumber
	}
	return ""
}

func (m *StateBlockResponse) GetBlockTime() string {
	if m != nil {
		return m.BlockTime
	}
	return ""
}

func (m *StateBlockResponse) GetImageHash() string {
	if m != nil {
		return m.ImageHash
	}
	return ""
}

func (m *StateBlockResponse) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *StateBlockResponse) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *StateBlockResponse) GetStatePrevDiffHash() string {
	if m != nil {
		return m.StatePrevDiffHash
	}
	return ""
}

func (m *StateBlockResponse) GetStateCurrentHash() string {
	if m != nil {
		return m.StateCurrentHash
	}
	return ""
}

type ImageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40d598b2c1d131a, []int{9}
}
func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (dst *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(dst, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*ErrorResponse)(nil), "protos.ErrorResponse")
	proto.RegisterType((*PingResponse)(nil), "protos.PingResponse")
	proto.RegisterType((*LatestBlockResponse)(nil), "protos.LatestBlockResponse")
	proto.RegisterType((*BlockResponse)(nil), "protos.BlockResponse")
	proto.RegisterType((*Signature)(nil), "protos.Signature")
	proto.RegisterType((*TransactionResponse)(nil), "protos.TransactionResponse")
	proto.RegisterType((*StateBlockResponse)(nil), "protos.StateBlockResponse")
	proto.RegisterType((*ImageResponse)(nil), "protos.ImageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// C3ServiceClient is the client API for C3Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type C3ServiceClient interface {
	Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type c3ServiceClient struct {
	cc *grpc.ClientConn
}

func NewC3ServiceClient(cc *grpc.ClientConn) C3ServiceClient {
	return &c3ServiceClient{cc}
}

func (c *c3ServiceClient) Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.C3Service/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// C3ServiceServer is the server API for C3Service service.
type C3ServiceServer interface {
	Send(context.Context, *Request) (*Response, error)
}

func RegisterC3ServiceServer(s *grpc.Server, srv C3ServiceServer) {
	s.RegisterService(&_C3Service_serviceDesc, srv)
}

func _C3Service_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(C3ServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.C3Service/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(C3ServiceServer).Send(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _C3Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.C3Service",
	HandlerType: (*C3ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _C3Service_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/pb/c3.proto",
}

func init() { proto.RegisterFile("rpc/pb/c3.proto", fileDescriptor_b40d598b2c1d131a) }

var fileDescriptor_b40d598b2c1d131a = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x5d, 0xd7, 0xae, 0x77, 0x2b, 0xdb, 0xbc, 0x31, 0x85, 0x09, 0xa1, 0xca, 0x20, 0x31,
	0x60, 0xb4, 0xd2, 0xc6, 0x03, 0x2f, 0x3c, 0x6c, 0x03, 0x09, 0x24, 0x40, 0x53, 0xba, 0x1f, 0x70,
	0x93, 0xdb, 0xcc, 0xac, 0xb1, 0x83, 0xed, 0x4c, 0xf4, 0x8f, 0xf8, 0x02, 0xbe, 0x8b, 0x4f, 0x40,
	0xb9, 0x71, 0xd3, 0x96, 0x0d, 0x04, 0x6f, 0x3c, 0xc5, 0xe7, 0xdc, 0xe3, 0xdc, 0xeb, 0x73, 0xaf,
	0x0d, 0x9b, 0x26, 0x8f, 0x07, 0xf9, 0x68, 0x10, 0x1f, 0xf7, 0x73, 0xa3, 0x9d, 0x66, 0x2d, 0xfa,
	0xd8, 0xfd, 0xfb, 0xa9, 0xd6, 0xe9, 0x04, 0x07, 0x04, 0x47, 0xc5, 0x78, 0x20, 0xd4, 0xb4, 0x92,
	0xf0, 0x18, 0xda, 0x11, 0x7e, 0x29, 0xd0, 0x3a, 0x16, 0x42, 0xfb, 0xb3, 0xd5, 0xca, 0xe4, 0x71,
	0x18, 0xf4, 0x82, 0x83, 0x4e, 0x34, 0x83, 0xec, 0x2e, 0x34, 0x64, 0x12, 0x36, 0x7a, 0xc1, 0x41,
	0x33, 0x6a, 0xc8, 0x84, 0xed, 0x41, 0x2b, 0x43, 0x77, 0xa9, 0x93, 0x70, 0x85, 0x84, 0x1e, 0x95,
	0x7c, 0x2e, 0x8c, 0xc8, 0x6c, 0xd8, 0xec, 0xad, 0x94, 0x7c, 0x85, 0xf8, 0x08, 0xd6, 0x22, 0xb4,
	0xb9, 0x56, 0x16, 0xff, 0x21, 0xcb, 0x21, 0xb4, 0x0c, 0xda, 0x62, 0xe2, 0x28, 0xcb, 0xfa, 0xd1,
	0x6e, 0xbf, 0x3a, 0x46, 0x7f, 0x76, 0x8c, 0xfe, 0x89, 0x9a, 0x46, 0x5e, 0xc3, 0x5f, 0x43, 0xf7,
	0xad, 0x31, 0xda, 0xd4, 0x89, 0x18, 0x34, 0x63, 0x9d, 0x20, 0x65, 0x69, 0x46, 0xb4, 0x2e, 0x93,
	0x67, 0x68, 0xad, 0x48, 0x91, 0xf2, 0x74, 0xa2, 0x19, 0xe4, 0x1c, 0x36, 0xce, 0xa5, 0x4a, 0x17,
	0x77, 0x27, 0xc2, 0x09, 0x5f, 0x23, 0xad, 0xf9, 0x53, 0xd8, 0xf9, 0x20, 0x1c, 0x5a, 0x77, 0x3a,
	0xd1, 0xf1, 0xd5, 0x1f, 0xa5, 0x3f, 0x1a, 0xd0, 0x5d, 0x56, 0x3d, 0x80, 0xce, 0xa8, 0x24, 0xde,
	0x09, 0x7b, 0xe9, 0xa5, 0x73, 0x82, 0xf5, 0x60, 0x9d, 0xc0, 0xa7, 0x22, 0x1b, 0xa1, 0xf1, 0xc5,
	0x2d, 0x52, 0xf5, 0xfe, 0x0b, 0x99, 0xa1, 0xb7, 0x7d, 0x4e, 0x94, 0x51, 0x99, 0x89, 0x14, 0xe9,
	0xef, 0xcd, 0x2a, 0x5a, 0x13, 0xec, 0x25, 0xdc, 0xb3, 0x4e, 0x38, 0xa4, 0x8a, 0xec, 0x47, 0x34,
	0x57, 0x93, 0x4a, 0xb9, 0x4a, 0xca, 0xdb, 0x83, 0xec, 0x31, 0x74, 0x73, 0x83, 0xd7, 0xa7, 0x75,
	0xd5, 0x2d, 0x52, 0x2f, 0x93, 0x6c, 0x17, 0x56, 0x95, 0x56, 0x31, 0x86, 0x6d, 0x8a, 0x56, 0x80,
	0x3d, 0x04, 0x48, 0xe4, 0x78, 0x2c, 0xe3, 0x62, 0xe2, 0xa6, 0xe1, 0x1a, 0x85, 0x16, 0x18, 0xc6,
	0x61, 0x23, 0x93, 0x0a, 0xcd, 0x49, 0x92, 0x18, 0xb4, 0x36, 0xec, 0x90, 0x62, 0x89, 0x63, 0x2f,
	0x60, 0x8d, 0xf0, 0x50, 0xa6, 0x21, 0xd0, 0x04, 0x6c, 0x57, 0xad, 0xb7, 0xfd, 0xa1, 0x4c, 0x95,
	0x70, 0x85, 0xc1, 0xa8, 0x96, 0xf0, 0x27, 0xd0, 0xa9, 0x69, 0xb6, 0x01, 0x81, 0xf1, 0x2e, 0x07,
	0xa6, 0x44, 0xd6, 0x7b, 0x1a, 0x58, 0xfe, 0x3d, 0x80, 0x9d, 0x0b, 0x23, 0x94, 0x15, 0xb1, 0x93,
	0x5a, 0xd5, 0x1d, 0xda, 0x83, 0x96, 0xfb, 0xba, 0xd0, 0x1e, 0x8f, 0x96, 0xbd, 0x6d, 0xfc, 0xea,
	0xed, 0xef, 0xee, 0x42, 0x08, 0xed, 0x5c, 0x4c, 0x27, 0x5a, 0x24, 0xfe, 0x32, 0xcc, 0x60, 0x39,
	0x2f, 0x63, 0xa3, 0x33, 0x6f, 0x3e, 0xad, 0xd9, 0x23, 0x58, 0xb1, 0x32, 0x25, 0x87, 0x6f, 0x3d,
	0x66, 0x19, 0xe5, 0xdf, 0x1a, 0xc0, 0x86, 0x75, 0xab, 0xfe, 0x8b, 0xc9, 0x9a, 0x7b, 0xb6, 0xba,
	0xe4, 0xd9, 0xdf, 0xcd, 0xce, 0x21, 0x6c, 0xd3, 0xe8, 0x9d, 0x1b, 0xbc, 0x7e, 0x23, 0xc7, 0x63,
	0x52, 0x56, 0x73, 0x74, 0x33, 0xc0, 0x9e, 0xc1, 0x16, 0x91, 0x67, 0x85, 0x31, 0xa8, 0x1c, 0x89,
	0xab, 0xc9, 0xba, 0xc1, 0xf3, 0x4d, 0xe8, 0xbe, 0x2f, 0x8b, 0x9c, 0x99, 0x74, 0xf4, 0x0a, 0x3a,
	0x67, 0xc7, 0x43, 0x34, 0xd7, 0x32, 0x46, 0xf6, 0x1c, 0x9a, 0x43, 0x54, 0x09, 0xdb, 0x9c, 0x19,
	0xed, 0x9f, 0xc0, 0xfd, 0xad, 0x39, 0x51, 0xed, 0xe3, 0x77, 0x46, 0xd5, 0x23, 0x7a, 0xfc, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x81, 0x1c, 0x84, 0x5e, 0x05, 0x00, 0x00,
}