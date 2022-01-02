package leetcode0487

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
		nums: []int{1, 0, 1, 1, 0},
		want: 4,
	},
	{
		name: "test 2",
		nums: []int{1, 0, 1, 1, 0, 1},
		want: 4,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := findMaxConsecutiveOnesV0(test.nums)
		assert.Equalf(t, test.want, have, "%s: max consecutive ones mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := findMaxConsecutiveOnesV1(test.nums)
		assert.Equalf(t, test.want, have, "%s: max consecutive ones mismatch", test.name)
	}
}
