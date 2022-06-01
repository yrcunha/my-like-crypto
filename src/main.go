package main

import (
	"context"
	"log"
	"net"

	"exemple.com/my-like-crypto-server/src/proto/gen"
	"exemple.com/my-like-crypto-server/src/server"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	clientOptions := options.Client().ApplyURI("mongodb://docker:mongo@localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("my-like-crypto").Collection("vote")
	listen, listenError := net.Listen("tcp", ":8200")
	if listenError != nil {
		log.Fatalf("failed to listen: %v", listenError)
	}
	vote := server.Server{
		Collection: collection,
	}
	grpcServer := grpc.NewServer()
	gen.RegisterScoreServiceServer(grpcServer, &vote)
	log.Println("Listening on Port: 8200!")
	if grpcError := grpcServer.Serve(listen); grpcError != nil {
		log.Fatalf("failed to serve: %s", grpcError)
	}
}
