package leetcode1480

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := []int{}
	result = append(result, nums[0])
	for i, n := range nums[1:] {
		result = append(result, result[i]+n)
	}

	return result
}
