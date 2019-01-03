package structures

import (
	"math/big"
	"testing"

	"github.com/azraeljack/algorithm-practice/common"
	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestStack(t *testing.T) {
	testData := testdata.GetAllSortedSliceCopies()

	for name, slice := range testData {
		s, err := NewStack(big.NewInt(int64(len(slice))))
		if err != nil {
			if name == "nilSlice" || name == "emptySlice" {
				continue
			}
			t.Errorf("failed to create testing stack for slice %s", name)
		}
		for index, number := range slice {
			err := s.Push(number)
			if err != nil {
				t.Errorf("failed to push the %d item of slice %s to stack", index, name)
			}
		}
		counter := 0
		result := make([]int, len(slice))
		for {
			popped, err := s.Pop()
			if err != nil {
				break
			}
			result[counter] = popped.(int)
			counter++
		}
		if !common.CompareArraysReversed(result, slice) {
			common.ReverseSortIntArray(slice)
			t.Errorf("result is not in the correct order, expected %v, got %v", slice, result)
		}
	}
}
