package sort

import "github.com/azraeljack/algorithm-practice/common"

func HeapSort(data []int) {
	size := len(data)
	for i := size/2 - 1; i >= 0; i-- {
		maxHeapify(data, i, size-1)
	}
	for i := size - 1; i > 0; i-- {
		common.SwapIntArrayElements(data, 0, i)
		maxHeapify(data, 0, i-1)
	}
}

// same as percolate down
func maxHeapify(data []int, start, end int) {
	parent := start
	child := parent*2 + 1
	for child <= end {
		if child+1 <= end && data[child] < data[child+1] {
			child++
		}
		if data[parent] > data[child] {
			break
		} else {
			common.SwapIntArrayElements(data, parent, child)
			parent = child
			child = parent*2 + 1
		}
	}
}
