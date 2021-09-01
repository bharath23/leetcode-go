package leetcode0167

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			want:   []int{1, 2},
		},
		{
			name:   "test 2",
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{1, 3},
		},
		{
			name:   "test 3",
			nums:   []int{-1, 0},
			target: -1,
			want:   []int{1, 2},
		},
	}

	for _, test := range tests {
		have := twoSum(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
