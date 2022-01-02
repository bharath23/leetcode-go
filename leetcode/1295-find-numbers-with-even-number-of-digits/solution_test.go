package leetcode1295

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			name: "test 2",
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
	}

	for _, test := range tests {
		have := findNumbers(test.nums)
		assert.Equalf(t, test.want, have, "%s: number of event digit numbers do not match", test.name)
	}
}
