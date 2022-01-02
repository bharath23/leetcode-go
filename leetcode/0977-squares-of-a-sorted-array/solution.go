package leetcode0977

import "sort"

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

/*
Simple one pass solution. We calculate the squares and then sort the result.
Time complexity would be O(n) for calculating squares and in case of sorting
we can assume it to be O(n logn). Total time complexity would be O(n logn).

Complexity
Time complexity: O(n logn)
Space complexity: O(n), assuming sorting doesnt use additional storage
*/

func sortedSquaresV0(nums []int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = n * n
	}

	sort.Ints(result)
	return result
}

/*
Two pointer solution. We use the fact that the original array is sorted in
non-decreasing order. We compare the absolute values between smallest and
largest number and calculate the square and store it in descending order.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func sortedSquaresV1(nums []int) []int {
	result := make([]int, len(nums))
	l := 0
	r := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		var n int
		if abs(nums[l]) > abs(nums[r]) {
			n = nums[l]
			l++
		} else {
			n = nums[r]
			r--
		}
		result[i] = n * n
	}

	return result
}
