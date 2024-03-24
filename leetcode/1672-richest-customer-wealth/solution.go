package leetcode1672

/*
Time complexity: O(n^2)
Space complexity: O(1)
*/
func maximumWealth(accounts [][]int) int {
	if len(accounts) == 0 {
		return 0
	}

	max := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
