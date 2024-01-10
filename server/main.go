package main

import (
	"log"
	"net"

	pbproduct "github.com/suryatresna/samplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = ":50051"
)

func main() {
	hdl := &Handler{}
	startServer(hdl)
}

func startServer(hdl *Handler) {
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbproduct.RegisterProductServiceServer(s, hdl)

	reflection.Register(s)
	log.Printf("connect to GRPC server at %s", defaultPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
