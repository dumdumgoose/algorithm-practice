package sort

import "github.com/azraeljack/algorithm-practice/common"

/*
Implementations of varies quick sort
*/

/*
Regular quick sort implementation
Time Complexity:
	Best: O(nlogn)
	Average: O(nlogn)
	Worst: O(n^2)
Space Complexity:
	Average: O(logn)
	Worst: O(n)
*/
func QuickSort(data []int) {
	if data == nil {
		return
	}
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		quickSort(data, low, pivot)
		quickSort(data, pivot+1, high)
	}
}

func partition(data []int, low, high int) int {
	start := low
	for low < high {
		if data[start] <= data[high] {
			high--
		} else if data[start] >= data[low] {
			low++
		} else {
			common.SwapIntArrayElements(data, low, high)
		}
	}
	if data[start] > data[high] {
		common.SwapIntArrayElements(data, start, high)
	}
	return high
}

func MultiPivotQuickSort(data []int) {

}

func ExternalQuickSort(data []int) {

}

func ThreeWayRadixQuickSort(data []int) {

}

func QuickRadixSort(data []int) {

}

func PartialIncrementalQuickSort(data []int) {

}
