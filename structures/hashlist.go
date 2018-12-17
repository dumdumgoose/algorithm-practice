package structures

import "github.com/azraeljack/algorithm-practice/common"

type HashLinkedList struct {
	root, tail *common.DoubleDirectionalNode
	hashMap    map[interface{}]*common.DoubleDirectionalNode
}

func NewHashLinkedList(size uint64) *HashLinkedList {
	return &HashLinkedList{
		root:    nil,
		tail:    nil,
		hashMap: make(map[interface{}]*common.DoubleDirectionalNode, size),
	}
}

func (h *HashLinkedList) Get(key interface{}) (interface{}, bool) {
	if value, ok := h.hashMap[key]; !ok {
		return nil, false
	} else {
		return value.Data, true
	}
}

func (h *HashLinkedList) Put(key interface{}, value interface{}) {
	_, ok := h.hashMap[key]
	if ok {
		h.hashMap[key].Data = value
		return
	}
	node := &common.DoubleDirectionalNode{
		Prev: nil,
		Next: nil,
		Data: value,
	}
	if h.root == nil {
		h.root = node
		h.tail = h.root
	} else {
		h.tail.Next = node
		node.Prev = h.tail
	}
	h.hashMap[key] = node
}

func (h *HashLinkedList) Delete(key interface{}) {
	node, ok := h.hashMap[key]
	if !ok {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	delete(h.hashMap, key)
}

func (h *HashLinkedList) Iterate() *common.DoubleDirectionalNode {
	return h.root
}
