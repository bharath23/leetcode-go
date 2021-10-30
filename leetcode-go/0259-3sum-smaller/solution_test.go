package leetcode0259

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
		nums:   []int{-2, 0, 1, 3},
		target: 2,
		want:   2,
	},
	{
		name:   "test 2",
		nums:   []int{},
		target: 0,
		want:   0,
	},
	{
		name:   "test 3",
		nums:   []int{0},
		target: 0,
		want:   0,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumSmallerV0(nums, test.target)
		assert.Equalf(t, test.want, have, "%s: number of combinations do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumSmallerV1(nums, test.target)
		assert.Equalf(t, test.want, have, "%s: number of combinations do not match", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumSmallerV2(nums, test.target)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
