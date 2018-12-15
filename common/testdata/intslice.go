package testdata

import (
	"errors"
	"math/rand"
)

var (
	errNotFound = errors.New("not found")
)

var sliceMap = map[string][]int{
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

func generateRandomIntSlice(size, max int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}
