package main

import (
	"log"
	"net"

	"exemple.com/my-like-crypto-server/src/proto/gen"
	"exemple.com/my-like-crypto-server/src/server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	listen, listenError := net.Listen("tcp", ":8200")
	if listenError != nil {
		log.Fatalf("failed to listen: %v", listenError)
	}
	vote := server.Server{}
	grpcServer := grpc.NewServer()
	gen.RegisterScoreServiceServer(grpcServer, &vote)
	log.Println("Listening on Port: 8200!")
	if grpcError := grpcServer.Serve(listen); grpcError != nil {
		log.Fatalf("failed to serve: %s", grpcError)
	}
}
