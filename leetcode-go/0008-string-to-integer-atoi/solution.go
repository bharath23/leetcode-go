package leetcode0008

import (
	"strings"
	"unicode"
)

/*
Traverse the string character by character.

Complexity:
Time complexity: O(n) the string is traveresed once
Space complexity: O(1)
*/

func atoi(str string) int {
	retval := 0
	isNeg := false
	minus := '-'
	plus := '+'
	minInt := 1 << 31
	maxInt := minInt - 1
	maxQuotient := maxInt / 10
	maxRemainder := maxInt % 10
	minRemainder := minInt % 10
	str = strings.TrimLeft(str, " ")
	for i, c := range str {
		if i == 0 {
			if c == minus {
				isNeg = true
				continue
			}

			if c == plus {
				continue
			}
		}

		if !unicode.IsDigit(c) {
			break
		}

		v := int(c - '0')
		if retval < maxQuotient {
			retval = 10*retval + v
			continue
		}

		if retval > maxQuotient {
			if isNeg {
				return -1 * minInt
			} else {
				return maxInt
			}
		}

		if isNeg {
			if v > minRemainder {
				return -1 * minInt
			} else {
				retval = 10*retval + v
				continue
			}
		} else {
			if v > maxRemainder {
				return maxInt
			} else {
				retval = 10*retval + v
				continue
			}
		}
	}

	if isNeg {
		return -1 * retval
	} else {
		return retval
	}
}
