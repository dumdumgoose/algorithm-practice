package sort

import "github.com/azraeljack/algorithm-practice/common"

func BubbleSort(data []int) {
	if data == nil {
		return
	}
	for i := 0; i < len(data); i++ {
		if i == 0 {
			continue
		}
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				common.SwapIntArrayElements(data, j, j-1)
			}
		}
	}
}
