package models

import (
	"context"
	"time"

	"fileserver/config"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert[T any](collectionName string, data *T) (interface{}, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID, nil
}

func UpdateOne[F any, T any](collectionName string, filter *F, update *T) (int64, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	flatBSON, err := flatbson.Flatten(update)
	if err != nil {
		return 0, err
	}

	updateQ := bson.M{
		"$set": flatBSON,
	}

	// Find documents
	res, err := collection.UpdateOne(ctx, filter, updateQ)

	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func UpdateMany(collectionName string, filter bson.M, update bson.M) (int64, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	// Find documents
	res, err := collection.UpdateMany(ctx, filter, update)

	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func Find[F any, T any](collectionName string, filter *F) ([]T, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	// Find documents
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	// Decode documents
	var results []T
	for cursor.Next(context.TODO()) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
func FindWithOptions[F any, T any](collectionName string, filter *F, options *options.FindOptions) ([]T, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	// Find documents
	cursor, err := collection.Find(ctx, filter, options)

	if err != nil {
		return nil, err
	}

	// Decode documents
	var results []T
	for cursor.Next(context.TODO()) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func DeleteOne[F any](collectionName string, filter *F) (int64, error) {

	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func DeleteMany[F any](collectionName string, filter *F) (int64, error) {
	// Get the collection
	collection := config.DB.GetCol(collectionName)

	// context
	ctx, cancle := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancle()

	res, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
