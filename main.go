package main

import (
	"log"
	"net"
	"sync"

	"github.com/lantos1618/yumyum/proto/go/reactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct {
	reactions.
	connectedClients map[peer.Peer]
	mu               sync.Mutex
}

func (s *server) SendEmojiReaction(stream reactions.MyEmoji) error {
	for {
		emoji, err := stream.Recv()
		if err != nil {
			// handle error
			return err
		}
		s.mu.Lock()
		for _, client := range s.connectedClients {
			if err := client.Send(emoji); err != nil {
				// handle error
			}
		}
		s.mu.Unlock()
	}
}

func (s *server) ReceiveEmojiReaction(e *pb.Empty, stream pb.MyService_ReceiveEmojiReactionServer) error {
	p, _ := peer.FromContext(stream.Context())
	s.mu.Lock()
	s.connectedClients[*p] = stream
	s.mu.Unlock()
	<-stream.Context().Done()
	s.mu.Lock()
	delete(s.connectedClients, *p)
	s.mu.Unlock()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{
		connectedClients: make(map[peer.Peer]pb.MyService_ReceiveEmojiReactionServer),
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
