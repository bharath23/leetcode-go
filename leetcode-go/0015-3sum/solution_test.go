package leetcode0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	nums []int
	want [][]int
}{
	{
		name: "test 1",
		nums: []int{-1, 0, 1, 2, -1, -4},
		want: [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
	},
	{
		name: "test 2",
		nums: []int{},
		want: [][]int{},
	},
	{
		name: "test 3",
		nums: []int{0},
		want: [][]int{},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumV0(nums)
		assert.ElementsMatchf(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumV1(nums)
		assert.ElementsMatchf(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := threeSumV2(nums)
		assert.ElementsMatchf(t, test.want, have, "%s failed", test.name)
	}
}
