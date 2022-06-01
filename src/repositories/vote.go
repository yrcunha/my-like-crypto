package repositorie

import (
	"context"

	"exemple.com/my-like-crypto-server/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVotes(collection *mongo.Collection, ctx context.Context, vote *model.Votes) error {
	_, err := collection.InsertOne(ctx, vote)
	if err == nil {
		return err
	}
	return nil
}
