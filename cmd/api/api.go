package api

import (
	"github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1/docker_rpcconnect"
	"net/http"
)

func SetupPaths() *http.ServeMux {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := docker_rpcconnect.NewDockerServiceHandler(greeter)
	mux.Handle(path, handler)

	return mux
}

//http.ListenAndServe(
//"localhost:8080",
//// Use h2c so we can serve HTTP/2 without TLS.
//h2c.NewHandler(mux, &http2.Server{}),
//)
