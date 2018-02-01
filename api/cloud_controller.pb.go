// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloud_controller.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	cloud_controller.proto
	istio.proto
	common.proto

It has these top-level messages:
	ListCfRoutesRequest
	ListCfRoutesResponse
	Route
	UpsertRouteRequest
	UpsertRouteResponse
	DeleteRouteRequest
	DeleteRouteResponse
	CapiProcess
	RouteMapping
	MapRouteRequest
	MapRouteResponse
	UnmapRouteRequest
	UnmapRouteResponse
	RoutesRequest
	RoutesResponse
	BackendSet
	Backend
	HealthResponse
	HealthRequest
*/
package api

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

type ListCfRoutesRequest struct {
}

func (m *ListCfRoutesRequest) Reset()                    { *m = ListCfRoutesRequest{} }
func (m *ListCfRoutesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListCfRoutesRequest) ProtoMessage()               {}
func (*ListCfRoutesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListCfRoutesResponse struct {
	// key is route guid
	// value is host
	Routes map[string]string `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListCfRoutesResponse) Reset()                    { *m = ListCfRoutesResponse{} }
func (m *ListCfRoutesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCfRoutesResponse) ProtoMessage()               {}
func (*ListCfRoutesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListCfRoutesResponse) GetRoutes() map[string]string {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Guid string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Route) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Route) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type UpsertRouteRequest struct {
	Route *Route `protobuf:"bytes,1,opt,name=route" json:"route,omitempty"`
}

func (m *UpsertRouteRequest) Reset()                    { *m = UpsertRouteRequest{} }
func (m *UpsertRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertRouteRequest) ProtoMessage()               {}
func (*UpsertRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpsertRouteRequest) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type UpsertRouteResponse struct {
}

func (m *UpsertRouteResponse) Reset()                    { *m = UpsertRouteResponse{} }
func (m *UpsertRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertRouteResponse) ProtoMessage()               {}
func (*UpsertRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DeleteRouteRequest struct {
	Guid string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
}

func (m *DeleteRouteRequest) Reset()                    { *m = DeleteRouteRequest{} }
func (m *DeleteRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRouteRequest) ProtoMessage()               {}
func (*DeleteRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRouteRequest) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type DeleteRouteResponse struct {
}

func (m *DeleteRouteResponse) Reset()                    { *m = DeleteRouteResponse{} }
func (m *DeleteRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRouteResponse) ProtoMessage()               {}
func (*DeleteRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type CapiProcess struct {
	Guid             string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	DiegoProcessGuid string `protobuf:"bytes,2,opt,name=diego_process_guid,json=diegoProcessGuid" json:"diego_process_guid,omitempty"`
}

func (m *CapiProcess) Reset()                    { *m = CapiProcess{} }
func (m *CapiProcess) String() string            { return proto.CompactTextString(m) }
func (*CapiProcess) ProtoMessage()               {}
func (*CapiProcess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CapiProcess) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *CapiProcess) GetDiegoProcessGuid() string {
	if m != nil {
		return m.DiegoProcessGuid
	}
	return ""
}

type RouteMapping struct {
	CapiProcess *CapiProcess `protobuf:"bytes,1,opt,name=capi_process,json=capiProcess" json:"capi_process,omitempty"`
	RouteGuid   string       `protobuf:"bytes,2,opt,name=route_guid,json=routeGuid" json:"route_guid,omitempty"`
}

func (m *RouteMapping) Reset()                    { *m = RouteMapping{} }
func (m *RouteMapping) String() string            { return proto.CompactTextString(m) }
func (*RouteMapping) ProtoMessage()               {}
func (*RouteMapping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RouteMapping) GetCapiProcess() *CapiProcess {
	if m != nil {
		return m.CapiProcess
	}
	return nil
}

func (m *RouteMapping) GetRouteGuid() string {
	if m != nil {
		return m.RouteGuid
	}
	return ""
}

type MapRouteRequest struct {
	RouteMapping *RouteMapping `protobuf:"bytes,1,opt,name=route_mapping,json=routeMapping" json:"route_mapping,omitempty"`
}

func (m *MapRouteRequest) Reset()                    { *m = MapRouteRequest{} }
func (m *MapRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*MapRouteRequest) ProtoMessage()               {}
func (*MapRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MapRouteRequest) GetRouteMapping() *RouteMapping {
	if m != nil {
		return m.RouteMapping
	}
	return nil
}

type MapRouteResponse struct {
}

func (m *MapRouteResponse) Reset()                    { *m = MapRouteResponse{} }
func (m *MapRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*MapRouteResponse) ProtoMessage()               {}
func (*MapRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type UnmapRouteRequest struct {
	CapiProcessGuid string `protobuf:"bytes,1,opt,name=capi_process_guid,json=capiProcessGuid" json:"capi_process_guid,omitempty"`
	RouteGuid       string `protobuf:"bytes,2,opt,name=route_guid,json=routeGuid" json:"route_guid,omitempty"`
}

func (m *UnmapRouteRequest) Reset()                    { *m = UnmapRouteRequest{} }
func (m *UnmapRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*UnmapRouteRequest) ProtoMessage()               {}
func (*UnmapRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UnmapRouteRequest) GetCapiProcessGuid() string {
	if m != nil {
		return m.CapiProcessGuid
	}
	return ""
}

func (m *UnmapRouteRequest) GetRouteGuid() string {
	if m != nil {
		return m.RouteGuid
	}
	return ""
}

type UnmapRouteResponse struct {
}

func (m *UnmapRouteResponse) Reset()                    { *m = UnmapRouteResponse{} }
func (m *UnmapRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*UnmapRouteResponse) ProtoMessage()               {}
func (*UnmapRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*ListCfRoutesRequest)(nil), "api.ListCfRoutesRequest")
	proto.RegisterType((*ListCfRoutesResponse)(nil), "api.ListCfRoutesResponse")
	proto.RegisterType((*Route)(nil), "api.Route")
	proto.RegisterType((*UpsertRouteRequest)(nil), "api.UpsertRouteRequest")
	proto.RegisterType((*UpsertRouteResponse)(nil), "api.UpsertRouteResponse")
	proto.RegisterType((*DeleteRouteRequest)(nil), "api.DeleteRouteRequest")
	proto.RegisterType((*DeleteRouteResponse)(nil), "api.DeleteRouteResponse")
	proto.RegisterType((*CapiProcess)(nil), "api.CapiProcess")
	proto.RegisterType((*RouteMapping)(nil), "api.RouteMapping")
	proto.RegisterType((*MapRouteRequest)(nil), "api.MapRouteRequest")
	proto.RegisterType((*MapRouteResponse)(nil), "api.MapRouteResponse")
	proto.RegisterType((*UnmapRouteRequest)(nil), "api.UnmapRouteRequest")
	proto.RegisterType((*UnmapRouteResponse)(nil), "api.UnmapRouteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CloudControllerCopilot service

type CloudControllerCopilotClient interface {
	ListCfRoutes(ctx context.Context, in *ListCfRoutesRequest, opts ...grpc.CallOption) (*ListCfRoutesResponse, error)
	UpsertRoute(ctx context.Context, in *UpsertRouteRequest, opts ...grpc.CallOption) (*UpsertRouteResponse, error)
	DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error)
	MapRoute(ctx context.Context, in *MapRouteRequest, opts ...grpc.CallOption) (*MapRouteResponse, error)
	UnmapRoute(ctx context.Context, in *UnmapRouteRequest, opts ...grpc.CallOption) (*UnmapRouteResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type cloudControllerCopilotClient struct {
	cc *grpc.ClientConn
}

func NewCloudControllerCopilotClient(cc *grpc.ClientConn) CloudControllerCopilotClient {
	return &cloudControllerCopilotClient{cc}
}

func (c *cloudControllerCopilotClient) ListCfRoutes(ctx context.Context, in *ListCfRoutesRequest, opts ...grpc.CallOption) (*ListCfRoutesResponse, error) {
	out := new(ListCfRoutesResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/ListCfRoutes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) UpsertRoute(ctx context.Context, in *UpsertRouteRequest, opts ...grpc.CallOption) (*UpsertRouteResponse, error) {
	out := new(UpsertRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/UpsertRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error) {
	out := new(DeleteRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/DeleteRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) MapRoute(ctx context.Context, in *MapRouteRequest, opts ...grpc.CallOption) (*MapRouteResponse, error) {
	out := new(MapRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/MapRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) UnmapRoute(ctx context.Context, in *UnmapRouteRequest, opts ...grpc.CallOption) (*UnmapRouteResponse, error) {
	out := new(UnmapRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/UnmapRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/Health", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudControllerCopilot service

type CloudControllerCopilotServer interface {
	ListCfRoutes(context.Context, *ListCfRoutesRequest) (*ListCfRoutesResponse, error)
	UpsertRoute(context.Context, *UpsertRouteRequest) (*UpsertRouteResponse, error)
	DeleteRoute(context.Context, *DeleteRouteRequest) (*DeleteRouteResponse, error)
	MapRoute(context.Context, *MapRouteRequest) (*MapRouteResponse, error)
	UnmapRoute(context.Context, *UnmapRouteRequest) (*UnmapRouteResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

func RegisterCloudControllerCopilotServer(s *grpc.Server, srv CloudControllerCopilotServer) {
	s.RegisterService(&_CloudControllerCopilot_serviceDesc, srv)
}

func _CloudControllerCopilot_ListCfRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCfRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).ListCfRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/ListCfRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).ListCfRoutes(ctx, req.(*ListCfRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_UpsertRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).UpsertRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/UpsertRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).UpsertRoute(ctx, req.(*UpsertRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/DeleteRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).DeleteRoute(ctx, req.(*DeleteRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_MapRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).MapRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/MapRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).MapRoute(ctx, req.(*MapRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_UnmapRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmapRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).UnmapRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/UnmapRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).UnmapRoute(ctx, req.(*UnmapRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudControllerCopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CloudControllerCopilot",
	HandlerType: (*CloudControllerCopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCfRoutes",
			Handler:    _CloudControllerCopilot_ListCfRoutes_Handler,
		},
		{
			MethodName: "UpsertRoute",
			Handler:    _CloudControllerCopilot_UpsertRoute_Handler,
		},
		{
			MethodName: "DeleteRoute",
			Handler:    _CloudControllerCopilot_DeleteRoute_Handler,
		},
		{
			MethodName: "MapRoute",
			Handler:    _CloudControllerCopilot_MapRoute_Handler,
		},
		{
			MethodName: "UnmapRoute",
			Handler:    _CloudControllerCopilot_UnmapRoute_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _CloudControllerCopilot_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud_controller.proto",
}

func init() { proto.RegisterFile("cloud_controller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x8b, 0xd3, 0x4e,
	0x10, 0xbd, 0x5c, 0x7f, 0x2d, 0xbf, 0x9b, 0x44, 0xae, 0x9d, 0xeb, 0xf5, 0x62, 0x40, 0x28, 0x01,
	0xa1, 0x88, 0x54, 0x68, 0xe1, 0xf0, 0x04, 0x11, 0x8c, 0x87, 0x0a, 0x1e, 0x4a, 0xe0, 0xbe, 0x5a,
	0x72, 0xe9, 0xda, 0x0b, 0xa6, 0xd9, 0x35, 0xbb, 0x11, 0xee, 0xbf, 0xf0, 0xa3, 0x7f, 0xae, 0x64,
	0x76, 0xdb, 0x6e, 0xda, 0x88, 0xdf, 0x76, 0xdf, 0xce, 0x9b, 0xf7, 0xde, 0x74, 0x52, 0x18, 0xa5,
	0x39, 0xaf, 0x96, 0x8b, 0x94, 0x17, 0xaa, 0xe4, 0x79, 0xce, 0xca, 0xa9, 0x28, 0xb9, 0xe2, 0xd8,
	0x49, 0x44, 0x16, 0x78, 0x29, 0x5f, 0xaf, 0x79, 0xa1, 0xa1, 0xf0, 0x1c, 0xce, 0x3e, 0x65, 0x52,
	0x45, 0xdf, 0x62, 0x5e, 0x29, 0x26, 0x63, 0xf6, 0xa3, 0x62, 0x52, 0x85, 0xbf, 0x1c, 0x18, 0x36,
	0x71, 0x29, 0x78, 0x21, 0x19, 0xbe, 0x86, 0x5e, 0x49, 0x88, 0xef, 0x8c, 0x3b, 0x13, 0x77, 0xf6,
	0x74, 0x9a, 0x88, 0x6c, 0xda, 0x56, 0x3a, 0xd5, 0xd7, 0xeb, 0x42, 0x95, 0x0f, 0xb1, 0x21, 0x05,
	0x57, 0xe0, 0x5a, 0x30, 0xf6, 0xa1, 0xf3, 0x9d, 0x3d, 0xf8, 0xce, 0xd8, 0x99, 0x9c, 0xc4, 0xf5,
	0x11, 0x87, 0xd0, 0xfd, 0x99, 0xe4, 0x15, 0xf3, 0x8f, 0x09, 0xd3, 0x97, 0x57, 0xc7, 0x2f, 0x9d,
	0xf0, 0x05, 0x74, 0x89, 0x8a, 0x08, 0xff, 0xad, 0xaa, 0x6c, 0x69, 0x58, 0x74, 0xae, 0xb1, 0x7b,
	0x2e, 0x95, 0x61, 0xd1, 0x39, 0xbc, 0x04, 0xbc, 0x15, 0x92, 0x95, 0x8a, 0x68, 0x26, 0x19, 0x8e,
	0xa1, 0x4b, 0x5e, 0x88, 0xee, 0xce, 0x80, 0xfc, 0xeb, 0x0a, 0xfd, 0x50, 0x8f, 0xa4, 0xc1, 0xd3,
	0x71, 0xc2, 0x09, 0xe0, 0x3b, 0x96, 0x33, 0xc5, 0x1a, 0xed, 0x5a, 0xcc, 0xd4, 0x0d, 0x1a, 0x95,
	0xa6, 0xc1, 0x67, 0x70, 0xa3, 0x44, 0x64, 0x5f, 0x4a, 0x9e, 0x32, 0x29, 0x5b, 0x63, 0x3c, 0x07,
	0x5c, 0x66, 0x6c, 0xc5, 0x17, 0x42, 0x17, 0x2d, 0xa8, 0x42, 0x87, 0xea, 0xd3, 0x8b, 0x61, 0xbf,
	0xaf, 0x75, 0xee, 0xc0, 0x23, 0x85, 0x9b, 0x44, 0x88, 0xac, 0x58, 0xe1, 0x1c, 0xbc, 0x34, 0x11,
	0xd9, 0x86, 0x6c, 0x12, 0xf6, 0x29, 0xa1, 0xa5, 0x1c, 0xbb, 0xa9, 0x65, 0xe3, 0x09, 0x00, 0xc5,
	0xb6, 0xa5, 0x4e, 0x08, 0x21, 0x8d, 0x8f, 0x70, 0x7a, 0x93, 0x88, 0x46, 0xe4, 0x4b, 0x78, 0xa4,
	0x19, 0x6b, 0xad, 0x6b, 0x74, 0x06, 0xbb, 0x49, 0x1a, 0x43, 0xb1, 0x57, 0x5a, 0xb7, 0x10, 0xa1,
	0xbf, 0x6b, 0x65, 0x66, 0xf2, 0x15, 0x06, 0xb7, 0xc5, 0x7a, 0x4f, 0xe0, 0x19, 0x0c, 0xec, 0x1c,
	0x0b, 0x6b, 0x4c, 0xa7, 0x96, 0xf5, 0xda, 0xdf, 0xbf, 0xec, 0x0f, 0x01, 0xed, 0xfe, 0x5a, 0x75,
	0xf6, 0xbb, 0x03, 0xa3, 0xa8, 0xfe, 0x44, 0xa2, 0xed, 0x17, 0x12, 0x71, 0x91, 0xe5, 0x5c, 0xe1,
	0x35, 0x78, 0xf6, 0x32, 0xa3, 0xdf, 0xb2, 0xdf, 0xe4, 0x32, 0x78, 0xfc, 0xd7, 0xcd, 0x0f, 0x8f,
	0xf0, 0x2d, 0xb8, 0xd6, 0x0e, 0xe1, 0x05, 0xd5, 0x1e, 0x6e, 0x63, 0xe0, 0x1f, 0x3e, 0xd8, 0x3d,
	0xac, 0x35, 0x32, 0x3d, 0x0e, 0x57, 0xd0, 0xf4, 0x68, 0xdb, 0xb8, 0x23, 0xbc, 0x82, 0xff, 0x37,
	0x33, 0xc7, 0x21, 0xd5, 0xed, 0xfd, 0x9a, 0xc1, 0xf9, 0x1e, 0xba, 0xa5, 0xbe, 0x01, 0xd8, 0x8d,
	0x0e, 0x47, 0xda, 0xe8, 0xfe, 0x6f, 0x15, 0x5c, 0x1c, 0xe0, 0xdb, 0x06, 0x73, 0xe8, 0x7d, 0x60,
	0x49, 0xae, 0xee, 0x11, 0xa9, 0x48, 0x5f, 0x36, 0xc4, 0xb3, 0x06, 0xb6, 0x21, 0xdd, 0xf5, 0xe8,
	0x6f, 0x69, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x33, 0x07, 0xed, 0x17, 0xc3, 0x04, 0x00, 0x00,
}
