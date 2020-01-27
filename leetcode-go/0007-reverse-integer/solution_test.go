package leetcode0007

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "test 1",
			x:    123,
			want: 321,
		},
		{
			name: "test 2",
			x:    -123,
			want: -321,
		},
		{
			name: "test 3",
			x:    120,
			want: 21,
		},
		{
			name: "test 4",
			x:    1534236469,
			want: 0,
		},
	}

	for _, test := range tests {
		have := reverse(test.x)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}
