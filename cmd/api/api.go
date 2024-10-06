package api

import (
	"github.com/docker/docker/client"
	dkclient "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1/v1connect"
	"net/http"
)

func SetupPaths(clientList []*client.Client) *http.ServeMux {
	greeter := &DockerServer{clientList}
	mux := http.NewServeMux()
	path, handler := dkclient.NewDockerServiceHandler(greeter)
	mux.Handle(path, handler)

	return mux
}
