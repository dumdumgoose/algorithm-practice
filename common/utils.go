package common

import "sort"

func SwapIntArrayElements(data []int, index1, index2 int) {
	temp := data[index1]
	data[index1] = data[index2]
	data[index2] = temp
}

func SortIntArray(data []int) []int {
	sorted := make([]int, len(data))
	copy(sorted, data)
	sort.Ints(sorted)
	return sorted
}

func IsIntArraySorted(data []int) bool {
	for i := 0; i < len(data); i++ {
		if i == 0 || data[i-1] <= data[i] {
			continue
		}
		return false
	}
	return true
}
