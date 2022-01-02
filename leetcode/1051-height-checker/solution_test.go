package leetcode1051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name    string
	heights []int
	want    int
}{
	{
		name:    "test 1",
		heights: []int{1, 1, 4, 2, 1, 3},
		want:    3,
	},
	{
		name:    "test 2",
		heights: []int{5, 1, 2, 3, 4},
		want:    5,
	},
	{
		name:    "test 3",
		heights: []int{1, 2, 3, 4, 5},
		want:    0,
	},
	{
		name:    "test 4",
		heights: []int{1, 1, 1, 1, 1, 1, 1, 1},
		want:    0,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		heights := make([]int, len(test.heights))
		copy(heights, test.heights)
		have := heightCheckerV0(heights)
		assert.Equalf(t, test.want, have, "%s: number of height mismatched indices differ", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		heights := make([]int, len(test.heights))
		copy(heights, test.heights)
		have := heightCheckerV1(heights)
		assert.Equalf(t, test.want, have, "%s: number of height mismatched indices differ", test.name)
	}
}
