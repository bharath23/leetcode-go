package leetcode0259

//import "fmt"
import "sort"

/*
Simple brute force approach, find all triplets (i,j,k) where i < j < k and
test for the condition.

Complexity:
Time complexity: O(n^3)
Space complexit: O(1)
*/
func threeSumSmallerV0(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] < target {
					count++
				}
			}
		}
	}

	return count
}

/*
Binary search approach. Time complexity would be O(n^2 logn), binary search
complexity would be O(logn) which makes complexity of twoSumSmaller O(n logn)
and outer loop would be O(n). Sorting complexity would be O(n logn) so total
time complexity would be O(n^2 logn).

Complexity:
Time complexity: O(n^2 logn)
Space complexity: O(1)
*/
func threeSumSmallerV1(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	twoSumSmaller := func(i int) {
		for j := i + 1; j < len(nums); j++ {
			l := j
			r := len(nums) - 1
			for l < r {
				m := (l + r + 1) / 2
				sum := nums[i] + nums[j] + nums[m]
				if sum < target {
					l = m
				} else {
					r = m - 1
				}
			}

			count += l - j
		}
	}

	for i := 0; i < len(nums); i++ {
		twoSumSmaller(i)
	}

	return count
}

/*
This uses the two pointer approach. Time complexity would be O(n^2). Sorting
complexity is O(n logn) and complexity is O(n^2) making total complexity
O(n^2). There is no additional space complexity.

Complexity:
Time complexity: O(n^2)
Space complexity: O(1)
*/
func threeSumSmallerV2(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	twoSumSmaller := func(i int) {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				count += r - l
				l++
			} else {
				r--
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		twoSumSmaller(i)
	}

	return count
}
