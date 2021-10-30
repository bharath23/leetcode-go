package leetcode0010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	s    string
	p    string
	want bool
}{
	{
		name: "test 1",
		s:    "aa",
		p:    "a",
		want: false,
	},
	{
		name: "test 2",
		s:    "aa",
		p:    "a*",
		want: true,
	},
	{
		name: "test 3",
		s:    "ab",
		p:    ".*",
		want: true,
	},
	{
		name: "test 4",
		s:    "aab",
		p:    "c*a*b",
		want: true,
	},
	{
		name: "test 5",
		s:    "mississippi",
		p:    "mis*is*p*.",
		want: false,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := isMatchV0(test.s, test.p)
		assert.Equalf(t, test.want, have, "%s: regex matching failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := isMatchV1(test.s, test.p)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
