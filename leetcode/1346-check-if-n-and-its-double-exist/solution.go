package leetcode1246

/*
Simple one pass solution with a map to check existence. For each value check
if the double or exact half of it already exists in the map. Time complexity is
O(n) as we scan the entire array once. Space complexity is O(n) as we need to
store the map.

Complexity:
Time complexity: O(n)
Space complexity: O(n), space to store
*/

func checkIfExist(arr []int) bool {
	numMap := make(map[int]bool, len(arr))
	for _, v := range arr {
		if numMap[v*2] || (numMap[v/2] && v%2 == 0) {
			return true
		}

		numMap[v] = true
	}

	return false
}
