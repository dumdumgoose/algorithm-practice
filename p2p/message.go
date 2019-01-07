package p2p

type Message interface {
	Code() int
	From() string
	To() string
}
