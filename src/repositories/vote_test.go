package repositorie_test

import (
	"context"
	"testing"

	"exemple.com/my-like-crypto-server/src/model"
	"exemple.com/my-like-crypto-server/src/repositories"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx           = context.TODO()
	clientOptions = options.Client().ApplyURI("mongodb://docker:mongo@localhost:27017/")
	client, _     = mongo.Connect(ctx, clientOptions)
	_             = client.Ping(ctx, nil)
	collection    = client.Database("my-like-crypto-test").Collection("vote-test")
	votes         = &model.Votes{}
)

func TestCreateRepositories(t *testing.T) {
	votes.ID = "6a4493a4-a449-4e43-b971-2f42eb591d9f"
	createError := repositorie.CreateVotes(collection, ctx, votes)
	assert.Nil(t, createError)
	collection.Drop(ctx)
}

func TestUpdateRepositories(t *testing.T) {
	votes.ID = "8353374a-efab-4401-b500-6da0101acc03"
	votes.Author = "Leoncio"
	votes.MyVotes[0].Name = "BTC"
	repositorie.CreateVotes(collection, ctx, votes)
	updateError := repositorie.UpdateVotes(collection, ctx, votes)
	assert.Nil(t, updateError)
	collection.Drop(ctx)
}

func TestDeleteRepositories(t *testing.T) {
	votes.ID = "f27a3faa-94ba-4c55-839a-06ce259dbdd6"
	repositorie.CreateVotes(collection, ctx, votes)
	deleteError := repositorie.DeleteVotes(collection, ctx, votes.ID)
	assert.Nil(t, deleteError)
	collection.Drop(ctx)
}
