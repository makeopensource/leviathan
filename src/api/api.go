package api

import (
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

	log.Info().
		Str("build_date", common.BuildDate).
		Str("build_commit", common.CommitInfo).
		Str("git_branch", common.Branch).
		Str("go_version", common.GoVersion).
		Str("build_version", common.Version).
		Msg("Leviathan initialized successfully")

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

	endpoints := []func() (string, http.Handler){
		// jobs endpoints
		func() (string, http.Handler) {
			jobSrv := &v1.JobServer{Service: job}
			return jobClient.NewJobServiceHandler(jobSrv)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &v1.DockerServer{Service: docker}
			return dkclient.NewDockerServiceHandler(dkSrv)
		},
	}

	mux := http.NewServeMux()
	for _, svc := range endpoints {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
