package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestQuickSort(t *testing.T) {
	err := testdata.RunSortTest(QuickSort)
	if err != nil {
		t.Error(err)
	}
}
