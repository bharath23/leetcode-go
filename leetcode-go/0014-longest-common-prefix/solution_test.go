package leetcode0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	strs []string
	want string
}{
	{
		name: "test 1",
		strs: []string{
			"flower",
			"flow",
			"flight",
		},
		want: "fl",
	},
	{
		name: "test 2",
		strs: []string{
			"dog",
			"racecar",
			"car",
		},
		want: "",
	},
	{
		name: "test 3",
		strs: []string{
			"",
			"c",
		},
		want: "",
	},
	{
		name: "test 4",
		strs: []string{
			"",
		},
		want: "",
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := longestCommonPrefixV0(test.strs)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := longestCommonPrefixV1(test.strs)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		have := longestCommonPrefixV2(test.strs)
		assert.Equalf(t, test.want, have, "%s failed", test.name)
	}
}
