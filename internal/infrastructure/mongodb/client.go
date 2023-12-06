package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
	database  *mongo.Database
}

func NewMongoClient(uri, dbName string) IDatabaseInterfaceProtocol {
	var m *MongoClient = &MongoClient{}
	if err := m.Connect(context.Background(), uri); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("[mongodb]: successfully connected.")
	m.database = m.client.Database(dbName)
	return m
}

func (conn *MongoClient) Connect(ctx context.Context, uri string) error {
	var err error
	conn.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return err
}

func (conn *MongoClient) Disconnect() {
	if err := conn.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (conn *MongoClient) GetDatabase() *mongo.Database {
	return conn.database
}
