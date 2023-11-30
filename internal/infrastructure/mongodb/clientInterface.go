package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type IDatabaseInterfaceProtocol interface {
	Connect() *mongo.Client
	Disconnect(*mongo.Client)
}
