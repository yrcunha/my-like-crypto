package server

import (
	"context"
	"log"

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
	result := &gen.CreateVoteRes{}
	unmarshalVote, unmarshalError := model.UnmarshalVote(vote)
	if unmarshalError != nil {
		log.Fatalf("Error in unmarshal method: %v", unmarshalError)
		result.Success = false
		return result, nil
	}
	repositorieError := repositorie.CreateVotes(server.Collection, ctx, unmarshalVote)
	if repositorieError != nil {
		log.Fatalf("Error in insert database: %v", repositorieError)
		result.Success = false
		return result, nil
	} else {
		result.Success = true
		return result, nil
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
