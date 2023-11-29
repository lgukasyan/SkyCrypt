package mongodb

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
  once   sync.Once
)

type MongoClient struct {
	Options *options.ClientOptions
	DBname  string
}

func NewMongoClient(uri string, dbName string) IDatabaseInterfaceProtocol {
	return &MongoClient{
		Options: options.Client().ApplyURI(uri),
		DBname: dbName,
	}
}

func (conn *MongoClient) Connect() {
	once.Do(func() {
		var err error
		client, err = mongo.Connect(context.TODO(), conn.Options)
		
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		log.Println("[mongodb]: successfully connected.")
	})
}

func (conn *MongoClient) Disconnect() {
	if client == nil {
		return
	}

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("[mongodb]: connection to mongodb closed.")
}

func GetMongoClientInstance() *mongo.Client {
	if client == nil {
		log.Fatal("[mongodb]: client is not initialized.")
		return nil
	}

	return client
}