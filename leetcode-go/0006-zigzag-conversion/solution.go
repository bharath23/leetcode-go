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
