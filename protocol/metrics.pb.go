// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	metrics.proto

It has these top-level messages:
	MetricWriteRequest
	QueryMetricsRequest
	MetricWriteReply
	QueryMetricsReply
	ErrorStatus
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the user's name.
type MetricWriteRequest struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *MetricWriteRequest) Reset()                    { *m = MetricWriteRequest{} }
func (m *MetricWriteRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricWriteRequest) ProtoMessage()               {}
func (*MetricWriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricWriteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetricWriteRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The request message containing the user's name.
type QueryMetricsRequest struct {
	Start uint32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   uint32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *QueryMetricsRequest) Reset()                    { *m = QueryMetricsRequest{} }
func (m *QueryMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsRequest) ProtoMessage()               {}
func (*QueryMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryMetricsRequest) GetStart() uint32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *QueryMetricsRequest) GetEnd() uint32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *QueryMetricsRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The response message containing the greetings
type MetricWriteReply struct {
	Error *ErrorStatus `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *MetricWriteReply) Reset()                    { *m = MetricWriteReply{} }
func (m *MetricWriteReply) String() string            { return proto.CompactTextString(m) }
func (*MetricWriteReply) ProtoMessage()               {}
func (*MetricWriteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetricWriteReply) GetError() *ErrorStatus {
	if m != nil {
		return m.Error
	}
	return nil
}

// The response message containing the greetings
type QueryMetricsReply struct {
	Error   *ErrorStatus                `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Results []*QueryMetricsReply_Result `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
}

func (m *QueryMetricsReply) Reset()                    { *m = QueryMetricsReply{} }
func (m *QueryMetricsReply) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsReply) ProtoMessage()               {}
func (*QueryMetricsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryMetricsReply) GetError() *ErrorStatus {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *QueryMetricsReply) GetResults() []*QueryMetricsReply_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryMetricsReply_Result struct {
	Time  uint32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *QueryMetricsReply_Result) Reset()                    { *m = QueryMetricsReply_Result{} }
func (m *QueryMetricsReply_Result) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsReply_Result) ProtoMessage()               {}
func (*QueryMetricsReply_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *QueryMetricsReply_Result) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *QueryMetricsReply_Result) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// The response message containing the greetings
type ErrorStatus struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ErrorStatus) Reset()                    { *m = ErrorStatus{} }
func (m *ErrorStatus) String() string            { return proto.CompactTextString(m) }
func (*ErrorStatus) ProtoMessage()               {}
func (*ErrorStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ErrorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MetricWriteRequest)(nil), "protocol.MetricWriteRequest")
	proto.RegisterType((*QueryMetricsRequest)(nil), "protocol.QueryMetricsRequest")
	proto.RegisterType((*MetricWriteReply)(nil), "protocol.MetricWriteReply")
	proto.RegisterType((*QueryMetricsReply)(nil), "protocol.QueryMetricsReply")
	proto.RegisterType((*QueryMetricsReply_Result)(nil), "protocol.QueryMetricsReply.Result")
	proto.RegisterType((*ErrorStatus)(nil), "protocol.ErrorStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metrics service

type MetricsClient interface {
	WriteMetric(ctx context.Context, in *MetricWriteRequest, opts ...grpc.CallOption) (*MetricWriteReply, error)
	QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsReply, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) WriteMetric(ctx context.Context, in *MetricWriteRequest, opts ...grpc.CallOption) (*MetricWriteReply, error) {
	out := new(MetricWriteReply)
	err := grpc.Invoke(ctx, "/protocol.Metrics/WriteMetric", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsReply, error) {
	out := new(QueryMetricsReply)
	err := grpc.Invoke(ctx, "/protocol.Metrics/QueryMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metrics service

type MetricsServer interface {
	WriteMetric(context.Context, *MetricWriteRequest) (*MetricWriteReply, error)
	QueryMetrics(context.Context, *QueryMetricsRequest) (*QueryMetricsReply, error)
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_WriteMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).WriteMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Metrics/WriteMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).WriteMetric(ctx, req.(*MetricWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_QueryMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).QueryMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Metrics/QueryMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).QueryMetrics(ctx, req.(*QueryMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMetric",
			Handler:    _Metrics_WriteMetric_Handler,
		},
		{
			MethodName: "QueryMetrics",
			Handler:    _Metrics_QueryMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics.proto",
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4b, 0xbc, 0x40,
	0x14, 0xc7, 0x57, 0xfd, 0xb9, 0xfe, 0x7a, 0x66, 0x6c, 0xaf, 0x0d, 0xc4, 0x0a, 0x64, 0x2e, 0x09,
	0x81, 0x07, 0xbb, 0x45, 0xd0, 0xa9, 0x43, 0x50, 0x87, 0x66, 0x0f, 0x9d, 0x4d, 0x87, 0x10, 0x74,
	0xb5, 0x99, 0xf1, 0xe0, 0xbf, 0xd3, 0x9f, 0xd0, 0x5f, 0x18, 0x3b, 0xa3, 0xe8, 0xb2, 0x18, 0x74,
	0x72, 0xbe, 0xf3, 0xde, 0xfb, 0xf8, 0x9d, 0xef, 0x03, 0xaf, 0x62, 0x92, 0x17, 0x99, 0x88, 0x1b,
	0x5e, 0xcb, 0x1a, 0xff, 0xab, 0x4f, 0x56, 0x97, 0xe4, 0x0e, 0xf0, 0x45, 0x95, 0xde, 0x78, 0x21,
	0x19, 0x65, 0x9f, 0x2d, 0x13, 0x12, 0x4f, 0xc0, 0x2c, 0x72, 0xdf, 0x08, 0x8d, 0xe8, 0x88, 0x9a,
	0x45, 0x8e, 0x6b, 0xb0, 0xb3, 0xba, 0xdd, 0x4a, 0xdf, 0x0c, 0x8d, 0xc8, 0xa3, 0x5a, 0x90, 0x0d,
	0x9c, 0xbd, 0xb6, 0x8c, 0x77, 0x1a, 0x20, 0x86, 0xe1, 0x35, 0xd8, 0x42, 0xa6, 0x5c, 0xaa, 0x79,
	0x8f, 0x6a, 0x81, 0x2b, 0xb0, 0xd8, 0x36, 0xef, 0x01, 0xbb, 0xe3, 0x08, 0xb5, 0xa6, 0xd0, 0x07,
	0x58, 0xed, 0x19, 0x6a, 0xca, 0x0e, 0x6f, 0xc0, 0x66, 0x9c, 0xd7, 0x5c, 0x11, 0xdd, 0xe4, 0x3c,
	0x1e, 0xec, 0xc7, 0x8f, 0xbb, 0xeb, 0x8d, 0x4c, 0x65, 0x2b, 0xa8, 0xee, 0x21, 0xdf, 0x06, 0x9c,
	0xee, 0xdb, 0xfa, 0x2b, 0x02, 0xef, 0xc1, 0xe1, 0x4c, 0xb4, 0xa5, 0x14, 0xbe, 0x19, 0x5a, 0x91,
	0x9b, 0x90, 0xb1, 0xfd, 0x00, 0x1d, 0x53, 0xd5, 0x4a, 0x87, 0x91, 0x20, 0x81, 0xa5, 0xbe, 0x42,
	0x84, 0x7f, 0xb2, 0xa8, 0x58, 0x1f, 0x84, 0x3a, 0xcf, 0x44, 0x79, 0x0d, 0xee, 0xc4, 0x07, 0xfa,
	0xe0, 0x54, 0x4c, 0x88, 0xf4, 0x83, 0xf5, 0x4b, 0x18, 0x64, 0xf2, 0x65, 0x80, 0xd3, 0xff, 0x1d,
	0x9f, 0xc0, 0x55, 0x21, 0x69, 0x8d, 0x97, 0xa3, 0xc9, 0xc3, 0x95, 0x06, 0xc1, 0x4c, 0xb5, 0x29,
	0x3b, 0xb2, 0xc0, 0x67, 0x38, 0x9e, 0x3e, 0x0c, 0xaf, 0xe6, 0x1e, 0xac, 0x61, 0x17, 0xbf, 0xe4,
	0x41, 0x16, 0xef, 0x4b, 0x55, 0xbd, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xad, 0x64, 0xfa,
	0x76, 0x02, 0x00, 0x00,
}
