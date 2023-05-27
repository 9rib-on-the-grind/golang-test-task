package golang_test_task

import (
	"context"
	"golang-test-task/rpc"
)

type SolutionServer struct {
	rpc.UnimplementedSolutionServer
}

func NewSolutionServer() *SolutionServer {
	return &SolutionServer{}
}

func (s *SolutionServer) MineBlock(ctx context.Context, request *rpc.MineBlockRequest) (*rpc.MineBlockResponse, error) {

	return &rpc.MineBlockResponse{}, nil
}

func (s *SolutionServer) CountCycles(ctx context.Context, request *rpc.CountCyclesRequest) (*rpc.CountCyclesResponse, error) {

	return &rpc.CountCyclesResponse{}, nil
}
