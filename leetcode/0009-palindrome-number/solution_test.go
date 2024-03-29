package leetcode0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	x    int
	want bool
}{
	{
		name: "test 1",
		x:    121,
		want: true,
	},
	{
		name: "test 2",
		x:    -121,
		want: false,
	},
	{
		name: "test 3",
		x:    10,
		want: false,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := isPalindromeV0(test.x)
		assert.Equalf(t, test.want, have, "%s: incorrect palindrome test", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := isPalindromeV1(test.x)
		assert.Equalf(t, test.want, have, "%s: incorrect palindrome test", test.name)
	}
}
