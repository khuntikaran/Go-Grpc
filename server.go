package main

import (
	"gogrpc/chat/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("error while listing on port 5050: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
