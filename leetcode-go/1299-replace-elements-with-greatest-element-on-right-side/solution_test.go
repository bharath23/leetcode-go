package leetcode1299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "test 1",
			arr:  []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			name: "test 2",
			arr:  []int{400},
			want: []int{-1},
		},
	}

	for _, test := range tests {
		arr := make([]int, len(test.arr))
		copy(arr, test.arr)
		have := replaceElements(arr)
		assert.Equalf(t, test.want, have, "%s: replaced elements dont match", test.name)
	}
}
