package api

import (
	"connectrpc.com/connect"
	"fmt"
	v1 "github.com/makeopensource/leviathan/api/v1"
	"github.com/makeopensource/leviathan/common"
	dkclient "github.com/makeopensource/leviathan/generated/docker_rpc/v1/v1connect"
	jobClient "github.com/makeopensource/leviathan/generated/jobs/v1/v1connect"
	"github.com/makeopensource/leviathan/service"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func StartGrpcServer() {
	mux := setupEndpoints()

	log.Info().Msg("Leviathan initialized successfully")

	srvAddr := fmt.Sprintf(":%s", common.ServerPort.GetStr())
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
	docker, job := service.InitServices()

	interceptor := connect.WithInterceptors()
	if common.ApiKey.GetStr() != "" {
		log.Info().Msg("ApiKey is set, endpoints now require authentication")
		interceptor = connect.WithInterceptors(&authInterceptor{common.ApiKey.GetStr()})
	}

	endpoints := []func() (string, http.Handler){
		// jobs endpoints
		func() (string, http.Handler) {
			jobSrv := v1.NewJobServer(job)
			return jobClient.NewJobServiceHandler(jobSrv, interceptor)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &v1.DockerServer{Service: docker}
			return dkclient.NewDockerServiceHandler(dkSrv, interceptor)
		},
	}

	mux := http.NewServeMux()
	for _, svc := range endpoints {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
