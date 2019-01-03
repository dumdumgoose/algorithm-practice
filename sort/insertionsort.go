package sort

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		k := data[i]
		j := i - 1
		for j >= 0 && data[j] > k {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = k
	}
}
