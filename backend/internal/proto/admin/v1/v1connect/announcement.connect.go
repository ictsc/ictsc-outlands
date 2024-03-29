// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: admin/v1/announcement.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ictsc/ictsc-outlands/backend/internal/proto/admin/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AnnouncementServiceName is the fully-qualified name of the AnnouncementService service.
	AnnouncementServiceName = "admin.v1.AnnouncementService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AnnouncementServiceGetAnnouncementsProcedure is the fully-qualified name of the
	// AnnouncementService's GetAnnouncements RPC.
	AnnouncementServiceGetAnnouncementsProcedure = "/admin.v1.AnnouncementService/GetAnnouncements"
	// AnnouncementServicePatchAnnouncementProcedure is the fully-qualified name of the
	// AnnouncementService's PatchAnnouncement RPC.
	AnnouncementServicePatchAnnouncementProcedure = "/admin.v1.AnnouncementService/PatchAnnouncement"
	// AnnouncementServicePostAnnouncementProcedure is the fully-qualified name of the
	// AnnouncementService's PostAnnouncement RPC.
	AnnouncementServicePostAnnouncementProcedure = "/admin.v1.AnnouncementService/PostAnnouncement"
	// AnnouncementServiceDeleteAnnouncementProcedure is the fully-qualified name of the
	// AnnouncementService's DeleteAnnouncement RPC.
	AnnouncementServiceDeleteAnnouncementProcedure = "/admin.v1.AnnouncementService/DeleteAnnouncement"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	announcementServiceServiceDescriptor                  = v1.File_admin_v1_announcement_proto.Services().ByName("AnnouncementService")
	announcementServiceGetAnnouncementsMethodDescriptor   = announcementServiceServiceDescriptor.Methods().ByName("GetAnnouncements")
	announcementServicePatchAnnouncementMethodDescriptor  = announcementServiceServiceDescriptor.Methods().ByName("PatchAnnouncement")
	announcementServicePostAnnouncementMethodDescriptor   = announcementServiceServiceDescriptor.Methods().ByName("PostAnnouncement")
	announcementServiceDeleteAnnouncementMethodDescriptor = announcementServiceServiceDescriptor.Methods().ByName("DeleteAnnouncement")
)

// AnnouncementServiceClient is a client for the admin.v1.AnnouncementService service.
type AnnouncementServiceClient interface {
	GetAnnouncements(context.Context, *connect.Request[v1.GetAnnouncementsRequest]) (*connect.Response[v1.GetAnnouncementsResponse], error)
	PatchAnnouncement(context.Context, *connect.Request[v1.PatchAnnouncementRequest]) (*connect.Response[v1.PatchAnnouncementResponse], error)
	PostAnnouncement(context.Context, *connect.Request[v1.PostAnnouncementRequest]) (*connect.Response[v1.PostAnnouncementResponse], error)
	DeleteAnnouncement(context.Context, *connect.Request[v1.DeleteAnnouncementRequest]) (*connect.Response[v1.DeleteAnnouncementResponse], error)
}

