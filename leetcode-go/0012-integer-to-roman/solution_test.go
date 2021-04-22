package leetcode0012

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

var tests = []struct {
	name string
	num  int
	want string
}{
	{
		name: "test 1",
		num:  3,
		want: "III",
	},
	{
		name: "test 2",
		num:  4,
		want: "IV",
	},
	{
		name: "test 3",
		num:  9,
		want: "IX",
	},
	{
		name: "test 4",
		num:  58,
		want: "LVIII",
	},
	{
		name: "test 5",
		num:  1994,
		want: "MCMXCIV",
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := intToRomanV0(test.num)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := intToRomanV1(test.num)
		assert.Equal(t, test.want, have, "%s failed", test.name)
	}
}
