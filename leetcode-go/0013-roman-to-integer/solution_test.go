package leetcode0013

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

var tests = []struct {
	name string
	s    string
	want int
}{
	{
		name: "test 1",
		s:    "III",
		want: 3,
	},
	{
		name: "test 2",
		s:    "IV",
		want: 4,
	},
	{
		name: "test 3",
		s:    "IX",
		want: 9,
	},
	{
		name: "test 4",
		s:    "LVIII",
		want: 58,
	},
	{
		name: "test 5",
		s:    "MCMXCIV",
		want: 1994,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := romanToIntV0(test.s)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := romanToIntV1(test.s)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}
