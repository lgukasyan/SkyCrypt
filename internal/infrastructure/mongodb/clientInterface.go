package mongodb

type IDatabaseInterfaceProtocol interface {
	Connect() error
	Disconnect() error
}

