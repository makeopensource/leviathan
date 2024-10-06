// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: labs/v1/labs.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/makeopensource/leviathan/internal/generated/labs/v1"
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
	// LabServiceName is the fully-qualified name of the LabService service.
	LabServiceName = "labs.v1.LabService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LabServiceNewLabProcedure is the fully-qualified name of the LabService's NewLab RPC.
	LabServiceNewLabProcedure = "/labs.v1.LabService/NewLab"
	// LabServiceEditLabProcedure is the fully-qualified name of the LabService's EditLab RPC.
	LabServiceEditLabProcedure = "/labs.v1.LabService/EditLab"
	// LabServiceDeleteLabProcedure is the fully-qualified name of the LabService's DeleteLab RPC.
	LabServiceDeleteLabProcedure = "/labs.v1.LabService/DeleteLab"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	labServiceServiceDescriptor         = v1.File_labs_v1_labs_proto.Services().ByName("LabService")
	labServiceNewLabMethodDescriptor    = labServiceServiceDescriptor.Methods().ByName("NewLab")
	labServiceEditLabMethodDescriptor   = labServiceServiceDescriptor.Methods().ByName("EditLab")
	labServiceDeleteLabMethodDescriptor = labServiceServiceDescriptor.Methods().ByName("DeleteLab")
)

// LabServiceClient is a client for the labs.v1.LabService service.
type LabServiceClient interface {
	NewLab(context.Context, *connect.Request[v1.NewLabRequest]) (*connect.Response[v1.NewLabResponse], error)
	EditLab(context.Context, *connect.Request[v1.EditLabRequest]) (*connect.Response[v1.EditLabResponse], error)
	DeleteLab(context.Context, *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error)
}

// NewLabServiceClient constructs a client for the labs.v1.LabService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLabServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LabServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &labServiceClient{
		newLab: connect.NewClient[v1.NewLabRequest, v1.NewLabResponse](
			httpClient,
			baseURL+LabServiceNewLabProcedure,
			connect.WithSchema(labServiceNewLabMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		editLab: connect.NewClient[v1.EditLabRequest, v1.EditLabResponse](
			httpClient,
			baseURL+LabServiceEditLabProcedure,
			connect.WithSchema(labServiceEditLabMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteLab: connect.NewClient[v1.DeleteLabRequest, v1.DeleteLabResponse](
			httpClient,
			baseURL+LabServiceDeleteLabProcedure,
			connect.WithSchema(labServiceDeleteLabMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// labServiceClient implements LabServiceClient.
type labServiceClient struct {
	newLab    *connect.Client[v1.NewLabRequest, v1.NewLabResponse]
	editLab   *connect.Client[v1.EditLabRequest, v1.EditLabResponse]
	deleteLab *connect.Client[v1.DeleteLabRequest, v1.DeleteLabResponse]
}

// NewLab calls labs.v1.LabService.NewLab.
func (c *labServiceClient) NewLab(ctx context.Context, req *connect.Request[v1.NewLabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	return c.newLab.CallUnary(ctx, req)
}

// EditLab calls labs.v1.LabService.EditLab.
func (c *labServiceClient) EditLab(ctx context.Context, req *connect.Request[v1.EditLabRequest]) (*connect.Response[v1.EditLabResponse], error) {
	return c.editLab.CallUnary(ctx, req)
}

// DeleteLab calls labs.v1.LabService.DeleteLab.
func (c *labServiceClient) DeleteLab(ctx context.Context, req *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {
	return c.deleteLab.CallUnary(ctx, req)
}

// LabServiceHandler is an implementation of the labs.v1.LabService service.
type LabServiceHandler interface {
	NewLab(context.Context, *connect.Request[v1.NewLabRequest]) (*connect.Response[v1.NewLabResponse], error)
	EditLab(context.Context, *connect.Request[v1.EditLabRequest]) (*connect.Response[v1.EditLabResponse], error)
	DeleteLab(context.Context, *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error)
}

// NewLabServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLabServiceHandler(svc LabServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	labServiceNewLabHandler := connect.NewUnaryHandler(
		LabServiceNewLabProcedure,
		svc.NewLab,
		connect.WithSchema(labServiceNewLabMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	labServiceEditLabHandler := connect.NewUnaryHandler(
		LabServiceEditLabProcedure,
		svc.EditLab,
		connect.WithSchema(labServiceEditLabMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	labServiceDeleteLabHandler := connect.NewUnaryHandler(
		LabServiceDeleteLabProcedure,
		svc.DeleteLab,
		connect.WithSchema(labServiceDeleteLabMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/labs.v1.LabService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LabServiceNewLabProcedure:
			labServiceNewLabHandler.ServeHTTP(w, r)
		case LabServiceEditLabProcedure:
			labServiceEditLabHandler.ServeHTTP(w, r)
		case LabServiceDeleteLabProcedure:
			labServiceDeleteLabHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLabServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLabServiceHandler struct{}

func (UnimplementedLabServiceHandler) NewLab(context.Context, *connect.Request[v1.NewLabRequest]) (*connect.Response[v1.NewLabResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("labs.v1.LabService.NewLab is not implemented"))
}

func (UnimplementedLabServiceHandler) EditLab(context.Context, *connect.Request[v1.EditLabRequest]) (*connect.Response[v1.EditLabResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("labs.v1.LabService.EditLab is not implemented"))
}

func (UnimplementedLabServiceHandler) DeleteLab(context.Context, *connect.Request[v1.DeleteLabRequest]) (*connect.Response[v1.DeleteLabResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("labs.v1.LabService.DeleteLab is not implemented"))
}
