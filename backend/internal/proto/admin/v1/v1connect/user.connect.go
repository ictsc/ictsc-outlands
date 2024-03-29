// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: admin/v1/user.proto

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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "admin.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/admin.v1.UserService/GetUser"
	// UserServiceGetUsersProcedure is the fully-qualified name of the UserService's GetUsers RPC.
	UserServiceGetUsersProcedure = "/admin.v1.UserService/GetUsers"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/admin.v1.UserService/DeleteUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor          = v1.File_admin_v1_user_proto.Services().ByName("UserService")
	userServiceGetUserMethodDescriptor    = userServiceServiceDescriptor.Methods().ByName("GetUser")
	userServiceGetUsersMethodDescriptor   = userServiceServiceDescriptor.Methods().ByName("GetUsers")
	userServiceDeleteUserMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("DeleteUser")
)

// UserServiceClient is a client for the admin.v1.UserService service.
type UserServiceClient interface {
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	GetUsers(context.Context, *connect.Request[v1.GetUsersRequest]) (*connect.Response[v1.GetUsersResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceClient constructs a client for the admin.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUser: connect.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect.WithSchema(userServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUsers: connect.NewClient[v1.GetUsersRequest, v1.GetUsersResponse](
			httpClient,
			baseURL+UserServiceGetUsersProcedure,
			connect.WithSchema(userServiceGetUsersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			connect.WithSchema(userServiceDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUser    *connect.Client[v1.GetUserRequest, v1.GetUserResponse]
	getUsers   *connect.Client[v1.GetUsersRequest, v1.GetUsersResponse]
	deleteUser *connect.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
}

// GetUser calls admin.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// GetUsers calls admin.v1.UserService.GetUsers.
func (c *userServiceClient) GetUsers(ctx context.Context, req *connect.Request[v1.GetUsersRequest]) (*connect.Response[v1.GetUsersResponse], error) {
	return c.getUsers.CallUnary(ctx, req)
}

// DeleteUser calls admin.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the admin.v1.UserService service.
type UserServiceHandler interface {
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	GetUsers(context.Context, *connect.Request[v1.GetUsersRequest]) (*connect.Response[v1.GetUsersResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceGetUserHandler := connect.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetUsersHandler := connect.NewUnaryHandler(
		UserServiceGetUsersProcedure,
		svc.GetUsers,
		connect.WithSchema(userServiceGetUsersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserHandler := connect.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(userServiceDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/admin.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceGetUsersProcedure:
			userServiceGetUsersHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUsers(context.Context, *connect.Request[v1.GetUsersRequest]) (*connect.Response[v1.GetUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.UserService.GetUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("admin.v1.UserService.DeleteUser is not implemented"))
}
