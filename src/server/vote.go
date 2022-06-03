package server

import (
	"context"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/proto/gen"

	"exemple.com/my-like-crypto-server/src/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	gen.UnimplementedVotesServiceServer
	Collection *mongo.Collection
}

func (server *Server) Upvote(ctx context.Context, vote *gen.UpvoteReq) (*gen.UpvoteRes, error) {
	unmarshal := model.UnmarshalVote(vote.Name.String(), "upvote")
	err := repositorie.UpvoteOrDownvote(server.Collection, ctx, unmarshal, true)
	if err != nil {
		return nil, err
	} else {
		return &gen.UpvoteRes{
			Success: true,
		}, nil
	}
}

func (server *Server) Downvote(ctx context.Context, vote *gen.DownvoteReq) (*gen.DownvoteRes, error) {
	unmarshal := model.UnmarshalVote(vote.Name.String(), "downvote")
	err := repositorie.UpvoteOrDownvote(server.Collection, ctx, unmarshal, false)
	if err != nil {
		return nil, err
	} else {
		return &gen.DownvoteRes{
			Success: true,
		}, nil
	}
}

func (server *Server) CreateCrypto(ctx context.Context, vote *gen.CreateCryptoReq) (*gen.CreateCryptoRes, error) {
	unmarshal := model.UnmarshalCrypto(vote.Name.String())
	repositorieErr := repositorie.CreateCrypto(server.Collection, ctx, unmarshal)
	if repositorieErr != nil {
		return nil, repositorieErr
	} else {
		return &gen.CreateCryptoRes{
			Success: true,
		}, nil
	}
}

func (server *Server) DeleteCrypto(ctx context.Context, vote *gen.DeleteCryptoReq) (*gen.DeleteCryptoRes, error) {
	repositorieError := repositorie.DeleteCrypto(server.Collection, ctx, vote.Id)
	if repositorieError != nil {
		return nil, repositorieError
	} else {
		return &gen.DeleteCryptoRes{
			Success: true,
		}, nil
	}
}
