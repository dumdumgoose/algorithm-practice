package common

type ForwardDirectionalNode struct {
	Next *ForwardDirectionalNode
	Data interface{}
}

type BackwardDirectionalNode struct {
	Prev *BackwardDirectionalNode
	Data interface{}
}

type DoubleDirectionalNode struct {
	Prev *DoubleDirectionalNode
	Next *DoubleDirectionalNode
	Data interface{}
}
