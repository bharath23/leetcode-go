package leetcode1480

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "test 2",
			nums: []int{1, 1, 1, 1, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test 3",
			nums: []int{3, 1, 2, 10, 1},
			want: []int{3, 4, 6, 16, 17},
		},
	}

	for _, test := range tests {
		have := runningSum(test.nums)
		assert.Equalf(t, test.want, have, "%s: runningSum result does not match", test.name)
	}
}
