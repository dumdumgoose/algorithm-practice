package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestHeapSort(t *testing.T) {
	err := testdata.RunSortTest(HeapSort)
	if err != nil {
		t.Error(err)
	}
}
