package leetcode0167

/*
We use the property that all the numbers are sorted.

Time Complexity: O(n), we traverse the list once
Space Complexity: O(1), we do not need any additional space
*/

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum > target {
			r--
		} else {
			l++
		}
	}

	return []int{-1, -1}
}
