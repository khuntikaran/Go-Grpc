package main

import (
	"context"
	"gogrpc/chat/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while conneting client %v", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "hii from client!",
	}
	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("error %v", err)
	}
	log.Printf("response from client:%s", response.Body)

}
