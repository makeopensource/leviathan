package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1"
)

type DockerServer struct{}

func CreateContainer(context.Context, *connect.Request[docker_rpc.CreateContainerRequest]) (*connect.Response[docker_rpc.CreateContainerResponse], error)
func DeleteContainer(context.Context, *connect.Request[docker_rpc.DeleteContainerRequest]) (*connect.Response[docker_rpc.DeleteContainerResponse], error)
func ListContainers(context.Context, *connect.Request[docker_rpc.ListContainersRequest]) (*connect.Response[docker_rpc.ListContainersResponse], error)
func Echo(context.Context, *connect.Request[docker_rpc.EchoRequest]) (*connect.Response[docker_rpc.EchoResponse], error)
