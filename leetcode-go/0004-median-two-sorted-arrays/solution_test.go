package leetcode0004

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

var tests = []struct {
	nums1 []int
	nums2 []int
	want  float64
}{
	{
		nums1: []int{1, 3},
		nums2: []int{2},
		want:  2.0,
	},
	{
		nums1: []int{1, 2},
		nums2: []int{3, 4},
		want:  2.5,
	},
	{
		nums1: []int{1, 2},
		nums2: []int{1, 2, 3},
		want:  2.0,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		have := findMedianSortedArraysV0(test.nums1, test.nums2)
		assert.Equal(t, test.want, have, "want: %f, have: %f", test.want, have)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		have := findMedianSortedArraysV1(test.nums1, test.nums2)
		assert.Equal(t, test.want, have, "want: %f, have: %f", test.want, have)
	}
}
