package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoClient(uri, dbName string) IDatabaseInterfaceProtocol {
	var m *MongoClient = &MongoClient{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := m.Connect(ctx, uri); err != nil {
		log.Fatal(err.Error())
	}

	m.database = m.client.Database(dbName)
	return m
}

func (conn *MongoClient) Connect(ctx context.Context, uri string) error {
	log.Println("attempting to connect to the database...")

	var err error

	conn.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = conn.client.Ping(ctx, nil); err != nil {
		panic("connection to mongodb failed.")
	}

	log.Println("[mongodb]: successfully connected.")
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
