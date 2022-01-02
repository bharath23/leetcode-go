package leetcode0003

/*
Remember if we have seen a character in the substring and remember the length
of the longest substring we have seen till then. Do this for every character
from the start

Complexity:
Time complexity: O(n), we might traverse the entire list twice for i and j.
Space complexity: O(n), we need to remember the characters we have seen.
*/

func lengthOfLongestSubstringV0(s string) int {
	retval := 0
	charMap := make(map[uint8]bool)
	i := 0
	j := 0
	for (i < len(s)) && (j < len(s)) {
		isPresent, ok := charMap[s[j]]
		if !ok || !isPresent {
			charMap[s[j]] = true
			l := j - i + 1
			if retval < l {
				retval = l
			}
			j++
		} else {
			charMap[s[i]] = false
			i++
		}
	}

	return retval
}

/*
Remember where have seen a character in the past, remember the substring and
its lenght. When we find a character that we have seen before, then start the
new substring from the character after the one that repeated.

Complexity:
Time complexity: O(n), we might traverse the entire list twice for i and j.
Space complexity: O(n), we need to remember the characters we have seen.
*/

func lengthOfLongestSubstringV1(s string) int {
	retval := 0
	charToIdx := make(map[uint8]int)
	i := 0
	for j := 0; j < len(s); j++ {
		idx, ok := charToIdx[s[j]]
		if ok {
			if i < idx {
				i = idx
			}
		}
		charToIdx[s[j]] = j + 1
		l := j - i + 1
		if retval < l {
			retval = l
		}
	}

	return retval
}
