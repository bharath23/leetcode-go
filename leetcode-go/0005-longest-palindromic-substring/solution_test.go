package leetcode0005

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "test 2",
			s:    "cbbd",
			want: "bb",
		},
	}

	for _, test := range tests {
		have := longestPalindrome(test.s)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}
