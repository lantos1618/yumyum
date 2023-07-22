package main

import (
	"context"
	"log"
	"net"

	// load our proto file from ./proto/go/reactions/reactions.proto
	pb "github.com/lantos1618/yumyum/proto/go/reactions"
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
