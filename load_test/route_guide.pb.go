// Code generated by protoc-gen-go.
// source: route_guide.proto
// DO NOT EDIT!

/*
Package routeguide is a generated protocol buffer package.

It is generated from these files:
	route_guide.proto

It has these top-level messages:
	Point
	Rectangle
	Feature
	FeatureDatabase
	RouteNote
	RouteSummary
*/
package load_test

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

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi *Point `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The point where the feature is detected.
	Location *Point `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

// Not used in the RPC.  Instead, this is here for the form serialized to disk.
type FeatureDatabase struct {
	Feature []*Feature `protobuf:"bytes,1,rep,name=feature" json:"feature,omitempty"`
}

func (m *FeatureDatabase) Reset()                    { *m = FeatureDatabase{} }
func (m *FeatureDatabase) String() string            { return proto.CompactTextString(m) }
func (*FeatureDatabase) ProtoMessage()               {}
func (*FeatureDatabase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FeatureDatabase) GetFeature() []*Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// The message to be sent.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RouteNote) Reset()                    { *m = RouteNote{} }
func (m *RouteNote) String() string            { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()               {}
func (*RouteNote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
}

func (m *RouteSummary) Reset()                    { *m = RouteSummary{} }
func (m *RouteSummary) String() string            { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()               {}
func (*RouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "routeguide.Feature")
	proto.RegisterType((*FeatureDatabase)(nil), "routeguide.FeatureDatabase")
	proto.RegisterType((*RouteNote)(nil), "routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routeguide.RouteSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteGuide service

type RouteGuideClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := grpc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[0], c.cc, "/routeguide.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[1], c.cc, "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[2], c.cc, "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route_guide.proto",
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x5d, 0xa7, 0xbb, 0x94, 0x4c, 0x82, 0x56, 0x6b, 0x84, 0x14, 0x15, 0x24, 0xc0, 0x5c, 0xf6,
	0x42, 0x54, 0x15, 0xa9, 0x47, 0x3e, 0x5a, 0x44, 0x2f, 0x15, 0x2a, 0xa1, 0xf7, 0xca, 0x4d, 0x4c,
	0x6a, 0x29, 0x89, 0xa3, 0xc4, 0x91, 0xe0, 0x07, 0xf0, 0x0b, 0xf8, 0x07, 0xfc, 0x52, 0x6c, 0xc7,
	0x69, 0x53, 0x48, 0xc5, 0x2d, 0x7e, 0xf3, 0xe6, 0xbd, 0xf1, 0xbc, 0x18, 0xee, 0x2a, 0xd1, 0x48,
	0xb6, 0x4b, 0x1b, 0x9e, 0xb0, 0xb0, 0xac, 0x84, 0x14, 0x18, 0x0c, 0x64, 0x10, 0xf2, 0x01, 0x6e,
	0x36, 0x82, 0x17, 0x12, 0x4f, 0xe0, 0x61, 0x46, 0x25, 0x97, 0x4d, 0xc2, 0x02, 0xf4, 0x02, 0xdd,
	0xdf, 0x44, 0xc7, 0x33, 0x7e, 0x06, 0x6e, 0x26, 0x8a, 0xb4, 0x2d, 0x3a, 0xa6, 0x78, 0x02, 0xc8,
	0x17, 0x70, 0x23, 0x16, 0x4b, 0x5a, 0xa4, 0x19, 0xc3, 0x2f, 0xc1, 0xc9, 0x84, 0x11, 0xf0, 0x66,
	0x77, 0xe1, 0xc9, 0x28, 0x34, 0x2e, 0x91, 0x2a, 0x6a, 0xca, 0x81, 0x1b, 0x99, 0x61, 0xca, 0x81,
	0x93, 0x35, 0x8c, 0x3f, 0x31, 0x2a, 0x9b, 0x8a, 0x61, 0x0c, 0xd7, 0x05, 0xcd, 0xdb, 0x99, 0xdc,
	0xc8, 0x7c, 0xe3, 0xd7, 0x6a, 0x56, 0x11, 0xab, 0xe9, 0x44, 0x71, 0x59, 0xe7, 0x48, 0x21, 0xef,
	0xe1, 0xd6, 0xaa, 0x7d, 0xa4, 0x92, 0xee, 0x69, 0xad, 0x15, 0xc6, 0xdf, 0x5a, 0x48, 0x09, 0x8f,
	0x94, 0xc0, 0xe3, 0xbe, 0x80, 0x65, 0x47, 0x1d, 0x87, 0x6c, 0xd5, 0x15, 0x75, 0xf9, 0xb3, 0x90,
	0xe7, 0xee, 0xe8, 0xbf, 0xee, 0x38, 0x80, 0x71, 0xce, 0xea, 0x9a, 0xa6, 0xed, 0xea, 0xdc, 0xa8,
	0x3b, 0x92, 0x5f, 0x08, 0x7c, 0x23, 0xfb, 0xb5, 0xc9, 0x73, 0x5a, 0xfd, 0xc0, 0xcf, 0xc1, 0x2b,
	0x75, 0xf7, 0x2e, 0x16, 0x4d, 0x21, 0x6d, 0x0c, 0x60, 0xa0, 0xa5, 0x46, 0xf0, 0x2b, 0x78, 0x64,
	0x47, 0xb2, 0x94, 0x36, 0x0c, 0xdf, 0x82, 0x2d, 0x49, 0x25, 0x99, 0xf0, 0x5a, 0xe5, 0x11, 0xb3,
	0x60, 0xd4, 0x26, 0xd9, 0x9d, 0xd5, 0xee, 0x7d, 0x96, 0xd1, 0xb2, 0x66, 0xc9, 0x4e, 0x72, 0xb5,
	0xd5, 0x6b, 0x53, 0xf7, 0x2c, 0xb6, 0x55, 0xd0, 0xec, 0xa7, 0x03, 0x60, 0xa6, 0x5a, 0xe9, 0xeb,
	0xe0, 0x39, 0xc0, 0x8a, 0xc9, 0x2e, 0x8d, 0x7f, 0x6f, 0x3a, 0x19, 0xda, 0x1c, 0xb9, 0xc2, 0x6f,
	0xc1, 0x5f, 0x2b, 0x57, 0x0b, 0xd4, 0xf8, 0x49, 0x9f, 0x76, 0xfc, 0x5f, 0x2e, 0x74, 0x4f, 0x91,
	0xea, 0xf7, 0x14, 0x4b, 0x54, 0x89, 0x99, 0x65, 0xc8, 0x38, 0x38, 0x53, 0xec, 0xed, 0x91, 0x5c,
	0xdd, 0x23, 0xfc, 0xce, 0x46, 0xb6, 0x3c, 0x50, 0xf9, 0x97, 0x79, 0x97, 0xe4, 0x64, 0x18, 0xd6,
	0xed, 0x53, 0xb4, 0x98, 0xc3, 0x53, 0x2e, 0xc2, 0xb4, 0x2a, 0xe3, 0x90, 0x7d, 0xa7, 0x79, 0x99,
	0xb1, 0xba, 0x47, 0x5f, 0xdc, 0x9e, 0x76, 0xb4, 0xd1, 0xaf, 0x6a, 0x83, 0x7e, 0x3b, 0xa3, 0x68,
	0xbb, 0xda, 0x3f, 0x30, 0x8f, 0xec, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x52, 0xa9,
	0x7a, 0x79, 0x03, 0x00, 0x00,
}