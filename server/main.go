package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"sync"

	reactions "github.com/lantos1618/yumyum/proto/go/reactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type server struct {
	reactions.UnimplementedReactionsServiceServer
	connectedClients map[peer.Peer]reactions.ReactionsService_ReceiveEmojiReactionServer
	mu               sync.Mutex
}

func (s *server) SendEmojiReaction(stream reactions.ReactionsService_SendEmojiReactionServer) error {
	p, _ := peer.FromContext(stream.Context())
	for {
		emoji, err := stream.Recv()
		if err != nil {
			// handle error
			return err
		}
		s.mu.Lock()
		for client, clientStream := range s.connectedClients {
			if client != *p {
				if err := clientStream.Send(emoji); err != nil {
					// check to see if client is still connected
					// if not, remove from connectedClients
					// lets drop the client for now
					fmt.Println("Error sending emoji to client: ", err)
				}
			}
		}
		s.mu.Unlock()
	}
	return nil
}

func (s *server) ReceiveEmojiReaction(e *reactions.Empty, stream reactions.ReactionsService_ReceiveEmojiReactionServer) error {
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

func loadTLSCredentials(certFile string, keyFile string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return credentials.NewServerTLSFromCert(&cert), nil
}

func main() {
	certFile, keyFile := "server.crt", "server.key"
	tlsCreds, err := loadTLSCredentials(certFile, keyFile)
	if err != nil {
		log.Println("Launching server without TLS: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if tlsCreds != nil {
		opts = append(opts, grpc.Creds(tlsCreds))
	}

	s := grpc.NewServer(opts...)
	reactions.RegisterReactionsServiceServer(s, &server{
		connectedClients: make(map[peer.Peer]reactions.ReactionsService_ReceiveEmojiReactionServer),
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
