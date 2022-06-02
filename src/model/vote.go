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
	Author  string    `json:"author" validate:"required"`
	MyVotes [5]Crypto `json:"cryptos" validate:"required,len=5,dive,required"`
}

var validate *validator.Validate

func UnmarshalVote(vote *gen.Vote) (*Votes, error) {
	unmarshalVote := &Votes{
		Author: vote.Author,
	}
	if vote.Id == "" {
		uuid, _ := uuid.NewUUID()
		unmarshalVote.ID = uuid.String()
	} else {
		unmarshalVote.ID = vote.Id
	}
	for key := range vote.Cryptos {
		if vote.Cryptos[key].Downvote == vote.Cryptos[key].Upvote {
			unmarshalVote.MyVotes[key].Downvote = false
			unmarshalVote.MyVotes[key].Upvote = false
		} else {
			unmarshalVote.MyVotes[key].Downvote = vote.Cryptos[key].Downvote
			unmarshalVote.MyVotes[key].Upvote = vote.Cryptos[key].Upvote
		}
		unmarshalVote.MyVotes[key].Name = vote.Cryptos[key].Name.String()
	}
	validate = validator.New()
	validationErrors := validate.StructExcept(unmarshalVote)
	if validationErrors != nil {
		return nil, validationErrors
	}
	return unmarshalVote, nil
}
