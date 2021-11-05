package leetcode0941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "test 1",
			arr:  []int{2, 1},
			want: false,
		},
		{
			name: "test 2",
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			name: "test 3",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "test 4",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
		{
			name: "test 5",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
	}

	for _, test := range tests {
		have := validMountainArray(test.arr)
		assert.Equalf(t, test.want, have, "%s: mountain array validity mismatch", test.name)
	}
}
