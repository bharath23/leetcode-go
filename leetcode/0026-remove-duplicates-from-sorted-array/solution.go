package leetcode0026

/*
Simple one pass solution using two pointers. Slower pointer is used to write
while the faster pointer is used to read. Time complexity is O(n) space
complexity is O(1) as no additional space is needed.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
