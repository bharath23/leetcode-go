package leetcode1295

/*
Simple one pass solution, find the number of digits for each number, if it
even increment count.

Complexity:
Time Complexity: O(N)
Space Complexity

*/
func findNumbers(nums []int) int {
	isEvenNumberOfDigits := func(n int) bool {
		digits := 0
		for n > 0 {
			digits++
			n = n / 10
		}

		return (digits % 2) == 0
	}

	count := 0
	for _, n := range nums {
		if isEvenNumberOfDigits(n) {
			count++
		}
	}

	return count
}
