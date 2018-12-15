package sort

import "math"

func BucketSort(data []int) {
	if data == nil || len(data) == 0 {
		return
	}
	max := int(math.MinInt64)
	// find the largest number in the array
	for _, number := range data {
		if number > max {
			max = number
		}
	}
	// create a bucket with the length of the maximum value + 1 and count the appearance of values
	bucket := make([]int, max+1)
	for _, number := range data {
		bucket[number]++
	}
	// dump the bucket back to the array to be sorted
	arrayPointer := 0
	for index, number := range bucket {
		for i := 0; i < number; i++ {
			data[arrayPointer] = index
			arrayPointer++
		}
	}
}
