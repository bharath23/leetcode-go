package leetcode0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int
		k    int
	}{
		{
			name: "test 1",
			nums: []int{1, 1, 2},
			want: []int{1, 2},
			k:    2,
		},
		{
			name: "test 2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
			k:    5,
		},
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		k := removeDuplicates(nums)
		if assert.Equalf(t, test.k, k, "%s: expected length not equal", test.name) {
			assert.ElementsMatchf(t, test.want, nums[:k], "%s: elements do not match", test.name)
		}
	}
}
