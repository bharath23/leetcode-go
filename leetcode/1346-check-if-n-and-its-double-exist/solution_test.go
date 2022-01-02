package leetcode1246

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
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "test 2",
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
		{
			name: "test 2",
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
	}

	for _, test := range tests {
		have := checkIfExist(test.arr)
		assert.Equalf(t, test.want, have, "%s: double check result mismatch", test.name)
	}
}
