package leetcode0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name   string
	height []int
	want   int
}{
	{
		name:   "test 1",
		height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		want:   49,
	},
	{
		name:   "test 2",
		height: []int{1, 1},
		want:   1,
	},
	{
		name:   "test 3",
		height: []int{4, 3, 2, 1, 4},
		want:   16,
	},
	{
		name:   "test 4",
		height: []int{1, 2, 1},
		want:   2,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := maxAreaV0(test.height)
		assert.Equalf(t, test.want, have, "%s: maximum area does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := maxAreaV1(test.height)
		assert.Equalf(t, test.want, have, "%s: maximum area does not match", test.name)
	}
}
