package routes

import (
	"context"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/stats"
)

type StatsServiceSrv struct {
	stats.UnimplementedStatsServiceServer
}

func (d *StatsServiceSrv) Echo(_ context.Context, request *stats.EchoRequest) (*stats.EchoResponse, error) {
	return &stats.EchoResponse{MessageResponse: request.GetMessage()}, nil
}
