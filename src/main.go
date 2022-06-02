package main

import (
	"context"
	"log"
	"net"
	"os"

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
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))
	listen, listenError := net.Listen("tcp", os.Getenv("PORT"))
	if listenError != nil {
		log.Fatalf("failed to listen: %v", listenError)
	}
	vote := server.Server{
		Collection: collection,
	}
	grpcServer := grpc.NewServer()
	gen.RegisterScoreServiceServer(grpcServer, &vote)
	log.Printf("Listening on Port %v!", os.Getenv("PORT"))
	if grpcError := grpcServer.Serve(listen); grpcError != nil {
		log.Fatalf("failed to serve: %s", grpcError)
	}
}
