// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: admin/v1/rule.proto

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
	// RuleServiceName is the fully-qualified name of the RuleService service.
	RuleServiceName = "admin.v1.RuleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RuleServiceGetRuleProcedure is the fully-qualified name of the RuleService's GetRule RPC.
	RuleServiceGetRuleProcedure = "/admin.v1.RuleService/GetRule"
	// RuleServicePatchRuleProcedure is the fully-qualified name of the RuleService's PatchRule RPC.
	RuleServicePatchRuleProcedure = "/admin.v1.RuleService/PatchRule"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	ruleServiceServiceDescriptor         = v1.File_admin_v1_rule_proto.Services().ByName("RuleService")
	ruleServiceGetRuleMethodDescriptor   = ruleServiceServiceDescriptor.Methods().ByName("GetRule")
	ruleServicePatchRuleMethodDescriptor = ruleServiceServiceDescriptor.Methods().ByName("PatchRule")
)

// RuleServiceClient is a client for the admin.v1.RuleService service.
type RuleServiceClient interface {
	GetRule(context.Context, *connect.Request[v1.GetRuleRequest]) (*connect.Response[v1.GetRuleResponse], error)
	PatchRule(context.Context, *connect.Request[v1.PatchRuleRequest]) (*connect.Response[v1.PatchRuleResponse], error)
}

// NewRuleServiceClient constructs a client for the admin.v1.RuleService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRuleServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RuleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ruleServiceClient{
		getRule: connect.NewClient[v1.GetRuleRequest, v1.GetRuleResponse](
			httpClient,
			baseURL+RuleServiceGetRuleProcedure,
			connect.WithSchema(ruleServiceGetRuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		patchRule: connect.NewClient[v1.PatchRuleRequest, v1.PatchRuleResponse](
			httpClient,
			baseURL+RuleServicePatchRuleProcedure,
			connect.WithSchema(ruleServicePatchRuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// ruleServiceClient implements RuleServiceClient.
type ruleServiceClient struct {
	getRule   *connect.Client[v1.GetRuleRequest, v1.GetRuleResponse]
	patchRule *connect.Client[v1.PatchRuleRequest, v1.PatchRuleResponse]
}

// GetRule calls admin.v1.RuleService.GetRule.
func (c *ruleServiceClient) GetRule(ctx context.Context, req *connect.Request[v1.GetRuleRequest]) (*connect.Response[v1.GetRuleResponse], error) {
	return c.getRule.CallUnary(ctx, req)
}

// PatchRule calls admin.v1.RuleService.PatchRule.
func (c *ruleServiceClient) PatchRule(ctx context.Context, req *connect.Request[v1.PatchRuleRequest]) (*connect.Response[v1.PatchRuleResponse], error) {
	return c.patchRule.CallUnary(ctx, req)
}

// RuleServiceHandler is an implementation of the admin.v1.RuleService service.
type RuleServiceHandler interface {
	GetRule(context.Context, *connect.Request[v1.GetRuleRequest]) (*connect.Response[v1.GetRuleResponse], error)
	PatchRule(context.Context, *connect.Request[v1.PatchRuleRequest]) (*connect.Response[v1.PatchRuleResponse], error)
}

// NewRuleServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRuleServiceHandler(svc RuleServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	ruleServiceGetRuleHandler := connect.NewUnaryHandler(
		RuleServiceGetRuleProcedure,
		svc.GetRule,
		connect.WithSchema(ruleServiceGetRuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	ruleServicePatchRuleHandler := connect.NewUnaryHandler(
		RuleServicePatchRuleProcedure,
		svc.PatchRule,
		connect.WithSchema(ruleServicePatchRuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/admin.v1.RuleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RuleServiceGetRuleProcedure:
			ruleServiceGetRuleHandler.ServeHTTP(w, r)
		case RuleServicePatchRuleProcedure:
			ruleServicePatchRuleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRuleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRuleServiceHandler struct{}

func (UnimplementedRuleServiceHandler) GetRule(context.Context, *connect.Request[v1.GetRuleRequest]) (*connect.Response[v1.GetRuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.RuleService.GetRule is not implemented"))
}

func (UnimplementedRuleServiceHandler) PatchRule(context.Context, *connect.Request[v1.PatchRuleRequest]) (*connect.Response[v1.PatchRuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.RuleService.PatchRule is not implemented"))
}
