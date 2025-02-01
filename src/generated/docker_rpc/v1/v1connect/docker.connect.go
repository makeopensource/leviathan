// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: docker_rpc/v1/docker.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/makeopensource/leviathan/generated/docker_rpc/v1"
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
	// DockerServiceName is the fully-qualified name of the DockerService service.
	DockerServiceName = "docker_rpc.v1.DockerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DockerServiceCreateContainerProcedure is the fully-qualified name of the DockerService's
	// CreateContainer RPC.
	DockerServiceCreateContainerProcedure = "/docker_rpc.v1.DockerService/CreateContainer"
	// DockerServiceDeleteContainerProcedure is the fully-qualified name of the DockerService's
	// DeleteContainer RPC.
	DockerServiceDeleteContainerProcedure = "/docker_rpc.v1.DockerService/DeleteContainer"
	// DockerServiceListContainersProcedure is the fully-qualified name of the DockerService's
	// ListContainers RPC.
	DockerServiceListContainersProcedure = "/docker_rpc.v1.DockerService/ListContainers"
	// DockerServiceStartContainerProcedure is the fully-qualified name of the DockerService's
	// StartContainer RPC.
	DockerServiceStartContainerProcedure = "/docker_rpc.v1.DockerService/StartContainer"
	// DockerServiceStopContainerProcedure is the fully-qualified name of the DockerService's
	// StopContainer RPC.
	DockerServiceStopContainerProcedure = "/docker_rpc.v1.DockerService/StopContainer"
	// DockerServiceGetContainerLogsProcedure is the fully-qualified name of the DockerService's
	// GetContainerLogs RPC.
	DockerServiceGetContainerLogsProcedure = "/docker_rpc.v1.DockerService/GetContainerLogs"
	// DockerServiceCreateNewImageProcedure is the fully-qualified name of the DockerService's
	// CreateNewImage RPC.
	DockerServiceCreateNewImageProcedure = "/docker_rpc.v1.DockerService/CreateNewImage"
	// DockerServiceListImagesProcedure is the fully-qualified name of the DockerService's ListImages
	// RPC.
	DockerServiceListImagesProcedure = "/docker_rpc.v1.DockerService/ListImages"
)

// DockerServiceClient is a client for the docker_rpc.v1.DockerService service.
type DockerServiceClient interface {
	CreateContainer(context.Context, *connect.Request[v1.CreateContainerRequest]) (*connect.Response[v1.CreateContainerResponse], error)
	DeleteContainer(context.Context, *connect.Request[v1.DeleteContainerRequest]) (*connect.Response[v1.DeleteContainerResponse], error)
	ListContainers(context.Context, *connect.Request[v1.ListContainersRequest]) (*connect.Response[v1.ListContainersResponse], error)
	StartContainer(context.Context, *connect.Request[v1.StartContainerRequest]) (*connect.Response[v1.StartContainerResponse], error)
	StopContainer(context.Context, *connect.Request[v1.StopContainerRequest]) (*connect.Response[v1.StopContainerResponse], error)
	GetContainerLogs(context.Context, *connect.Request[v1.GetContainerLogRequest]) (*connect.ServerStreamForClient[v1.GetContainerLogResponse], error)
	CreateNewImage(context.Context, *connect.Request[v1.NewImageRequest]) (*connect.Response[v1.NewImageResponse], error)
	ListImages(context.Context, *connect.Request[v1.ListImageRequest]) (*connect.Response[v1.ListImageResponse], error)
}

// NewDockerServiceClient constructs a client for the docker_rpc.v1.DockerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDockerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DockerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	dockerServiceMethods := v1.File_docker_rpc_v1_docker_proto.Services().ByName("DockerService").Methods()
	return &dockerServiceClient{
		createContainer: connect.NewClient[v1.CreateContainerRequest, v1.CreateContainerResponse](
			httpClient,
			baseURL+DockerServiceCreateContainerProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("CreateContainer")),
			connect.WithClientOptions(opts...),
		),
		deleteContainer: connect.NewClient[v1.DeleteContainerRequest, v1.DeleteContainerResponse](
			httpClient,
			baseURL+DockerServiceDeleteContainerProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("DeleteContainer")),
			connect.WithClientOptions(opts...),
		),
		listContainers: connect.NewClient[v1.ListContainersRequest, v1.ListContainersResponse](
			httpClient,
			baseURL+DockerServiceListContainersProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("ListContainers")),
			connect.WithClientOptions(opts...),
		),
		startContainer: connect.NewClient[v1.StartContainerRequest, v1.StartContainerResponse](
			httpClient,
			baseURL+DockerServiceStartContainerProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("StartContainer")),
			connect.WithClientOptions(opts...),
		),
		stopContainer: connect.NewClient[v1.StopContainerRequest, v1.StopContainerResponse](
			httpClient,
			baseURL+DockerServiceStopContainerProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("StopContainer")),
			connect.WithClientOptions(opts...),
		),
		getContainerLogs: connect.NewClient[v1.GetContainerLogRequest, v1.GetContainerLogResponse](
			httpClient,
			baseURL+DockerServiceGetContainerLogsProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("GetContainerLogs")),
			connect.WithClientOptions(opts...),
		),
		createNewImage: connect.NewClient[v1.NewImageRequest, v1.NewImageResponse](
			httpClient,
			baseURL+DockerServiceCreateNewImageProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("CreateNewImage")),
			connect.WithClientOptions(opts...),
		),
		listImages: connect.NewClient[v1.ListImageRequest, v1.ListImageResponse](
			httpClient,
			baseURL+DockerServiceListImagesProcedure,
			connect.WithSchema(dockerServiceMethods.ByName("ListImages")),
			connect.WithClientOptions(opts...),
		),
	}
}

