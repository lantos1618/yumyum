package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/lantos1618/yumyum/proto/go"
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

func receiveMessages(client pb.YumYumServiceClient) {
	stream, err := client.EmojiChat(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ReceiveEmojiReaction: %v", err)
	}
	for {
		emoji, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive emoji: %v", err)
		}
		log.Printf("Received emoji: %v", emoji)
	}
}

func sendMessages(client pb.YumYumServiceClient) {
	stream, err := client.EmojiChat(context.Background(), grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("Failed to call SendEmojiReaction: %v", err)
	}
	defer stream.CloseSend()

	for {
		time.Sleep(time.Second * 1)
		if err := stream.Send(&pb.Emoji{Reaction: pb.EmojiReaction_LOVE}); err != nil {
			log.Fatalf("Failed to send emoji: %v", err)
		}
		log.Printf("Sent emoji: %v", pb.EmojiReaction_LOVE)
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

	go receiveMessages(client)
	sendMessages(client)
}
