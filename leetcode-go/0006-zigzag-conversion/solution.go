package leetcode0006

import (
	"strings"
)

/*
Simple solution, iterating throught the string and appending each row with the
appropriate letter.
Complexity:

Time complexity: O(n) the string is just traversed once
Space complexity: O(n) We have to store all the characters of the string in
	another array of strings
*/
func convertV0(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows []string
	if len(s) < numRows {
		rows = make([]string, len(s))
	} else {
		rows = make([]string, numRows)
	}

	goingDown := true
	idx := 0
	for _, c := range s {
		rows[idx] += string(c)
		if goingDown {
			idx++
			if idx == numRows {
				goingDown = false
				idx = idx - 2
			}
		} else {
			idx--
			if idx == -1 {
				goingDown = true
				idx = idx + 2
			}
		}
	}

	return strings.Join(rows, "")
}

/*
Visit characters as they would appear in a row i.e row 0, row 1 and so on.

A cycle starts at row 0 and ends at row 1

Cycle length: 2n - 2
For kth cycle letters are
row 0: k*(cycle length)
row i: i + k*(cycle length), (k+1)*(cycle length) - i
row n-1: n-1 + k*(cycle length)

Complexity:
Time complexity: O(n) the string is just traversed once
Space complexity: O(1) we do not need additional storage
*/
func convertV1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var retval strings.Builder
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j = j + cycleLen {
			retval.WriteByte(s[j+i])
			if (i == 0) || (i == numRows-1) {
				continue
			}

			if (j + cycleLen - i) < len(s) {
				retval.WriteByte(s[j+cycleLen-i])
			}
		}
	}

	return retval.String()
}
