package leetcode2235

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "test 1",
			num1: 12,
			num2: 5,
			want: 17,
		},
		{
			name: "test 2",
			num1: -10,
			num2: 4,
			want: -6,
		},
	}

	for _, test := range tests {
		have := sum(test.num1, test.num2)
		assert.Equalf(t, test.want, have, "%s: sum does not match", test.name)
	}
}
