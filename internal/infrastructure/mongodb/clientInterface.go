package mongodb

type IDatabaseInterfaceProtocol interface {
	Connect()
	Disconnect()
}

