package database

import (
	"context"
	"fmt"
	"getirAssignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

func Close() {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect() {
	uri := "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Minute)

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func Query(dataBase, col string, matchStage bson.D, groupStage bson.D) (cursor *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	cursor, err = collection.Aggregate(ctx, mongo.Pipeline{groupStage, matchStage})
	if err != nil {
		panic(err)
	}
	return
}

func FetchAllData(cursor *mongo.Cursor) (result []models.Record) {
	if err := cursor.All(ctx, &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
	return
}
