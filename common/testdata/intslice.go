package testdata

import "errors"

var (
	errNotFound = errors.New("not found")
)

var sliceMap = map[string][]int{
	"nilSlice":            nil,
	"emptySlice":          []int{},
	"singleSlice":         []int{0},
	"doubleSlice":         []int{5, 1},
	"tripleSlice":         []int{9, 1, 2},
	"equalSlice":          []int{1, 1, 1, 1, 1, 1, 1},
	"sortedSlice":         []int{1, 2, 3, 4, 5, 6, 7},
	"reverseSortedSlice":  []int{7, 6, 5, 4, 3, 2, 1},
	"randomizedSlice":     []int{1, 5, 4, 9, 2, 3, 10},
	"longRandomizedSlice": []int{9, 2, 3, 3, 5, 2, 1, 8, 4, 6, 1, 2, 3, 88, 1, 1, 9, 8, 6, 51, 6, 581, 6, 8, 65, 16, 65, 6, 5, 1, 2, 5},
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
