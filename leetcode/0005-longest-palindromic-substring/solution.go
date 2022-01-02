package leetcode0005

/*
Palindrome are mirror around the center. So we expand from the center to
generate longest palindrome.

Complexity:
Time complexity: O(n2) There are 2n - 1 possible centers. Each needs another n
	comparison to find the palindrome string
Space complexity: O(n) There is no additional space needed
*/

func longestPalindrome(s string) string {
	retvalLen := 0
	retval := ""
	for i := 0; i < len(s); i++ {
		str := findLongestPalindrome(s, i, i)
		if len(str) > retvalLen {
			retval = str
			retvalLen = len(str)
		}

		str = findLongestPalindrome(s, i, i+1)
		if len(str) > retvalLen {
			retval = str
			retvalLen = len(str)
		}
	}

	return retval
}

func findLongestPalindrome(s string, l, r int) string {
	for (l >= 0) && (r < len(s)) && (s[l] == s[r]) {
		l--
		r++
	}

	return s[l+1 : r]
}
