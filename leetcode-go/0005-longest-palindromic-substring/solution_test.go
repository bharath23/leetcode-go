package leetcode0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equalf(t, test.want, have, "%s: longest palindrome does not match", test.name)
	}
}
