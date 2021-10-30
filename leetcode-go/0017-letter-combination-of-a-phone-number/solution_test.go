package leetcode0017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name   string
	digits string
	want   []string
}{
	{
		name:   "test 1",
		digits: "23",
		want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		name:   "test 2",
		digits: "",
		want:   []string{},
	},
	{
		name:   "test 3",
		digits: "2",
		want:   []string{"a", "b", "c"},
	},
	{
		name:   "test 4",
		digits: "234",
		want: []string{
			"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
			"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
			"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
		},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := letterCombinationsV0(test.digits)
		assert.ElementsMatchf(t, test.want, have, "%s: possible combination generated does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := letterCombinationsV1(test.digits)
		assert.ElementsMatchf(t, test.want, have, "%s: possible combination generated does not match", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		have := letterCombinationsV2(test.digits)
		assert.ElementsMatchf(t, test.want, have, "%s failed", test.name)
	}
}

func TestSolutionV3(t *testing.T) {
	for _, test := range tests {
		have := letterCombinationsV3(test.digits)
		assert.ElementsMatchf(t, test.want, have, "%s failed", test.name)
	}
}
