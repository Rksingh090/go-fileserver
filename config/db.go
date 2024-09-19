package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_DB_URL))

	if err != nil {
		fmt.Println(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	return client
}

// Client instance
var mongoClient *mongo.Client = ConnectDB()

type DbStruct struct {
	db *mongo.Database
}

var DB = &DbStruct{
	db: mongoClient.Database(DATABASE_NAME),
}

func (r *DbStruct) GetCol(collectionName string) *mongo.Collection {
	return r.db.Collection(collectionName)
}
