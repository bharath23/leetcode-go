package leetcode0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "test1",
			str:  "42",
			want: 42,
		},
		{
			name: "test2",
			str:  "   -42",
			want: -42,
		},
		{
			name: "test3",
			str:  "4193 with words",
			want: 4193,
		},
		{
			name: "test4",
			str:  "words and 987",
			want: 0,
		},
		{
			name: "test5",
			str:  "+1",
			want: 1,
		},
	}

	for _, test := range tests {
		have := atoi(test.str)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
