package structures

import (
	"errors"

	"github.com/azraeljack/algorithm-practice/common"
)

/*
Implementation of min heap
Structure:
	parent: i
	left child: 2i + 1
	right child: 2i + 2
*/

var (
	errHeapEmpty      = errors.New("heap is empty")
	errHeapFull       = errors.New("heap is full")
	errNotImplemented = errors.New("not implemented")
)

type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap struct {
	data        []int
	capacity    int                 // 0 means infinite
	compareFunc func(a, b int) bool // if a has greater priority than b then return true else false
}

func greater(a, b int) bool {
	return a > b
}

func less(a, b int) bool {
	return a < b
}

func NewHeap(data []int, heapType HeapType, capacity int) (*Heap, error) {
	heap := &Heap{capacity: capacity}

	if heapType == MinHeap {
		heap.compareFunc = greater
	} else if heapType == MaxHeap {
		heap.compareFunc = less
	} else {
		return nil, errNotImplemented
	}

	if data != nil {
		for _, number := range data {
			if err := heap.Insert(number); err != nil {
				return nil, err
			}
		}
	}
	return heap, nil
}

func (h *Heap) Insert(data int) error {
	if h.capacity != 0 && h.Size() >= h.capacity {
		return errHeapFull
	}
	h.data = append(h.data, data)
	h.percolateUp(h.Size() - 1)
	return nil
}

func (h *Heap) Get() (int, error) {
	if h.isEmpty() {
		return 0, errHeapEmpty
	}
	return h.data[0], nil
}

func (h *Heap) Pop() (int, error) {
	if h.isEmpty() {
		return 0, errHeapEmpty
	}
	popped := h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.percolateDown(0)
	return popped, nil
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) isFull() bool {
	if len(h.data) >= h.capacity {
		return true
	}
	return false
}

func (h *Heap) isEmpty() bool {
	if h.Size() <= 0 {
		return true
	}
	return false
}

func (h *Heap) percolateDown(node int) {
	size := len(h.data)
	parent := node
	child := node*2 + 1
	for child < size {
		if child+1 < size && h.compareFunc(h.data[child], h.data[child+1]) {
			child++
		}
		if h.compareFunc(h.data[parent], h.data[child]) {
			h.swap(parent, child)
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
}

func (h *Heap) percolateUp(node int) {
	child := node
	for {
		parent := (child - 1) / 2
		if parent < 0 || parent == child || h.compareFunc(h.data[child], h.data[parent]) {
			break
		}
		h.swap(parent, child)
		child = parent
	}
}

func (h *Heap) swap(index1, index2 int) {
	common.SwapIntArrayElements(h.data, index1, index2)
}
