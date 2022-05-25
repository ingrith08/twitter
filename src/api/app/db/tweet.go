package db

import (
	"context"
	"time"
	"twitter_gin/internal/tweet/core/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const nameCollectionTweet = "tweet"

func (r *MongoRepository) InsertTweet(tweet entity.SaveTweet) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	register := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}
	result, err := r.db.Collection(nameCollectionTweet).InsertOne(ctx, register)
	objID, _ := result.InsertedID.(primitive.ObjectID)
	if err != nil {
		return "", err
	}

	return objID.String(), nil
}

func (r *MongoRepository) GetTweet(ID string, page int64) ([]*entity.ResponseTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result []*entity.ResponseTweet
	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"userid": objId,
	}

	tweetOptions := options.Find()
	tweetOptions.SetLimit(20)
	tweetOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	tweetOptions.SetSkip((page - 1) * 20)

	cursor, err := r.db.Collection(nameCollectionTweet).Find(ctx, condition, tweetOptions)
	if err != nil {
		return result, false
	}

	for cursor.Next(context.TODO()) {

		var register entity.ResponseTweet
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}

func (r *MongoRepository) DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(ID)
	objUserID, _ := primitive.ObjectIDFromHex(UserID)

	condition := bson.M{
		"_id":    objID,
		"userid": objUserID,
	}

	_, err := r.db.Collection(nameCollectionTweet).DeleteOne(ctx, condition)
	return err
}
