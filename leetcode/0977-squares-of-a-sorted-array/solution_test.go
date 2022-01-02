package leetcode0977

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
		nums: []int{-4, -1, 0, 3, 10},
		want: []int{0, 1, 9, 16, 100},
	},
	{
		name: "test 1",
		nums: []int{-7, -3, 2, 3, 11},
		want: []int{4, 9, 9, 49, 121},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := sortedSquaresV0(test.nums)
		assert.Equalf(t, test.want, have, "%s: sorted squares do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := sortedSquaresV1(test.nums)
		assert.Equalf(t, test.want, have, "%s: sorted squares do not match", test.name)
	}
}
