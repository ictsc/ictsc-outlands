// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: contestant/v1/recreation.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ictsc/ictsc-outlands/backend/internal/proto/contestant/v1"
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
	// RecreationServiceName is the fully-qualified name of the RecreationService service.
	RecreationServiceName = "contestant.v1.RecreationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RecreationServiceGetRecreationsProcedure is the fully-qualified name of the RecreationService's
	// GetRecreations RPC.
	RecreationServiceGetRecreationsProcedure = "/contestant.v1.RecreationService/GetRecreations"
	// RecreationServicePostRecreationProcedure is the fully-qualified name of the RecreationService's
	// PostRecreation RPC.
	RecreationServicePostRecreationProcedure = "/contestant.v1.RecreationService/PostRecreation"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	recreationServiceServiceDescriptor              = v1.File_contestant_v1_recreation_proto.Services().ByName("RecreationService")
	recreationServiceGetRecreationsMethodDescriptor = recreationServiceServiceDescriptor.Methods().ByName("GetRecreations")
	recreationServicePostRecreationMethodDescriptor = recreationServiceServiceDescriptor.Methods().ByName("PostRecreation")
)

// RecreationServiceClient is a client for the contestant.v1.RecreationService service.
type RecreationServiceClient interface {
	GetRecreations(context.Context, *connect.Request[v1.GetRecreationsRequest]) (*connect.Response[v1.GetRecreationsResponse], error)
	PostRecreation(context.Context, *connect.Request[v1.PostRecreationRequest]) (*connect.Response[v1.PostRecreationResponse], error)
}

// NewRecreationServiceClient constructs a client for the contestant.v1.RecreationService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRecreationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RecreationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &recreationServiceClient{
		getRecreations: connect.NewClient[v1.GetRecreationsRequest, v1.GetRecreationsResponse](
			httpClient,
			baseURL+RecreationServiceGetRecreationsProcedure,
			connect.WithSchema(recreationServiceGetRecreationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		postRecreation: connect.NewClient[v1.PostRecreationRequest, v1.PostRecreationResponse](
			httpClient,
			baseURL+RecreationServicePostRecreationProcedure,
			connect.WithSchema(recreationServicePostRecreationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// recreationServiceClient implements RecreationServiceClient.
type recreationServiceClient struct {
	getRecreations *connect.Client[v1.GetRecreationsRequest, v1.GetRecreationsResponse]
	postRecreation *connect.Client[v1.PostRecreationRequest, v1.PostRecreationResponse]
}

// GetRecreations calls contestant.v1.RecreationService.GetRecreations.
func (c *recreationServiceClient) GetRecreations(ctx context.Context, req *connect.Request[v1.GetRecreationsRequest]) (*connect.Response[v1.GetRecreationsResponse], error) {
	return c.getRecreations.CallUnary(ctx, req)
}

// PostRecreation calls contestant.v1.RecreationService.PostRecreation.
func (c *recreationServiceClient) PostRecreation(ctx context.Context, req *connect.Request[v1.PostRecreationRequest]) (*connect.Response[v1.PostRecreationResponse], error) {
	return c.postRecreation.CallUnary(ctx, req)
}

// RecreationServiceHandler is an implementation of the contestant.v1.RecreationService service.
type RecreationServiceHandler interface {
	GetRecreations(context.Context, *connect.Request[v1.GetRecreationsRequest]) (*connect.Response[v1.GetRecreationsResponse], error)
	PostRecreation(context.Context, *connect.Request[v1.PostRecreationRequest]) (*connect.Response[v1.PostRecreationResponse], error)
}

// NewRecreationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRecreationServiceHandler(svc RecreationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	recreationServiceGetRecreationsHandler := connect.NewUnaryHandler(
		RecreationServiceGetRecreationsProcedure,
		svc.GetRecreations,
		connect.WithSchema(recreationServiceGetRecreationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	recreationServicePostRecreationHandler := connect.NewUnaryHandler(
		RecreationServicePostRecreationProcedure,
		svc.PostRecreation,
		connect.WithSchema(recreationServicePostRecreationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/contestant.v1.RecreationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RecreationServiceGetRecreationsProcedure:
			recreationServiceGetRecreationsHandler.ServeHTTP(w, r)
		case RecreationServicePostRecreationProcedure:
			recreationServicePostRecreationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRecreationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRecreationServiceHandler struct{}

func (UnimplementedRecreationServiceHandler) GetRecreations(context.Context, *connect.Request[v1.GetRecreationsRequest]) (*connect.Response[v1.GetRecreationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("contestant.v1.RecreationService.GetRecreations is not implemented"))
}

func (UnimplementedRecreationServiceHandler) PostRecreation(context.Context, *connect.Request[v1.PostRecreationRequest]) (*connect.Response[v1.PostRecreationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("contestant.v1.RecreationService.PostRecreation is not implemented"))
}