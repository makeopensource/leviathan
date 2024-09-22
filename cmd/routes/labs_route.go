package routes

import (
	"context"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/labs"
)

// grpc implementation for lab service

type LabServiceSrv struct {
	labs.UnimplementedLabServiceServer
}

func (d *LabServiceSrv) NewLab(_ context.Context, request *labs.NewLabRequest) (*labs.NewLabResponse, error) {
	return &labs.NewLabResponse{}, nil
}

func (d *LabServiceSrv) EditLab(_ context.Context, request *labs.EditLabRequest) (*labs.EditLabResponse, error) {
	return &labs.EditLabResponse{}, nil
}

func (d *LabServiceSrv) DeleteLab(_ context.Context, request *labs.DeleteLabRequest) (*labs.DeleteLabResponse, error) {
	return &labs.DeleteLabResponse{}, nil
}

func (d *LabServiceSrv) Echo(_ context.Context, request *labs.EchoRequest) (*labs.EchoResponse, error) {
	return &labs.EchoResponse{MessageResponse: request.GetMessage()}, nil
}
