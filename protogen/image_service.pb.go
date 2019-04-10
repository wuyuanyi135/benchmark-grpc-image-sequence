// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/image_service.proto

package protogen

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ImageStreamRequest struct {
	IntervalMs           uint64   `protobuf:"varint,1,opt,name=interval_ms,json=intervalMs,proto3" json:"interval_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageStreamRequest) Reset()         { *m = ImageStreamRequest{} }
func (m *ImageStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ImageStreamRequest) ProtoMessage()    {}
func (*ImageStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df85201e21caee65, []int{0}
}

func (m *ImageStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageStreamRequest.Unmarshal(m, b)
}
func (m *ImageStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageStreamRequest.Marshal(b, m, deterministic)
}
func (m *ImageStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageStreamRequest.Merge(m, src)
}
func (m *ImageStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ImageStreamRequest.Size(m)
}
func (m *ImageStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageStreamRequest proto.InternalMessageInfo

func (m *ImageStreamRequest) GetIntervalMs() uint64 {
	if m != nil {
		return m.IntervalMs
	}
	return 0
}

type ImageStream struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageStream) Reset()         { *m = ImageStream{} }
func (m *ImageStream) String() string { return proto.CompactTextString(m) }
func (*ImageStream) ProtoMessage()    {}
func (*ImageStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_df85201e21caee65, []int{1}
}

func (m *ImageStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageStream.Unmarshal(m, b)
}
func (m *ImageStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageStream.Marshal(b, m, deterministic)
}
func (m *ImageStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageStream.Merge(m, src)
}
func (m *ImageStream) XXX_Size() int {
	return xxx_messageInfo_ImageStream.Size(m)
}
func (m *ImageStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageStream.DiscardUnknown(m)
}

var xxx_messageInfo_ImageStream proto.InternalMessageInfo

func (m *ImageStream) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageStream) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageStreamRequest)(nil), "image_service.ImageStreamRequest")
	proto.RegisterType((*ImageStream)(nil), "image_service.ImageStream")
}

func init() { proto.RegisterFile("proto/image_service.proto", fileDescriptor_df85201e21caee65) }

var fileDescriptor_df85201e21caee65 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x03, 0x8b, 0x09, 0xf1, 0xa2, 0x08, 0x2a, 0x99, 0x72, 0x09, 0x79, 0x82, 0x04, 0x82, 0x4b, 0x8a,
	0x52, 0x13, 0x73, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb9, 0xb8, 0x33, 0xf3,
	0x4a, 0x52, 0x8b, 0xca, 0x12, 0x73, 0xe2, 0x73, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82,
	0xb8, 0x60, 0x42, 0xbe, 0xc5, 0x4a, 0xe6, 0x5c, 0xdc, 0x48, 0xda, 0x84, 0x44, 0xb8, 0x58, 0xc1,
	0xc6, 0x82, 0x55, 0xf2, 0x04, 0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x51, 0x26, 0x97, 0x08, 0x44, 0x23, 0xc8, 0xa6,
	0xbc, 0xe4, 0xd4, 0x60, 0x88, 0x3b, 0x84, 0x02, 0xb9, 0xf8, 0x90, 0x0c, 0xcc, 0xcc, 0x4b, 0x17,
	0x52, 0xd4, 0x43, 0x75, 0x3e, 0xa6, 0x33, 0xa5, 0xa4, 0x70, 0x2b, 0x31, 0x60, 0x74, 0xb2, 0x8f,
	0xb2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x2f, 0xad, 0x2c,
	0x4d, 0xcc, 0xab, 0xcc, 0x34, 0x34, 0x36, 0xd5, 0x4f, 0x4a, 0xcd, 0x4b, 0xce, 0xc8, 0x4d, 0x2c,
	0xca, 0xd6, 0x4d, 0x2f, 0x2a, 0x48, 0xd6, 0x05, 0x1b, 0xa2, 0x5b, 0x0c, 0x75, 0x91, 0x3e, 0x38,
	0x9c, 0xd2, 0x53, 0xf3, 0x92, 0xd8, 0xc0, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x13, 0xf5, 0xc2, 0x4e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageSequenceServiceClient is the client API for ImageSequenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageSequenceServiceClient interface {
	ImageStreaming(ctx context.Context, in *ImageStreamRequest, opts ...grpc.CallOption) (ImageSequenceService_ImageStreamingClient, error)
}

type imageSequenceServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageSequenceServiceClient(cc *grpc.ClientConn) ImageSequenceServiceClient {
	return &imageSequenceServiceClient{cc}
}

func (c *imageSequenceServiceClient) ImageStreaming(ctx context.Context, in *ImageStreamRequest, opts ...grpc.CallOption) (ImageSequenceService_ImageStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageSequenceService_serviceDesc.Streams[0], "/image_service.ImageSequenceService/ImageStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageSequenceServiceImageStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageSequenceService_ImageStreamingClient interface {
	Recv() (*ImageStream, error)
	grpc.ClientStream
}

type imageSequenceServiceImageStreamingClient struct {
	grpc.ClientStream
}

func (x *imageSequenceServiceImageStreamingClient) Recv() (*ImageStream, error) {
	m := new(ImageStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageSequenceServiceServer is the server API for ImageSequenceService service.
type ImageSequenceServiceServer interface {
	ImageStreaming(*ImageStreamRequest, ImageSequenceService_ImageStreamingServer) error
}

func RegisterImageSequenceServiceServer(s *grpc.Server, srv ImageSequenceServiceServer) {
	s.RegisterService(&_ImageSequenceService_serviceDesc, srv)
}

func _ImageSequenceService_ImageStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageSequenceServiceServer).ImageStreaming(m, &imageSequenceServiceImageStreamingServer{stream})
}

type ImageSequenceService_ImageStreamingServer interface {
	Send(*ImageStream) error
	grpc.ServerStream
}

type imageSequenceServiceImageStreamingServer struct {
	grpc.ServerStream
}

func (x *imageSequenceServiceImageStreamingServer) Send(m *ImageStream) error {
	return x.ServerStream.SendMsg(m)
}

var _ImageSequenceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image_service.ImageSequenceService",
	HandlerType: (*ImageSequenceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImageStreaming",
			Handler:       _ImageSequenceService_ImageStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/image_service.proto",
}