package structures

import (
	"errors"
	"github.com/azraeljack/algorithm-practice/common"
	"math/big"
)

var (
	errQueueEmpty = errors.New("queue is empty")
	errQueueFull  = errors.New("queue is full")
)

type NonBlockingQueue struct {
	root, tail    *common.ForwardDirectionalNode
	length, limit *big.Int
}

func NewQueue(size *big.Int) *NonBlockingQueue {
	return &NonBlockingQueue{
		root: nil, tail: nil, length: new(big.Int).Set(common.Big0), limit: new(big.Int).Set(size),
	}
}

func (q *NonBlockingQueue) Peek() (interface{}, error) {
	if q.root == nil {
		return nil, errQueueEmpty
	}
	return q.root.Data, nil
}

func (q *NonBlockingQueue) Size() *big.Int {
	return q.length
}

func (q *NonBlockingQueue) Get() (interface{}, error) {
	if q.root == nil {
		return nil, errQueueEmpty
	}
	q.length.Sub(q.length, common.Big1)
	result := q.root.Data
	q.root = q.root.Next
	return result, nil
}

func (q *NonBlockingQueue) Put(data interface{}) error {
	node := common.ForwardDirectionalNode{Next: nil, Data: data}
	if q.length.Cmp(q.limit) >= 0 {
		return errQueueFull
	}
	if q.root == nil {
		q.root = &node
		q.tail = q.root
	} else {
		q.tail.Next = &node
	}
	q.length.Add(q.length, common.Big1)
	return nil
}

func (q *NonBlockingQueue) Empty() bool {
	if q.root == nil {
		return true
	}
	return false
}
