package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	options *options.ClientOptions
}

func NewMongoClient(uri string) (IDatabaseInterfaceProtocol) {
	return &MongoClient{
		options: options.Client().ApplyURI(uri),
	}
}

func (conn *MongoClient) Connect() (*mongo.Client) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, conn.options)
	if err != nil {
		panic(err)
	}

	log.Println("[mongodb]: successful connection to the database.")
	return client
}

func (conn *MongoClient) Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}



