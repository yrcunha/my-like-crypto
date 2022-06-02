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
	collection *mongo.Collection
	ctx        = context.TODO()
)

func TestRepositories(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://docker:mongo@localhost:27017/")
	client, _ := mongo.Connect(ctx, clientOptions)
	_ = client.Ping(ctx, nil)
	collection = client.Database("my-like-crypto-test").Collection("vote-test")
	votes := &model.Votes{}
	createError := repositorie.CreateVotes(collection, ctx, votes)
	assert.Nil(t, createError)
}
