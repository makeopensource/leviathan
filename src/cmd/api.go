package cmd

import (
	"connectrpc.com/connect"
	"fmt"
	dockerrpc "github.com/makeopensource/leviathan/generated/docker_rpc/v1/v1connect"
	jobrpc "github.com/makeopensource/leviathan/generated/jobs/v1/v1connect"
	labrpc "github.com/makeopensource/leviathan/generated/labs/v1/v1connect"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/docker"
	fm "github.com/makeopensource/leviathan/internal/file_manager"
	"github.com/makeopensource/leviathan/internal/jobs"
	"github.com/makeopensource/leviathan/internal/labs"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

// StartServerWithAddr starts a server using the port specified in the config
func StartServerWithAddr() {
	srvAddr := fmt.Sprintf(":%s", config.ServerPort.GetStr())
	StartServer(srvAddr)
}

func StartServer(srvAddr string) {
	mux := setupEndpoints()

	log.Info().Msg("Leviathan initialized successfully")
	log.Info().Msgf("starting server on %s", srvAddr)
	err := http.ListenAndServe(
		srvAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server on %s", srvAddr)
		return
	}
}

func setupEndpoints() *http.ServeMux {
	dk, job, lab := initServices()

	interceptor := connect.WithInterceptors()
	if config.ApiKey.GetStr() != "" {
		log.Info().Msg("ApiKey is set, endpoints now require authentication")
		interceptor = connect.WithInterceptors(&authInterceptor{config.ApiKey.GetStr()})
	}

	v1Endpoints := []func() (string, http.Handler){
		// jobs endpoints
		func() (string, http.Handler) {
			jobSrv := jobs.NewJobServer(job)
			return jobrpc.NewJobServiceHandler(jobSrv, interceptor)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &docker.Server{Service: dk}
			return dockerrpc.NewDockerServiceHandler(dkSrv, interceptor)
		},
		func() (string, http.Handler) {
			labSrv := labs.LabServer{Srv: lab}
			return labrpc.NewLabServiceHandler(labSrv, interceptor)
		},
		func() (string, http.Handler) {
			fileHandler := fm.NewFileManagerHandler("/files.v1")
			return fileHandler.BasePath + "/", fileHandler
		},
	}

	mux := http.NewServeMux()
	for _, svc := range v1Endpoints {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
