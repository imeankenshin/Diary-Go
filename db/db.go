package db

import (
	"context"
	"fmt"

	"first_go/ui"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	// MongoDBに接続するためのURIを定義する
	uri := "mongodb://localhost:27017"
	fmt.Println("Trying to connect " + uri + " ...")
	// MongoDBに接続するためのオプションを定義する
	clientOptions := options.Client().ApplyURI(uri)
	// MongoDBに接続する
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	// MongoDBに接続できたことを確認する
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	ui.Greenln("Connected to MongoDB!")
	return client, nil
}
