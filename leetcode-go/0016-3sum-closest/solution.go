package leetcode0016

import (
	"math"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

/*
This solution uses two pointer approach. Time complexity would be O(n^2).
Sorting complexity is O(n logn) and actual complexity for finding sum is
O(n^2). There is no additional space requirement.

Complexity:
Time complexity: O(n^2)
Space complexity: O(1)
*/

func threeSumClosestV0(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt
	for i := 0; i < len(nums) && diff != 0; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}

	return target - diff
}

/*
Binary search approach. We use binary search to find a value in the array that
is larger than the complement. We compare the diff between the upper bound and
previous lower value and complement. Time complexity of O(n^2 logn), binary
search has a complexity of O(logn). With two outer loops it would be
O(n^2 logn).

Complexity:
Time complexity: O(n^2 logn)
Space complexity: O(1)
*/

func threeSumClosestV1(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt
	upperBound := func(l, r, v int) int {
		for l < r {
			m := (l + r) / 2
			if nums[m] < v {
				l = m + 1
			} else {
				r = m
			}
		}

		return r
	}

	for i := 0; i < len(nums) && diff != 0; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			complement := target - nums[i] - nums[j]
			k := upperBound(j+1, len(nums)-1, complement)
			if k < len(nums) && abs(complement-nums[k]) < abs(diff) {
				diff = complement - nums[k]
			}
			if j < k-1 && abs(complement-nums[k-1]) < abs(diff) {
				diff = complement - nums[k-1]
			}
		}
	}

	return target - diff
}
