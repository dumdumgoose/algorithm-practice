package common

import "sort"

func SwapIntArrayElements(data []int, index1, index2 int) {
	temp := data[index1]
	data[index1] = data[index2]
	data[index2] = temp
}

func CompareArrays(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func SortIntArray(data []int) {
	sort.Ints(data)
}

func SortIntArrayWithCopy(data []int) []int {
	sorted := make([]int, len(data))
	copy(sorted, data)
	sort.Ints(sorted)
	return sorted
}

func ReverseSortIntArray(data []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
}

func IsIntArraySorted(data []int) bool {
	return sort.IntsAreSorted(data)
}

func IsIntArrayReverseSorted(data []int) bool {
	for i := 0; i < len(data); i++ {
		if i > 0 && data[i] > data[i-1] {
			return false
		}
	}
	return true
}
