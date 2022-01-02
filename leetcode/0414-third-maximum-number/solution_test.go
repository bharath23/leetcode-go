package leetcode0414

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	nums []int
	want int
}{
	{
		name: "test 1",
		nums: []int{3, 2, 1},
		want: 1,
	},
	{
		name: "test 2",
		nums: []int{1, 2},
		want: 2,
	},
	{
		name: "test 3",
		nums: []int{2, 2, 3, 1},
		want: 1,
	},
	{
		name: "test 4",
		nums: []int{1, 2, -2147483648},
		want: -2147483648,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := thirdMaxV0(nums)
		assert.Equalf(t, test.want, have, "%s: incorrect third maximum", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := thirdMaxV1(nums)
		assert.Equalf(t, test.want, have, "%s: incorrect third maximum", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := thirdMaxV2(nums)
		assert.Equalf(t, test.want, have, "%s: incorrect third maximum", test.name)
	}
}

func TestSolutionV3(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := thirdMaxV3(nums)
		assert.Equalf(t, test.want, have, "%s: incorrect third maximum", test.name)
	}
}