// NewAnnouncementServiceClient constructs a client for the admin.v1.AnnouncementService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAnnouncementServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AnnouncementServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &announcementServiceClient{
		getAnnouncements: connect.NewClient[v1.GetAnnouncementsRequest, v1.GetAnnouncementsResponse](
			httpClient,
			baseURL+AnnouncementServiceGetAnnouncementsProcedure,
			connect.WithSchema(announcementServiceGetAnnouncementsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		patchAnnouncement: connect.NewClient[v1.PatchAnnouncementRequest, v1.PatchAnnouncementResponse](
			httpClient,
			baseURL+AnnouncementServicePatchAnnouncementProcedure,
			connect.WithSchema(announcementServicePatchAnnouncementMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		postAnnouncement: connect.NewClient[v1.PostAnnouncementRequest, v1.PostAnnouncementResponse](
			httpClient,
			baseURL+AnnouncementServicePostAnnouncementProcedure,
			connect.WithSchema(announcementServicePostAnnouncementMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteAnnouncement: connect.NewClient[v1.DeleteAnnouncementRequest, v1.DeleteAnnouncementResponse](
			httpClient,
			baseURL+AnnouncementServiceDeleteAnnouncementProcedure,
			connect.WithSchema(announcementServiceDeleteAnnouncementMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// announcementServiceClient implements AnnouncementServiceClient.
type announcementServiceClient struct {
	getAnnouncements   *connect.Client[v1.GetAnnouncementsRequest, v1.GetAnnouncementsResponse]
	patchAnnouncement  *connect.Client[v1.PatchAnnouncementRequest, v1.PatchAnnouncementResponse]
	postAnnouncement   *connect.Client[v1.PostAnnouncementRequest, v1.PostAnnouncementResponse]
	deleteAnnouncement *connect.Client[v1.DeleteAnnouncementRequest, v1.DeleteAnnouncementResponse]
}

// GetAnnouncements calls admin.v1.AnnouncementService.GetAnnouncements.
func (c *announcementServiceClient) GetAnnouncements(ctx context.Context, req *connect.Request[v1.GetAnnouncementsRequest]) (*connect.Response[v1.GetAnnouncementsResponse], error) {
	return c.getAnnouncements.CallUnary(ctx, req)
}

// PatchAnnouncement calls admin.v1.AnnouncementService.PatchAnnouncement.
func (c *announcementServiceClient) PatchAnnouncement(ctx context.Context, req *connect.Request[v1.PatchAnnouncementRequest]) (*connect.Response[v1.PatchAnnouncementResponse], error) {
	return c.patchAnnouncement.CallUnary(ctx, req)
}

// PostAnnouncement calls admin.v1.AnnouncementService.PostAnnouncement.
func (c *announcementServiceClient) PostAnnouncement(ctx context.Context, req *connect.Request[v1.PostAnnouncementRequest]) (*connect.Response[v1.PostAnnouncementResponse], error) {
	return c.postAnnouncement.CallUnary(ctx, req)
}

// DeleteAnnouncement calls admin.v1.AnnouncementService.DeleteAnnouncement.
func (c *announcementServiceClient) DeleteAnnouncement(ctx context.Context, req *connect.Request[v1.DeleteAnnouncementRequest]) (*connect.Response[v1.DeleteAnnouncementResponse], error) {
	return c.deleteAnnouncement.CallUnary(ctx, req)
}

// AnnouncementServiceHandler is an implementation of the admin.v1.AnnouncementService service.
type AnnouncementServiceHandler interface {
	GetAnnouncements(context.Context, *connect.Request[v1.GetAnnouncementsRequest]) (*connect.Response[v1.GetAnnouncementsResponse], error)
	PatchAnnouncement(context.Context, *connect.Request[v1.PatchAnnouncementRequest]) (*connect.Response[v1.PatchAnnouncementResponse], error)
	PostAnnouncement(context.Context, *connect.Request[v1.PostAnnouncementRequest]) (*connect.Response[v1.PostAnnouncementResponse], error)
	DeleteAnnouncement(context.Context, *connect.Request[v1.DeleteAnnouncementRequest]) (*connect.Response[v1.DeleteAnnouncementResponse], error)
}

// NewAnnouncementServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAnnouncementServiceHandler(svc AnnouncementServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	announcementServiceGetAnnouncementsHandler := connect.NewUnaryHandler(
		AnnouncementServiceGetAnnouncementsProcedure,
		svc.GetAnnouncements,
		connect.WithSchema(announcementServiceGetAnnouncementsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	announcementServicePatchAnnouncementHandler := connect.NewUnaryHandler(
		AnnouncementServicePatchAnnouncementProcedure,
		svc.PatchAnnouncement,
		connect.WithSchema(announcementServicePatchAnnouncementMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	announcementServicePostAnnouncementHandler := connect.NewUnaryHandler(
		AnnouncementServicePostAnnouncementProcedure,
		svc.PostAnnouncement,
		connect.WithSchema(announcementServicePostAnnouncementMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	announcementServiceDeleteAnnouncementHandler := connect.NewUnaryHandler(
		AnnouncementServiceDeleteAnnouncementProcedure,
		svc.DeleteAnnouncement,
		connect.WithSchema(announcementServiceDeleteAnnouncementMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/admin.v1.AnnouncementService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AnnouncementServiceGetAnnouncementsProcedure:
			announcementServiceGetAnnouncementsHandler.ServeHTTP(w, r)
		case AnnouncementServicePatchAnnouncementProcedure:
			announcementServicePatchAnnouncementHandler.ServeHTTP(w, r)
		case AnnouncementServicePostAnnouncementProcedure:
			announcementServicePostAnnouncementHandler.ServeHTTP(w, r)
		case AnnouncementServiceDeleteAnnouncementProcedure:
			announcementServiceDeleteAnnouncementHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAnnouncementServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAnnouncementServiceHandler struct{}

func (UnimplementedAnnouncementServiceHandler) GetAnnouncements(context.Context, *connect.Request[v1.GetAnnouncementsRequest]) (*connect.Response[v1.GetAnnouncementsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.AnnouncementService.GetAnnouncements is not implemented"))
}

func (UnimplementedAnnouncementServiceHandler) PatchAnnouncement(context.Context, *connect.Request[v1.PatchAnnouncementRequest]) (*connect.Response[v1.PatchAnnouncementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.AnnouncementService.PatchAnnouncement is not implemented"))
}

func (UnimplementedAnnouncementServiceHandler) PostAnnouncement(context.Context, *connect.Request[v1.PostAnnouncementRequest]) (*connect.Response[v1.PostAnnouncementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.AnnouncementService.PostAnnouncement is not implemented"))
}

func (UnimplementedAnnouncementServiceHandler) DeleteAnnouncement(context.Context, *connect.Request[v1.DeleteAnnouncementRequest]) (*connect.Response[v1.DeleteAnnouncementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.AnnouncementService.DeleteAnnouncement is not implemented"))
}
