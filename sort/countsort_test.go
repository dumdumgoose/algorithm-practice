package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestCountSort(t *testing.T) {
	err := testdata.RunSortTest(CountSort)
	if err != nil {
		t.Error(err)
	}
}
