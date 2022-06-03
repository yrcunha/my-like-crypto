package repositorie

import (
	"context"

	"exemple.com/my-like-crypto-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVotes(collection *mongo.Collection, ctx context.Context, vote *model.Votes) error {
	data := bson.M{
		"_id":     vote.ID,
		"author":  vote.Author,
		"myvotes": vote.MyVotes,
	}
	_, insertError := collection.InsertOne(ctx, data)
	if insertError != nil {
		return insertError
	}
	return nil
}

func UpdateVotes(collection *mongo.Collection, ctx context.Context, vote *model.Votes) error {
	filter := bson.M{"_id": bson.M{"$eq": vote.ID}}
	data := bson.M{
		"$set": bson.M{
			"author":  vote.Author,
			"myvotes": vote.MyVotes,
		},
	}
	_, insertError := collection.UpdateOne(ctx, filter, data)
	if insertError != nil {
		return insertError
	}
	return nil
}

func DeleteVotes(collection *mongo.Collection, ctx context.Context, vote string) error {
	_, insertError := collection.DeleteOne(ctx, bson.M{"_id": vote})
	if insertError != nil {
		return insertError
	}
	return nil
}
