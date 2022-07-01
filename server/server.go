package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Zzocker/multicast/algo"
	pb "github.com/Zzocker/multicast/protos"
	"google.golang.org/grpc"
)

var lg = log.New(os.Stdout, "[MULTICAST-SERVER] ", 0)

func Run(port int, algoName algo.ALGO_NAME, peerCount int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	simulator := &algo.Simulator{}

	pb.RegisterMulticastServer(srv, &multicastService{simulator: simulator})
	simulator.Start(algoName, peerCount)

	if err = srv.Serve(lis); err != nil {
		panic(err)
	}
}

type multicastService struct {
	simulator *algo.Simulator
	pb.UnimplementedMulticastServer
}

func (s *multicastService) Set(ctx context.Context, in *pb.Data) (*pb.Empty, error) {
	s.simulator.Set(in.Value)
	return &pb.Empty{}, nil
}

func (s *multicastService) Get(ctx context.Context, in *pb.Empty) (*pb.Data, error) {
	value := s.simulator.Get()
	return &pb.Data{
		Value: value,
	}, nil
}
