package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestMergeSort(t *testing.T) {
	err := testdata.RunSortTest(MergeSort)
	if err != nil {
		t.Error(err)
	}
}
