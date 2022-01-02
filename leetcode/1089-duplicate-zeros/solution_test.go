package leetcode1089

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	arr  []int
	want []int
}{
	{
		name: "test 1",
		arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
		want: []int{1, 0, 0, 2, 3, 0, 0, 4},
	},
	{
		name: "test 2",
		arr:  []int{1, 2, 3},
		want: []int{1, 2, 3},
	},
	{
		name: "test 3",
		arr:  []int{8, 4, 5, 0, 0, 0, 0, 7},
		want: []int{8, 4, 5, 0, 0, 0, 0, 0},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		arr := make([]int, len(test.arr))
		copy(arr, test.arr)
		duplicateZerosV0(arr)
		assert.Equalf(t, test.want, arr, "%s: duplicated array do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		arr := make([]int, len(test.arr))
		copy(arr, test.arr)
		duplicateZerosV1(arr)
		assert.Equalf(t, test.want, arr, "%s: duplicated array do not match", test.name)
	}
}
