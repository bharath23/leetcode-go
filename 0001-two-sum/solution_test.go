package solution

import (
	//"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgroeneveld/trial/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}

	for _, test := range tests {
		have := twoSums(test.nums, test.target)
		assert.True(
			t,
			cmp.Equal(test.want, have),
			"want: %#v, have: %#v",
			test.want,
			have,
		)
	}
}
