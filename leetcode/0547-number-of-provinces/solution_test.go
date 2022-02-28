package leetcode0547

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name        string
	isConnected [][]int
	want        int
}{
	{
		name:        "test1",
		isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		want:        2,
	},
	{
		name:        "test2",
		isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		want:        3,
	},
	{
		name:        "test3",
		isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
		want:        1,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := findCircleNumV0(test.isConnected)
		assert.Equalf(t, test.want, have, "%s: number of province mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := findCircleNumV1(test.isConnected)
		assert.Equalf(t, test.want, have, "%s: number of province mismatch", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		have := findCircleNumV2(test.isConnected)
		assert.Equalf(t, test.want, have, "%s: number of province mismatch", test.name)
	}
}
