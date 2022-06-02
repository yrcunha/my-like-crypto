package model

import (
	"exemple.com/my-like-crypto-server/src/proto/gen"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Crypto struct {
	Name     string `json:"name"`
	Upvote   bool   `json:"upvote"`
	Downvote bool   `json:"downvote"`
}

type Votes struct {
	ID      string    `json:"id,omitempty" validate:"required,uuid"`
	Author  string    `json:"author" validate:"required,base64"`
	MyVotes [5]Crypto `json:"cryptos" validate:"required,len=5,dive,required"`
}

var validate *validator.Validate

func UnmarshalVote(vote *gen.CreateVoteReq) (*Votes, error) {
	unmarshalVote := &Votes{
		Author: vote.Vote.Author,
	}
	if vote.Vote.Id == "" {
		uuid, _ := uuid.NewUUID()
		unmarshalVote.ID = uuid.String()
	} else {
		unmarshalVote.ID = vote.Vote.Id
	}
	for key := range vote.Vote.Cryptos {
		if vote.Vote.Cryptos[key].Downvote == vote.Vote.Cryptos[key].Upvote {
			unmarshalVote.MyVotes[key].Downvote = false
			unmarshalVote.MyVotes[key].Upvote = false
		} else {
			unmarshalVote.MyVotes[key].Downvote = vote.Vote.Cryptos[key].Downvote
			unmarshalVote.MyVotes[key].Upvote = vote.Vote.Cryptos[key].Upvote
		}
		unmarshalVote.MyVotes[key].Name = vote.Vote.Cryptos[key].Name.String()
	}
	validate = validator.New()
	validationErrors := validate.StructExcept(unmarshalVote)
	if validationErrors != nil {
		return nil, validationErrors
	}
	return unmarshalVote, nil
}
