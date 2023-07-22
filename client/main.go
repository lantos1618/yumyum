package main

import (
	"context"
	"flag"
	"log"
	"time"

	reactions "github.com/lantos1618/yumyum/proto/go/reactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Command-line flags for the certificate and key files
	certFile := flag.String("cert", "", "The TLS cert file")
	keyFile := flag.String("key", "", "The TLS key file")
	flag.Parse()

	// If no credentials were provided, return nil
	if *certFile == "" || *keyFile == "" {
		return nil, nil
	}

	// Load the certificate and key files
	return credentials.NewServerTLSFromFile(*certFile, *keyFile)
}

func main() {
	creds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	// Create a connection to the server
	var conn *grpc.ClientConn
	if creds != nil {
		conn, err = grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	} else {
		conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	}
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := reactions.NewReactionsServiceClient(conn)

	// Start a goroutine to receive emoji reactions
	go func() {
		stream, err := client.ReceiveEmojiReaction(context.Background(), &reactions.Empty{})
		if err != nil {
			log.Fatalf("Failed to call ReceiveEmojiReaction: %v", err)
		}
		for {
			emoji, err := stream.Recv()
			if err != nil {
				// handle error
				log.Fatalf("Failed to receive emoji: %v", err)
			}
			log.Printf("Received emoji: %v", emoji)
		}
	}()

	// Send an emoji reaction
	stream, err := client.SendEmojiReaction(context.Background())
	if err != nil {
		log.Fatalf("Failed to call SendEmojiReaction: %v", err)
	}
	if err := stream.Send(&reactions.Emoji{Reaction: reactions.EmojiReaction_LOVE}); err != nil {
		log.Fatalf("Failed to send emoji: %v", err)
	}
	stream.CloseSend()

	// Let the goroutine run for a while before exiting
	time.Sleep(time.Second * 10)
}
