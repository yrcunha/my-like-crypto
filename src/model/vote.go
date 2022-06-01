package model

import (
	"exemple.com/my-like-crypto-server/src/proto/gen"
	"github.com/google/uuid"
)

type Crypto struct {
	Name     string `json:"name,omitempty"`
	Upvote   bool   `json:"upvote,omitempty"`
	Downvote bool   `json:"downvote,omitempty"`
}

type Votes struct {
	ID      string    `json:"id,omitempty"`
	Author  string    `json:"author,omitempty"`
	MyVotes [5]Crypto `json:"cryptos,omitempty"`
}

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
		unmarshalVote.MyVotes[key].Downvote = vote.Vote.Cryptos[key].Downvote
		unmarshalVote.MyVotes[key].Upvote = vote.Vote.Cryptos[key].Upvote
		unmarshalVote.MyVotes[key].Name = vote.Vote.Cryptos[key].Name.String()
	}
	return unmarshalVote, nil
}
