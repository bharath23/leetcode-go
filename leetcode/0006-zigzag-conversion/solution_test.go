package leetcode0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name    string
	s       string
	numRows int
	want    string
}{
	{
		name:    "test 1",
		s:       "AB",
		numRows: 1,
		want:    "AB",
	},
	{
		name:    "test 2",
		s:       "PAYPALISHIRING",
		numRows: 3,
		want:    "PAHNAPLSIIGYIR",
	},
	{
		name:    "test 3",
		s:       "PAYPALISHIRING",
		numRows: 4,
		want:    "PINALSIGYAHRPI",
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := convertV0(test.s, test.numRows)
		assert.Equalf(t, test.want, have, "%s: generated pattern does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := convertV1(test.s, test.numRows)
		assert.Equalf(t, test.want, have, "%s: generated pattern does not match", test.name)
	}
}
