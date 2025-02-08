package v1

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/makeopensource/leviathan/generated/stats/v1"
	"github.com/makeopensource/leviathan/service/stats"
)

type StatsServer struct {
	Service *stats.StatService
}

func (stats *StatsServer) Echo(_ context.Context, req *connect.Request[v1.EchoRequest]) (*connect.Response[v1.EchoResponse], error) {
	res := connect.NewResponse(&v1.EchoResponse{MessageResponse: req.Msg.GetMessage()})
	return res, nil
}
