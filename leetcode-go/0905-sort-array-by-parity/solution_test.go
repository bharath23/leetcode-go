package leetcode0905

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
		nums: []int{3, 1, 2, 4},
		want: [][]int{{2, 4, 1, 3}, {2, 4, 3, 1}, {4, 2, 1, 3}, {4, 2, 3, 1}},
	},
	{
		name: "test 2",
		nums: []int{0},
		want: [][]int{{0}},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := sortArrayByParityV0(nums)
		assert.Containsf(t, test.want, have, "%s: does not match any valid parity sortint", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := sortArrayByParityV1(nums)
		assert.Containsf(t, test.want, have, "%s: does not match any valid parity sortint", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := sortArrayByParityV2(nums)
		assert.Containsf(t, test.want, have, "%s: does not match any valid parity sortint", test.name)
	}
}
