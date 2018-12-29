package structures

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common"
	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestMinHeap(t *testing.T) {
	testData := testdata.GetAllTestSliceCopies()
	for name, slice := range testData {
		heap, err := NewHeap(slice, MinHeap, len(slice))
		if err != nil {
			t.Errorf("failed to create new heap for slice %s, err: %s", name, err)
		}
		sortedSlice := make([]int, heap.Size())
		counter := 0
		for {
			popped, err := heap.Pop()
			if err != nil {
				break
			}
			sortedSlice[counter] = popped
			counter++
		}
		expected, _ := testdata.GetSortedSliceCopy(name)
		if !common.CompareArrays(sortedSlice, expected) {
			t.Errorf("result is not properly sorted, expected: %v, actual: %v", expected, sortedSlice)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	testData := testdata.GetAllTestSliceCopies()
	for name, slice := range testData {
		heap, err := NewHeap(slice, MaxHeap, len(slice))
		if err != nil {
			t.Errorf("failed to create new heap for slice %s, err: %s", name, err)
		}
		sortedSlice := make([]int, heap.Size())
		counter := 0
		for {
			popped, err := heap.Pop()
			if err != nil {
				break
			}
			sortedSlice[counter] = popped
			counter++
		}
		expected, _ := testdata.GetSortedSliceCopy(name)
		common.ReverseSortIntArray(expected)
		if !common.CompareArrays(sortedSlice, expected) {
			t.Errorf("result is not properly sorted, expected: %v, actual: %v", expected, sortedSlice)
		}
	}
}
