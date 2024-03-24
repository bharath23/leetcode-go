package leetcode1672

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{
			name:     "test 1",
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			want:     6,
		},
		{
			name:     "test 2",
			accounts: [][]int{{1, 5}, {7, 3}, {3, 5}},
			want:     10,
		},
		{
			name:     "test 3",
			accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			want:     17,
		},
	}

	for _, test := range tests {
		have := maximumWealth(test.accounts)
		assert.Equalf(t, test.want, have, "%s: maximumWealth does not match", test.name)
	}
}
