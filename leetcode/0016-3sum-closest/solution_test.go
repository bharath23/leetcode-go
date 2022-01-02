package leetcode0016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name   string
	nums   []int
	target int
	want   int
}{
	{
		name:   "test 1",
		nums:   []int{-1, 2, 1, -4},
		target: 1,
		want:   2,
	},
	{
		name:   "test 2",
		nums:   []int{0, 0, 0},
		target: 1,
		want:   0,
	},
	{
		name:   "test 3",
		nums:   []int{1, 1, -1, -1, 3},
		target: -1,
		want:   -1,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumClosestV0(nums, test.target)
		assert.Equalf(t, test.want, have, "%s: sum doesnt match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumClosestV1(nums, test.target)
		assert.Equalf(t, test.want, have, "%s: sum doesnt match", test.name)
	}
}
