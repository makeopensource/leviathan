package api

import (
	dkclient "github.com/makeopensource/leviathan/generated/docker_rpc/v1/v1connect"
	jobClient "github.com/makeopensource/leviathan/generated/jobs/v1/v1connect"
	labClient "github.com/makeopensource/leviathan/generated/labs/v1/v1connect"
	statsClient "github.com/makeopensource/leviathan/generated/stats/v1/v1connect"
	"github.com/makeopensource/leviathan/service"
	"net/http"
)

func SetupEndpoints() *http.ServeMux {
	docker, lab, job, stats := service.InitServices()

	endpoints := []func() (string, http.Handler){
		// jobs endpoints
		func() (string, http.Handler) {
			jobSrv := &JobServer{service: job}
			return jobClient.NewJobServiceHandler(jobSrv)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &DockerServer{service: docker}
			return dkclient.NewDockerServiceHandler(dkSrv)
		},
		// lab endpoints
		func() (string, http.Handler) {
			labSrv := &LabServer{service: lab}
			return labClient.NewLabServiceHandler(labSrv)
		},
		// stats endpoints
		func() (string, http.Handler) {
			statsSrv := &StatsServer{service: stats}
			return statsClient.NewStatsServiceHandler(statsSrv)
		},
	}

	mux := http.NewServeMux()
	for _, svc := range endpoints {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
