package leetcode0003

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

var tests = []struct {
	name string
	s    string
	want int
}{
	{
		name: "test 1",
		s:    "abcabcbb",
		want: 3,
	},
	{
		name: "test 2",
		s:    "bbbbb",
		want: 1,
	},
	{
		name: "test 3",
		s:    "pwwkew",
		want: 3,
	},
}

func TestSolutionv0(t *testing.T) {
	for _, test := range tests {
		have := lengthOfLongestSubstringV0(test.s)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionv1(t *testing.T) {
	for _, test := range tests {
		have := lengthOfLongestSubstringV1(test.s)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}
