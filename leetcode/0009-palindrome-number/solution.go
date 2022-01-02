package leetcode0009

import (
	"strconv"
)

/*
Convert the number to string and verify if the number is palindrome
Complexity:
Time complexity: O(n), we need to traverse the string
Space complexity: we need extra space to store string
*/
func isPalindromeV0(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.FormatInt(int64(x), 10)
	for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
		if str[l] != str[r] {
			return false
		}
	}

	return true
}

/*
Reverse half of the number and compare
Time complexity: O(n), we need to partially reverse the number
Space complexity: O(1), we need to store the partially reveresed number
*/
func isPalindromeV1(x int) bool {
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}

	rev := 0
	for x > rev {
		r := x % 10
		rev = 10*rev + r
		x = x / 10
	}

	return x == rev || x == rev/10
}
