package leetcode0005

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "cbbd",
			want: "bb",
		},
	}

	for _, test := range tests {
		have := longestPalindrome(test.s)
		assert.Equal(t, test.want, have, "want: %s, have: %s", test.want, have)
	}
}
