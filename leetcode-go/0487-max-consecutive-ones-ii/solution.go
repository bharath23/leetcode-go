package leetcode0487

/*
Two pass solution. We scan the array till we encounter a second zero. Time
complexity for this O(n ^ 2), space complexity is O(1).

Complexity:
Time complexity: O(n^2)
Space complexity: O(1), no additional space required
*/

func findMaxConsecutiveOnesV0(nums []int) int {
	lenNums := len(nums)
	maxConsecutiveOnes := 0
	i := 0
	j := i
	for ; i < lenNums; i++ {
		j = i

		numZeros := 0
		for ; j < lenNums; j++ {
			if nums[j] == 0 {
				numZeros++
			}
			if numZeros > 1 {
				consecutiveOnes := j - i
				if consecutiveOnes > maxConsecutiveOnes {
					maxConsecutiveOnes = consecutiveOnes
				}
				break
			}
		}

		consecutiveOnes := j - i
		if consecutiveOnes > maxConsecutiveOnes {
			maxConsecutiveOnes = consecutiveOnes
		}
	}

	return maxConsecutiveOnes
}

/*
Simple one pass solution using two pointers. Two pointers track the largest
possible window. The window becomes invalid the instant we have two or more
zero values. We adjust the left pointer till we have a valid window. The time
complexity is O(n) and space complexity is O(1) as no additional storage is nee
*/

func findMaxConsecutiveOnesV1(nums []int) int {
	lenNums := len(nums)
	maxConsecutiveOnes := 0
	l := 0
	r := 0
	numZeros := 0
	for r < lenNums {
		if nums[r] == 0 {
			numZeros++
		}

		if numZeros > 1 {
			if nums[l] == 0 {
				numZeros--
			}
			l++
		}
		consecutiveOnes := r - l + 1
		if consecutiveOnes > maxConsecutiveOnes {
			maxConsecutiveOnes = consecutiveOnes
		}
		r++
	}

	return maxConsecutiveOnes
}
