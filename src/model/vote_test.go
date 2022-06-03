package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yrcunha/my-like-crypto-server/src/model"
	"github.com/yrcunha/my-like-crypto-server/src/proto/gen"
)

func TestUnmarshalVote(t *testing.T) {
	body := &gen.UpvoteReq{
		Name: 0,
	}
	upvote := model.UnmarshalVote(body.Name.String(), "upvote")
	assert.Equal(t, upvote.Name, "BTC")
	assert.Equal(t, upvote.Upvote, true)
	assert.Equal(t, upvote.Downvote, false)
	downvote := model.UnmarshalVote(body.Name.String(), "downvote")
	assert.Equal(t, downvote.Name, "BTC")
	assert.Equal(t, downvote.Upvote, false)
	assert.Equal(t, downvote.Downvote, true)
}

func TestUnmarshalCrypto(t *testing.T) {
	body := &gen.CreateCryptoReq{
		Name: 0,
	}
	upvote := model.UnmarshalCrypto(body.Name.String())
	assert.Equal(t, upvote.Name, "BTC")
	assert.Equal(t, upvote.Upvote, 0)
	assert.Equal(t, upvote.Downvote, 0)
}
