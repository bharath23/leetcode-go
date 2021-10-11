package leetcode0485

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "test 2",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}

	for _, test := range tests {
		have := findMaxConsecutiveOnes(test.nums)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
