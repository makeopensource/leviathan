package api

import (
	"github.com/docker/docker/client"
	dkclient "github.com/makeopensource/leviathan/generated/docker_rpc/v1/v1connect"
	jobClient "github.com/makeopensource/leviathan/generated/jobs/v1/v1connect"
	labClient "github.com/makeopensource/leviathan/generated/labs/v1/v1connect"
	statsClient "github.com/makeopensource/leviathan/generated/stats/v1/v1connect"
	"net/http"
)

func SetupPaths(clientList map[string]*client.Client) *http.ServeMux {
	mux := http.NewServeMux()

	services := []func() (string, http.Handler){
		// jobs endpoints
		func() (string, http.Handler) {
			jobSrv := &JobServer{clientList: clientList}
			return jobClient.NewJobServiceHandler(jobSrv)
		},
		// docker endpoints
		func() (string, http.Handler) {
			dkSrv := &DockerServer{clientList: clientList}
			return dkclient.NewDockerServiceHandler(dkSrv)
		},
		// lab endpoints
		func() (string, http.Handler) {
			labSrv := &LabServer{clientList: clientList}
			return labClient.NewLabServiceHandler(labSrv)
		},
		// stats endpoints
		func() (string, http.Handler) {
			statsSrv := &StatsServer{clientList: clientList}
			return statsClient.NewStatsServiceHandler(statsSrv)
		},
	}

	for _, svc := range services {
		path, handler := svc()
		mux.Handle(path, handler)
	}

	return mux
}
