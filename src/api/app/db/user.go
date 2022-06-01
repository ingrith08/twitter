package db

import (
	"context"
	"time"
	"twitter_gin/internal/user/core/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const nameCollection = "user"

func (r *MongoRepository) InsertUser(user entity.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := r.db.Collection(nameCollection).InsertOne(ctx, user)
	objID, _ := result.InsertedID.(primitive.ObjectID)
	if err != nil {
		return "", err
	}

	return objID.String(), nil
}

func (r *MongoRepository) CheckUser(email string) (entity.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	condition := bson.M{"email": email}

	var result entity.User
	err := r.db.Collection(nameCollection).FindOne(ctx, condition).Decode(&result)

	if err != nil {
		return result, false
	}
	return result, true

}

func (r *MongoRepository) GetUser(ID string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result entity.User
	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objId}
	err := r.db.Collection(nameCollection).FindOne(ctx, condition).Decode(&result)
	result.Password = ""
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *MongoRepository) UpdateUser(ID string, register map[string]interface{}) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := r.db.Collection(nameCollection).UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoRepository) ListUser(page int64, search string) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*entity.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := r.db.Collection(nameCollection).Find(ctx, query, findOptions)
	if err != nil {
		return results, err
	}

	for cursor.Next(context.TODO()) {
		var s entity.User
		err := cursor.Decode(&s)
		if err != nil {
			return results, err
		}

		s.Password = ""
		s.Biography = ""
		s.WebSite = ""
		s.Location = ""
		s.Banner = ""
		s.Email = ""

		results = append(results, &s)
	}
	err = cursor.Err()
	if err != nil {
		return results, err
	}
	cursor.Close(ctx)
	return results, nil
}
