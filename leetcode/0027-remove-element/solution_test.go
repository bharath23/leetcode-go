package leetcode0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	nums []int
	val  int
	want []int
	k    int
}{
	{
		name: "test 1",
		nums: []int{3, 2, 2, 3},
		val:  3,
		want: []int{2, 2},
		k:    2,
	},
	{
		name: "test 2",
		nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:  2,
		want: []int{0, 1, 4, 0, 3},
		k:    5,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		k := removeElementV0(nums, test.val)
		if assert.Equalf(t, test.k, k, "%s: expected length not equal", test.name) {
			assert.ElementsMatchf(t, test.want, nums[:k], "%s: elements do not match", test.name)
		}
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		k := removeElementV1(nums, test.val)
		if assert.Equalf(t, test.k, k, "%s: expected length not equal", test.name) {
			assert.ElementsMatchf(t, test.want, nums[:k], "%s: elements do not match", test.name)
		}
	}
}
