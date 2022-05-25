package db

import (
	"context"
	"time"
	"twitter_gin/internal/relation/core/entity"

	"go.mongodb.org/mongo-driver/bson"
)

const nameCollectionRelation = "relation"

func (r *MongoRepository) InsertRelation(relation entity.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.db.Collection(nameCollectionRelation).InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoRepository) DeleteRelation(relation entity.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.db.Collection(nameCollectionRelation).DeleteOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MongoRepository) GetRelation(relation entity.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	condicion := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	var result entity.Relation
	err := r.db.Collection(nameCollectionRelation).FindOne(ctx, condicion).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MongoRepository) ListTweets(ID string, page int) ([]entity.ListTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	var result []entity.ListTweets
	cursor, err := r.db.Collection(nameCollectionRelation).Aggregate(ctx, conditions)
	if err != nil {
		return result, false
	}
	err = cursor.All(context.TODO(), &result)

	if err != nil {
		return result, false
	}

	return result, true
}
