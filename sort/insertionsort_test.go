package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestInsertionSort(t *testing.T) {
	err := testdata.RunSortTest(InsertionSort)
	if err != nil {
		t.Error(err)
	}
}