// dockerServiceClient implements DockerServiceClient.
type dockerServiceClient struct {
	createContainer  *connect.Client[v1.CreateContainerRequest, v1.CreateContainerResponse]
	deleteContainer  *connect.Client[v1.DeleteContainerRequest, v1.DeleteContainerResponse]
	listContainers   *connect.Client[v1.ListContainersRequest, v1.ListContainersResponse]
	startContainer   *connect.Client[v1.StartContainerRequest, v1.StartContainerResponse]
	stopContainer    *connect.Client[v1.StopContainerRequest, v1.StopContainerResponse]
	getContainerLogs *connect.Client[v1.GetContainerLogRequest, v1.GetContainerLogResponse]
	createNewImage   *connect.Client[v1.NewImageRequest, v1.NewImageResponse]
	listImages       *connect.Client[v1.ListImageRequest, v1.ListImageResponse]
}

// CreateContainer calls docker_rpc.v1.DockerService.CreateContainer.
func (c *dockerServiceClient) CreateContainer(ctx context.Context, req *connect.Request[v1.CreateContainerRequest]) (*connect.Response[v1.CreateContainerResponse], error) {
	return c.createContainer.CallUnary(ctx, req)
}

// DeleteContainer calls docker_rpc.v1.DockerService.DeleteContainer.
func (c *dockerServiceClient) DeleteContainer(ctx context.Context, req *connect.Request[v1.DeleteContainerRequest]) (*connect.Response[v1.DeleteContainerResponse], error) {
	return c.deleteContainer.CallUnary(ctx, req)
}

// ListContainers calls docker_rpc.v1.DockerService.ListContainers.
func (c *dockerServiceClient) ListContainers(ctx context.Context, req *connect.Request[v1.ListContainersRequest]) (*connect.Response[v1.ListContainersResponse], error) {
	return c.listContainers.CallUnary(ctx, req)
}

// StartContainer calls docker_rpc.v1.DockerService.StartContainer.
func (c *dockerServiceClient) StartContainer(ctx context.Context, req *connect.Request[v1.StartContainerRequest]) (*connect.Response[v1.StartContainerResponse], error) {
	return c.startContainer.CallUnary(ctx, req)
}

// StopContainer calls docker_rpc.v1.DockerService.StopContainer.
func (c *dockerServiceClient) StopContainer(ctx context.Context, req *connect.Request[v1.StopContainerRequest]) (*connect.Response[v1.StopContainerResponse], error) {
	return c.stopContainer.CallUnary(ctx, req)
}

// GetContainerLogs calls docker_rpc.v1.DockerService.GetContainerLogs.
func (c *dockerServiceClient) GetContainerLogs(ctx context.Context, req *connect.Request[v1.GetContainerLogRequest]) (*connect.ServerStreamForClient[v1.GetContainerLogResponse], error) {
	return c.getContainerLogs.CallServerStream(ctx, req)
}

// CreateNewImage calls docker_rpc.v1.DockerService.CreateNewImage.
func (c *dockerServiceClient) CreateNewImage(ctx context.Context, req *connect.Request[v1.NewImageRequest]) (*connect.Response[v1.NewImageResponse], error) {
	return c.createNewImage.CallUnary(ctx, req)
}

// ListImages calls docker_rpc.v1.DockerService.ListImages.
func (c *dockerServiceClient) ListImages(ctx context.Context, req *connect.Request[v1.ListImageRequest]) (*connect.Response[v1.ListImageResponse], error) {
	return c.listImages.CallUnary(ctx, req)
}

// DockerServiceHandler is an implementation of the docker_rpc.v1.DockerService service.
type DockerServiceHandler interface {
	CreateContainer(context.Context, *connect.Request[v1.CreateContainerRequest]) (*connect.Response[v1.CreateContainerResponse], error)
	DeleteContainer(context.Context, *connect.Request[v1.DeleteContainerRequest]) (*connect.Response[v1.DeleteContainerResponse], error)
	ListContainers(context.Context, *connect.Request[v1.ListContainersRequest]) (*connect.Response[v1.ListContainersResponse], error)
	StartContainer(context.Context, *connect.Request[v1.StartContainerRequest]) (*connect.Response[v1.StartContainerResponse], error)
	StopContainer(context.Context, *connect.Request[v1.StopContainerRequest]) (*connect.Response[v1.StopContainerResponse], error)
	GetContainerLogs(context.Context, *connect.Request[v1.GetContainerLogRequest], *connect.ServerStream[v1.GetContainerLogResponse]) error
	CreateNewImage(context.Context, *connect.Request[v1.NewImageRequest]) (*connect.Response[v1.NewImageResponse], error)
	ListImages(context.Context, *connect.Request[v1.ListImageRequest]) (*connect.Response[v1.ListImageResponse], error)
}

// NewDockerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDockerServiceHandler(svc DockerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	dockerServiceMethods := v1.File_docker_rpc_v1_docker_proto.Services().ByName("DockerService").Methods()
	dockerServiceCreateContainerHandler := connect.NewUnaryHandler(
		DockerServiceCreateContainerProcedure,
		svc.CreateContainer,
		connect.WithSchema(dockerServiceMethods.ByName("CreateContainer")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceDeleteContainerHandler := connect.NewUnaryHandler(
		DockerServiceDeleteContainerProcedure,
		svc.DeleteContainer,
		connect.WithSchema(dockerServiceMethods.ByName("DeleteContainer")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceListContainersHandler := connect.NewUnaryHandler(
		DockerServiceListContainersProcedure,
		svc.ListContainers,
		connect.WithSchema(dockerServiceMethods.ByName("ListContainers")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceStartContainerHandler := connect.NewUnaryHandler(
		DockerServiceStartContainerProcedure,
		svc.StartContainer,
		connect.WithSchema(dockerServiceMethods.ByName("StartContainer")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceStopContainerHandler := connect.NewUnaryHandler(
		DockerServiceStopContainerProcedure,
		svc.StopContainer,
		connect.WithSchema(dockerServiceMethods.ByName("StopContainer")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceGetContainerLogsHandler := connect.NewServerStreamHandler(
		DockerServiceGetContainerLogsProcedure,
		svc.GetContainerLogs,
		connect.WithSchema(dockerServiceMethods.ByName("GetContainerLogs")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceCreateNewImageHandler := connect.NewUnaryHandler(
		DockerServiceCreateNewImageProcedure,
		svc.CreateNewImage,
		connect.WithSchema(dockerServiceMethods.ByName("CreateNewImage")),
		connect.WithHandlerOptions(opts...),
	)
	dockerServiceListImagesHandler := connect.NewUnaryHandler(
		DockerServiceListImagesProcedure,
		svc.ListImages,
		connect.WithSchema(dockerServiceMethods.ByName("ListImages")),
		connect.WithHandlerOptions(opts...),
	)
	return "/docker_rpc.v1.DockerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DockerServiceCreateContainerProcedure:
			dockerServiceCreateContainerHandler.ServeHTTP(w, r)
		case DockerServiceDeleteContainerProcedure:
			dockerServiceDeleteContainerHandler.ServeHTTP(w, r)
		case DockerServiceListContainersProcedure:
			dockerServiceListContainersHandler.ServeHTTP(w, r)
		case DockerServiceStartContainerProcedure:
			dockerServiceStartContainerHandler.ServeHTTP(w, r)
		case DockerServiceStopContainerProcedure:
			dockerServiceStopContainerHandler.ServeHTTP(w, r)
		case DockerServiceGetContainerLogsProcedure:
			dockerServiceGetContainerLogsHandler.ServeHTTP(w, r)
		case DockerServiceCreateNewImageProcedure:
			dockerServiceCreateNewImageHandler.ServeHTTP(w, r)
		case DockerServiceListImagesProcedure:
			dockerServiceListImagesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDockerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDockerServiceHandler struct{}

func (UnimplementedDockerServiceHandler) CreateContainer(context.Context, *connect.Request[v1.CreateContainerRequest]) (*connect.Response[v1.CreateContainerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.CreateContainer is not implemented"))
}

func (UnimplementedDockerServiceHandler) DeleteContainer(context.Context, *connect.Request[v1.DeleteContainerRequest]) (*connect.Response[v1.DeleteContainerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.DeleteContainer is not implemented"))
}

func (UnimplementedDockerServiceHandler) ListContainers(context.Context, *connect.Request[v1.ListContainersRequest]) (*connect.Response[v1.ListContainersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.ListContainers is not implemented"))
}

func (UnimplementedDockerServiceHandler) StartContainer(context.Context, *connect.Request[v1.StartContainerRequest]) (*connect.Response[v1.StartContainerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.StartContainer is not implemented"))
}

func (UnimplementedDockerServiceHandler) StopContainer(context.Context, *connect.Request[v1.StopContainerRequest]) (*connect.Response[v1.StopContainerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.StopContainer is not implemented"))
}

func (UnimplementedDockerServiceHandler) GetContainerLogs(context.Context, *connect.Request[v1.GetContainerLogRequest], *connect.ServerStream[v1.GetContainerLogResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.GetContainerLogs is not implemented"))
}

func (UnimplementedDockerServiceHandler) CreateNewImage(context.Context, *connect.Request[v1.NewImageRequest]) (*connect.Response[v1.NewImageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.CreateNewImage is not implemented"))
}

func (UnimplementedDockerServiceHandler) ListImages(context.Context, *connect.Request[v1.ListImageRequest]) (*connect.Response[v1.ListImageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("docker_rpc.v1.DockerService.ListImages is not implemented"))
}
