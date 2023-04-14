package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	// MongoDBに接続するためのURIを定義する
	uri := "mongodb://root:root@localhost:27017"
	fmt.Println("Trying to connect " + uri + " ...")
	// MongoDBに接続するためのオプションを定義する
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDBに接続するためのクライアントを生成する
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
