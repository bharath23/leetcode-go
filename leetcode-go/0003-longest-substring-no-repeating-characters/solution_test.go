package solution

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

var tests = []struct {
	s    string
	want int
}{
	{
		s:    "abcabcbb",
		want: 3,
	},
	{
		s:    "bbbbb",
		want: 1,
	},
	{
		s:    "pwwkew",
		want: 3,
	},
}

func TestSolutionv0(t *testing.T) {
	for _, test := range tests {
		have := lengthOfLongestSubstringv0(test.s)
		assert.Equal(t, test.want, have, "want: %d, have: %d", test.want, have)
	}
}

func TestSolutionv1(t *testing.T) {
	for _, test := range tests {
		have := lengthOfLongestSubstringv1(test.s)
		assert.Equal(t, test.want, have, "want: %d, have: %d", test.want, have)
	}
}
