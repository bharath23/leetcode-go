package leetcode0485

/*
Simple one pass solution, count number of consecutive ones and remember the
maximum.

Complexity:
Time Complexity: O(N)
Space Complexity: O(1)
*/

func findMaxConsecutiveOnes(nums []int) int {
	max_count := 0
	count := 0
	for _, n := range nums {
		if n == 1 {
			count++
		} else {
			if (count > 0) && (max_count < count) {
				max_count = count
			}
			count = 0
		}
	}

	if max_count < count {
		max_count = count
	}

	return max_count
}
