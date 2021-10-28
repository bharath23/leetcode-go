package leetcode0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	nums1 []int
	m     int
	nums2 []int
	n     int
	want  []int
}{
	{
		name:  "test 1",
		nums1: []int{1, 2, 3, 0, 0, 0},
		m:     3,
		nums2: []int{2, 5, 6},
		n:     3,
		want:  []int{1, 2, 2, 3, 5, 6},
	},
	{
		name:  "test 2",
		nums1: []int{1},
		m:     1,
		nums2: []int{},
		n:     0,
		want:  []int{1},
	},
	{
		name:  "test 3",
		nums1: []int{0},
		m:     0,
		nums2: []int{1},
		n:     1,
		want:  []int{1},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		nums1 := make([]int, len(test.nums1))
		copy(nums1, test.nums1)
		mergeV0(nums1, test.m, test.nums2, test.n)
		assert.Equalf(t, test.want, nums1, "%s failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		nums1 := make([]int, len(test.nums1))
		copy(nums1, test.nums1)
		mergeV1(nums1, test.m, test.nums2, test.n)
		assert.Equalf(t, test.want, nums1, "%s failed", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		nums1 := make([]int, len(test.nums1))
		copy(nums1, test.nums1)
		mergeV2(nums1, test.m, test.nums2, test.n)
		assert.Equalf(t, test.want, nums1, "%s failed", test.name)
	}
}
