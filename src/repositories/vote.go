package repositorie

import (
	"context"

	"exemple.com/my-like-crypto-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var data primitive.M

func UpvoteOrDownvote(collection *mongo.Collection, ctx context.Context, vote *model.Crypto, upvote bool) error {
	filter := bson.M{"crypto": bson.M{"$eq": vote.Name}}
	if upvote {
		data = bson.M{
			"$inc": bson.M{
				"upvote": 1,
			},
		}
	} else {
		data = bson.M{
			"$inc": bson.M{
				"downvote": 1,
			},
		}
	}
	_, err := collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}
	return nil
}

func CreateCrypto(collection *mongo.Collection, ctx context.Context, vote *model.Data) error {
	data := bson.M{
		"crypto":   vote.Name,
		"upvote":   vote.Upvote,
		"downvote": vote.Downvote,
	}
	_, insertError := collection.InsertOne(ctx, data)
	if insertError != nil {
		return insertError
	}
	return nil
}

func DeleteCrypto(collection *mongo.Collection, ctx context.Context, vote string) error {
	id, _ := primitive.ObjectIDFromHex(vote)
	_, deleteError := collection.DeleteOne(ctx, bson.M{"_id": id})
	if deleteError != nil {
		return deleteError
	}
	return nil
}
