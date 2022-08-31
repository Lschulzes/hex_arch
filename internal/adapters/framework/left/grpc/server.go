package grpc

import (
	"github.com/lschulzes/hex_arch/internal/adapters/framework/left/grpc/pb"
	"github.com/lschulzes/hex_arch/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpcA Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	arithServiceServer := grpcA
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithServiceServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve GRPC server over port 9000: %v", err)
	}
}
