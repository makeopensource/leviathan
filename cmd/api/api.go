package api

import (
	"github.com/docker/docker/client"
	dkclient "github.com/makeopensource/leviathan/internal/generated/docker_rpc/v1/v1connect"
	jobClient "github.com/makeopensource/leviathan/internal/generated/jobs/v1/v1connect"
	labClient "github.com/makeopensource/leviathan/internal/generated/labs/v1/v1connect"
	statsClient "github.com/makeopensource/leviathan/internal/generated/stats/v1/v1connect"
	"net/http"
)

func SetupPaths(clientList map[string]*client.Client) *http.ServeMux {
	mux := http.NewServeMux()

	dkSrv := &DockerServer{clientList: clientList}
	dkPath, dkHandler := dkclient.NewDockerServiceHandler(dkSrv)
	mux.Handle(dkPath, dkHandler)

	jobSrv := &JobServer{clientList: clientList}
	jobPath, jobHandler := jobClient.NewJobServiceHandler(jobSrv)
	mux.Handle(jobPath, jobHandler)

	labSrv := &LabServer{clientList: clientList}
	labPath, labHandler := labClient.NewLabServiceHandler(labSrv)
	mux.Handle(labPath, labHandler)

	statsSrv := &StatsServer{clientList: clientList}
	statsPath, statsHandler := statsClient.NewStatsServiceHandler(statsSrv)
	mux.Handle(statsPath, statsHandler)

	return mux
}
