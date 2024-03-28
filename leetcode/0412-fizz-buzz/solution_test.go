package leetcode0412

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n int
		want []string
	}{
		{
			name: "test 1",
			n: 3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "test 2",
			n: 5,
			want: []string{"1","2","Fizz","4","Buzz"},
		},
		{
			name: "test 3",
			n: 15,
			want: []string{"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"},
		},
	}

	for _, test := range tests {
		have:= fizzBuzz(test.n)
		assert.Equalf(t, test.want, have, "%s: fizzBuzz does not match", test.name)
	}
}
