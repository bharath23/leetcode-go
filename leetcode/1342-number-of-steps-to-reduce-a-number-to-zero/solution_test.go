package leetcode1342

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	num  int
	want int
}{
	{
		name: "test 1",
		num:  14,
		want: 6,
	},
	{
		name: "test 2",
		num:  8,
		want: 4,
	},
	{
		name: "test 3",
		num:  123,
		want: 12,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := numberOfStepsV0(test.num)
		assert.Equalf(t, test.want, have, "%s: number of steps mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := numberOfStepsV1(test.num)
		assert.Equalf(t, test.want, have, "%s: number of steps mismatch", test.name)
	}
}
