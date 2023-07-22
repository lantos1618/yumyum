package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lantos1618/yumyum_backend/proto/yumyum_proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) React(ctx context.Context, req *pb.ReactionMessage) (*pb.ReactionResponse, error) {
	// your implementation here
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
