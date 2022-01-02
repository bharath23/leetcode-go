package leetcode0448

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
		nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
		want: []int{5, 6},
	},
	{
		name: "test 2",
		nums: []int{1, 1},
		want: []int{2},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := findDisappearedNumbersV0(nums)
		assert.ElementsMatchf(t, test.want, have, "%s: disappeared numbers do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := findDisappearedNumbersV1(nums)
		assert.ElementsMatchf(t, test.want, have, "%s: disappeared numbers do not match", test.name)
	}
}
