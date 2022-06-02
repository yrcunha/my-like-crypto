package repositorie

import (
	"context"
	"fmt"

	"exemple.com/my-like-crypto-server/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVotes(collection *mongo.Collection, ctx context.Context, vote *model.Votes) error {
	fmt.Print(vote)
	_, insertError := collection.InsertOne(ctx, vote)
	if insertError != nil {
		return insertError
	}
	return nil
}
