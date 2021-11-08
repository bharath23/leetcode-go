package leetcode0238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	nums []int
	want []int
}{
	{
		name: "test 1",
		nums: []int{0, 1, 0, 3, 12},
		want: []int{1, 3, 12, 0, 0},
	},
	{
		name: "test 2",
		nums: []int{0},
		want: []int{0},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		moveZeroesV0(nums)
		assert.Equalf(t, test.want, nums, "%s: moved array mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		moveZeroesV1(nums)
		assert.Equalf(t, test.want, nums, "%s: moved array mismatch", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		moveZeroesV2(nums)
		assert.Equalf(t, test.want, nums, "%s: moved array mismatch", test.name)
	}
}
