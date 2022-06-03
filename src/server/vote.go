package server

import (
	"context"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/proto/gen"
	"exemple.com/my-like-crypto-server/src/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	gen.UnimplementedScoreServiceServer
	Collection *mongo.Collection
}

func (server *Server) CreateVote(ctx context.Context, vote *gen.CreateVoteReq) (*gen.CreateVoteRes, error) {
	unmarshalVote, unmarshalError := model.UnmarshalVote(vote.Vote)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	repositorieError := repositorie.CreateVotes(server.Collection, ctx, unmarshalVote)
	if repositorieError != nil {
		return nil, repositorieError
	} else {
		return &gen.CreateVoteRes{
			Success: true,
		}, nil
	}
}

func (server *Server) UpdateVote(ctx context.Context, vote *gen.UpdateVoteReq) (*gen.UpdateVoteRes, error) {
	unmarshalVote, unmarshalError := model.UnmarshalVote(vote.Vote)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	repositorieError := repositorie.UpdateVotes(server.Collection, ctx, unmarshalVote)
	if repositorieError != nil {
		return nil, repositorieError
	} else {
		return &gen.UpdateVoteRes{
			Success: true,
		}, nil
	}
}

func (server *Server) DeleteVote(ctx context.Context, vote *gen.DeleteVoteReq) (*gen.DeleteVoteRes, error) {
	repositorieError := repositorie.DeleteVotes(server.Collection, ctx, vote.Id)
	if repositorieError != nil {
		return nil, repositorieError
	} else {
		return &gen.DeleteVoteRes{
			Success: true,
		}, nil
	}
}

func (server *Server) ListVotes(_ *gen.ListVotesReq, stream gen.ScoreService_ListVotesServer) error {
	// enviar 5 sends um com cada votação de crypto
	// melhorar mensagem de retorno, somente com a estatística já definida
	if err := stream.Send(&gen.ListVotesRes{
		Cryptos: []*gen.Crypto{},
	}); err != nil {
		return err
	}
	return nil
}

// preciso de mais um metodo para capturar o id do author
