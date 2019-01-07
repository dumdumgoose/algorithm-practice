package p2p

type Consensus interface {
	Update(key string, value string)
	Handle(message Message) error
}
