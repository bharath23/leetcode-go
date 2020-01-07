package leetcode0001

/*
One-pass solution:
While iterating insert the element into a map, check if the element's
complement already exists. If it exists return the indices.

Complexity:
Time complexity: O(n), we traverse the list once and map lookup is O(1)
Space complexity: O(n), we need extra space to store the map
*/

func twoSums(nums []int, target int) []int {
	numToIdxMap := make(map[int]int, len(nums))
	for i, n := range nums {
		complement := target - n
		if idx, ok := numToIdxMap[complement]; ok {
			return []int{idx, i}
		}

		numToIdxMap[n] = i
	}

	return []int{-1, -1}
}
