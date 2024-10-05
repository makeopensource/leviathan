package api

import (
	dkclient "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1/v1connect"
	"net/http"
)

func SetupPaths() *http.ServeMux {
	greeter := &DockerServer{}
	mux := http.NewServeMux()
	path, handler := dkclient.NewDockerServiceHandler(greeter)
	mux.Handle(path, handler)

	return mux
}
