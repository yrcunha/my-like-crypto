package server

import (
	"context"
	"encoding/json"
	"log"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/proto/gen"
)

type Server struct {
	gen.UnimplementedScoreServiceServer
}

func (server *Server) CreateVote(ctx context.Context, vote *gen.CreateVoteReq) (*gen.CreateVoteRes, error) {
	body, marchalError := json.Marshal(vote.Vote)
	if marchalError != nil {
		log.Fatal(marchalError)
	}
	unmarshalVote := model.Votes{}
	unmarshalError := json.Unmarshal(body, &unmarshalVote)
	if unmarshalError != nil {
		log.Fatalf("Error in unmarshal method: %v", unmarshalError)
	}
	return &gen.CreateVoteRes{
		Success: true,
	}, nil
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
