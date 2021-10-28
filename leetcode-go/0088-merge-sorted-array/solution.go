package leetcode0088

import "sort"

/*
Trivial solution, merge arrays and then sort them. Sorting time complexity is
O(n logn) and merging complexity is O(n). Space complexity is O(n).

Complexity:
Time complexity: O(n logn)
Space complexity: O(1), assuming sorting doesnt take any additional storage
*/

func mergeV0(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}

/*
Simple one pass solution. We first create a copy of nums1 as we will be
overwriting the array. Time complexity is O(m+n), space complexity is O(m)
additional storage

Complexity:
Time complexity: O(m+n)
Space complexity: O(m)
*/

func mergeV1(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])
	p1 := 0
	p2 := 0
	for i := 0; i < m+n; i++ {
		if (p1 >= m) || (p2 < n && nums2[p2] < nums1Copy[p1]) {
			nums1[i] = nums2[p2]
			p2++
		} else {
			nums1[i] = nums1Copy[p1]
			p1++
		}
	}
}

/*
One pass solution with no additional space requirement. Let us use the fact
that top entries of nums1 are all 0. We can use this to our advantage and
fill nums1 in descending order. Time complexity is O(m+n) and space complexity
is O(1) as no additional space is required.

Complexity:
Time complexity: O(m+n)
Space complexity: O(1)
*/

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if (p1 < 0) || (p2 >= 0 && nums2[p2] > nums1[p1]) {
			nums1[i] = nums2[p2]
			p2--
		} else {
			nums1[i] = nums1[p1]
			p1--
		}
	}
}
