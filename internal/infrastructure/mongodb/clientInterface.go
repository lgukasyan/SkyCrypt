package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDatabaseInterfaceProtocol interface {
	Connect(context.Context, string) error
	Disconnect()
	GetDatabase() *mongo.Database
}
