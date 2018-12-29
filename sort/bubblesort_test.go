package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestBubbleSort(t *testing.T) {
	err := testdata.RunSortTest(BubbleSort)
	if err != nil {
		t.Error(err)
	}
}
