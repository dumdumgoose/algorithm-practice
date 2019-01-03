package testdata

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/azraeljack/algorithm-practice/common"
)

var (
	errNotFound = errors.New("not found")
)

var (
	sliceMap = map[string][]int{
		"nilSlice":            nil,
		"emptySlice":          {},
		"singleSlice":         {0},
		"doubleSlice":         generateRandomIntSlice(2, 10000),
		"tripleSlice":         generateRandomIntSlice(3, 10000),
		"equalSlice":          {1, 1, 1, 1, 1, 1, 1},
		"sortedSlice":         {1, 2, 3, 4, 5, 6, 7},
		"reverseSortedSlice":  {7, 6, 5, 4, 3, 2, 1},
		"randomizedSlice":     generateRandomIntSlice(7, 10000),
		"longRandomizedSlice": generateRandomIntSlice(100000, 10000000),
	}

	sortedSliceMap = map[string][]int{
		"nilSlice":            common.SortIntArrayWithCopy(sliceMap["nilSlice"]),
		"emptySlice":          common.SortIntArrayWithCopy(sliceMap["emptySlice"]),
		"singleSlice":         common.SortIntArrayWithCopy(sliceMap["singleSlice"]),
		"doubleSlice":         common.SortIntArrayWithCopy(sliceMap["doubleSlice"]),
		"tripleSlice":         common.SortIntArrayWithCopy(sliceMap["tripleSlice"]),
		"equalSlice":          common.SortIntArrayWithCopy(sliceMap["equalSlice"]),
		"sortedSlice":         common.SortIntArrayWithCopy(sliceMap["sortedSlice"]),
		"reverseSortedSlice":  common.SortIntArrayWithCopy(sliceMap["reverseSortedSlice"]),
		"randomizedSlice":     common.SortIntArrayWithCopy(sliceMap["randomizedSlice"]),
		"longRandomizedSlice": common.SortIntArrayWithCopy(sliceMap["longRandomizedSlice"]),
	}
)

func GetAllTestSliceCopies() map[string][]int {
	copied := make(map[string][]int, len(sliceMap))
	for name, slice := range sliceMap {
		sliceCopy := make([]int, len(slice))
		copy(sliceCopy, slice)
		copied[name] = sliceCopy
	}
	return copied
}

func GetAllSortedSliceCopies() map[string][]int {
	copied := make(map[string][]int, len(sortedSliceMap))
	for name, slice := range sortedSliceMap {
		sliceCopy := make([]int, len(slice))
		copy(sliceCopy, slice)
		copied[name] = sliceCopy
	}
	return copied
}

func GetSliceCopy(name string) ([]int, error) {
	toBeCopied, ok := sliceMap[name]
	if !ok {
		return nil, errNotFound
	} else if toBeCopied == nil {
		return toBeCopied, nil
	}

	buf := make([]int, len(toBeCopied))
	copy(buf, toBeCopied)
	return buf, nil
}

func GetSortedSliceCopy(name string) ([]int, error) {
	toBeCopied, ok := sortedSliceMap[name]
	if !ok {
		return nil, errNotFound
	} else if toBeCopied == nil {
		return toBeCopied, nil
	}

	buf := make([]int, len(toBeCopied))
	copy(buf, toBeCopied)
	return buf, nil
}

func RunSortTest(sortFunc func([]int)) error {
	for name, slice := range sliceMap {
		testData := make([]int, len(slice))
		copy(testData, slice)
		sortFunc(testData)
		if !common.CompareArraysIdentical(testData, sortedSliceMap[name]) {
			return fmt.Errorf("failed to sort the array %s, expected %v, got %v", name, sortedSliceMap[name], testData)
		}
	}
	return nil
}

func generateRandomIntSlice(size, max int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}
