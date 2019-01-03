package sort

func MergeSort(data []int) {
	if len(data) <= 1 {
		return
	}

	mid := len(data) / 2
	left := make([]int, mid)
	right := make([]int, len(data)-mid)
	copy(left, data[:mid])
	copy(right, data[mid:])

	MergeSort(left)
	MergeSort(right)

	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		data[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		data[k] = right[j]
		j++
		k++
	}
}
