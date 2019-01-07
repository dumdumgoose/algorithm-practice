package p2p

type Peer interface {
	Connect() error
	Disconnect()
	Send(message Message) error
}
