package api

import (
	v1 "github.com/makeopensource/leviathan/api/v1"
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
			jobSrv := &v1.JobServer{Service: job}
			return jobClient.NewJobServiceHandler(jobSrv)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &v1.DockerServer{Service: docker}
			return dkclient.NewDockerServiceHandler(dkSrv)
		},
		// lab endpoints
		func() (string, http.Handler) {
			labSrv := &v1.LabServer{Service: lab}
			return labClient.NewLabServiceHandler(labSrv)
		},
		// stats endpoints
		func() (string, http.Handler) {
			statsSrv := &v1.StatsServer{Service: stats}
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
