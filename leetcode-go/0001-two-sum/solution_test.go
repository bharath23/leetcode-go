package leetcode0001

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgroeneveld/trial/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "test 1",
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
			"%s failed",
			test.name,
		)
	}
}
