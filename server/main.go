package main

import (
	"crypto/tls"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"

	pb "github.com/lantos1618/yumyum/proto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedYumYumServiceServer
	connectedClients map[peer.Peer]pb.YumYumService_EmojiChatServer
	db               *gorm.DB
	addClient        chan pb.YumYumService_EmojiChatServer
	removeClient     chan pb.YumYumService_EmojiChatServer
	emojiBroadcast   chan *pb.Emoji
	mu               sync.Mutex
}

type EmojiReaction struct {
	gorm.Model
	ReactionType pb.EmojiReaction
	Sender       string
}

func (s *server) addNewClient(stream pb.YumYumService_EmojiChatServer) {
	p, _ := peer.FromContext(stream.Context())
	s.connectedClients[*p] = stream
}

func (s *server) removeClientFromList(stream pb.YumYumService_EmojiChatServer) {
	p, _ := peer.FromContext(stream.Context())
	delete(s.connectedClients, *p)
}

func (s *server) broadcastEmojiToClients(emoji *pb.Emoji) {
	for _, stream := range s.connectedClients {
		if err := stream.Send(emoji); err != nil {
			log.Errorf("Error sending emoji to client: %v", err)
			s.removeClientFromList(stream)
		}
	}
}

func (s *server) manageClients() {
	for {
		select {
		case stream := <-s.addClient:
			s.addNewClient(stream)
		case stream := <-s.removeClient:
			s.removeClientFromList(stream)
		case emoji := <-s.emojiBroadcast:
			s.broadcastEmojiToClients(emoji)
		}
	}
}

func (s *server) EmojiChat(stream pb.YumYumService_EmojiChatServer) error {
	p, _ := peer.FromContext(stream.Context())

	for {
		emoji, err := s.receiveMessage(stream)
		if err != nil {
			return err
		}

		err = s.saveEmojiToDatabase(emoji, p)
		if err != nil {
			log.Fatalf("Error saving emoji to database: ", err)
		}

		log.Infof("Received emoji: %v", emoji)

		err = s.forwardMessageToClients(emoji, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) receiveMessage(stream pb.YumYumService_EmojiChatServer) (*pb.Emoji, error) {
	emoji, err := stream.Recv()
	if err != nil {
		// handle error
		return nil, err
	}
	return emoji, nil
}

func (s *server) saveEmojiToDatabase(emoji *pb.Emoji, p *peer.Peer) error {
	emojiReaction := EmojiReaction{
		ReactionType: emoji.Reaction,
		Sender:       p.Addr.String(), // or another way to identify the sender
	}
	result := s.db.Create(&emojiReaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *server) forwardMessageToClients(emoji *pb.Emoji, p *peer.Peer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
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

	log.SetLevel(log.DebugLevel)
	// // load logging
	// verbose := flag.Bool("v", false, "Enable verbose logging")
	// flag.Parse()

	// if *verbose {
	// 	log.SetLevel(log.DebugLevel)
	// } else {
	// 	log.SetLevel(log.InfoLevel)
	// }

	// load database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// migrate schema
	db.AutoMigrate(&EmojiReaction{})

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

	srv := &server{
		connectedClients: make(map[peer.Peer]pb.YumYumService_EmojiChatServer),
		db:               db,
		addClient:        make(chan pb.YumYumService_EmojiChatServer),
		removeClient:     make(chan pb.YumYumService_EmojiChatServer),
		emojiBroadcast:   make(chan *pb.Emoji),
	}

	go srv.manageClients()

	s := grpc.NewServer(opts...)
	pb.RegisterYumYumServiceServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
