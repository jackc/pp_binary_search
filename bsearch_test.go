package pp_binary_search_test

import (
	"testing"

	"github.com/jackc/pp_binary_search"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		collection []int
		query      int
		present    bool
		idx        int
	}{
		{
			[]int{1, 3, 7, 12, 65, 78},
			1,
			true,
			0,
		},
		{
			[]int{1, 3, 7, 12, 65, 78},
			78,
			true,
			5,
		},
		{
			[]int{1, 3, 7, 12, 65, 78},
			12,
			true,
			3,
		},
		{
			[]int{1, 3, 7, 12, 65, 78},
			8,
			false,
			-1,
		},
		{
			[]int{},
			8,
			false,
			-1,
		},
	}

	for i, tt := range tests {
		idx, present := pp_binary_search.Search(len(tt.collection), func(i int) int {
			return tt.query - tt.collection[i]
		})
		if present != tt.present {
			t.Errorf("%d. Expected present to be %v, but it was %v", i, tt.present, present)
		}
		if idx != tt.idx {
			t.Errorf("%d. Expected idx to be %v, but it was %v", i, tt.idx, idx)
		}
	}
}
