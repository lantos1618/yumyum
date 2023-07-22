package main

import (
	"crypto/tls"
	"flag"

	"net"
	"sync"

	log "github.com/sirupsen/logrus"

	reactions "github.com/lantos1618/yumyum/proto/go/reactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type server struct {
	reactions.UnimplementedReactionsServiceServer
	connectedClients map[peer.Peer]reactions.ReactionsService_ReceiveEmojiReactionServer
	db               *gorm.DB
	mu               sync.Mutex
}

// models
type EmojiReaction struct {
	gorm.Model
	ReactionType reactions.EmojiReaction
	Sender       string
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
					log.Errorf("Error sending emoji to client: %v", err)
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
	for {
		emoji := &reactions.Emoji{}
		err := stream.RecvMsg(emoji)
		if err != nil {
			// handle error
			break
		}

		// Save the received emoji into the database
		emojiReaction := EmojiReaction{
			ReactionType: emoji.Reaction,
			Sender:       p.Addr.String(), // or another way to identify the sender
		}
		result := s.db.Create(&emojiReaction)
		if result.Error != nil {
			log.Fatalf("Error saving emoji to database: ", result.Error)
		}
	}
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

	// load logging
	verbose := flag.Bool("v", false, "Enable verbose logging")
	flag.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// load database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// load certs
	certFile, keyFile := "server.crt", "server.key"
	tlsCreds, err := loadTLSCredentials(certFile, keyFile)
	if err != nil {
		log.Infof("Launching server without TLS: %v", err)
	}

	// start server
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
		db:               db,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
