package model_test

import (
	"testing"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/proto/gen"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalVote(t *testing.T) {
	body := &gen.Vote{
		Id:     "",
		Author: "Yuri Rodrigues",
		Cryptos: []*gen.Crypto{
			{
				Name:     0,
				Upvote:   false,
				Downvote: true,
			}, {
				Name:     1,
				Upvote:   true,
				Downvote: true,
			},
		},
	}
	unmarshal, _ := model.UnmarshalVote(body)
	assert.Equal(t, unmarshal.Author, "Yuri Rodrigues")
	assert.Equal(t, unmarshal.MyVotes[0].Name, "BTC")
	assert.Equal(t, unmarshal.MyVotes[1].Name, "ETH")
	assert.Equal(t, unmarshal.MyVotes[1].Upvote, false)
	assert.Equal(t, unmarshal.MyVotes[1].Downvote, false)
}
