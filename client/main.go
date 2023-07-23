package main

import (
	"context"
	"flag"

	"time"

	pb "github.com/lantos1618/yumyum/protos/go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	certFile := flag.String("cert", "", "The TLS cert file")
	keyFile := flag.String("key", "", "The TLS key file")
	flag.Parse()

	if *certFile == "" || *keyFile == "" {
		return nil, nil
	}

	return credentials.NewServerTLSFromFile(*certFile, *keyFile)
}

func createClientConnection(creds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	if creds != nil {
		return grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	}
	return grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func receiveMessages(stream pb.YumYumService_EmojiChatClient, quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			emoji, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive emoji: %v", err)
			}
			log.Printf("Received emoji: %v", emoji)
		}
	}
}

func sendMessages(stream pb.YumYumService_EmojiChatClient, quit <-chan bool) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-quit:
			ticker.Stop()
			return
		case <-ticker.C:
			if err := stream.Send(&pb.Emoji{Reaction: pb.EmojiReaction_LOVE}); err != nil {
				log.Fatalf("Failed to send emoji: %v", err)
			}
			log.Printf("Sent emoji: %v", pb.EmojiReaction_LOVE)
		}
	}
}

func main() {
	creds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	conn, err := createClientConnection(creds)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewYumYumServiceClient(conn)
	stream, err := client.EmojiChat(context.Background())
	if err != nil {
		log.Fatalf("Error on creating stream: %v", err)
	}
	quit := make(chan bool)

	go receiveMessages(stream, quit)
	go sendMessages(stream, quit)

	// Wait forever
	select {}

}
