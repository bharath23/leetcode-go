package leetcode0383

func canConstruct(ransomNote string, magazine string) bool {
	alphaCount := make([]int, 26)
	if len(ransomNote) > len(magazine) {
		return false
	}

	for _, m := range magazine {
		alphaCount[m-'a']++
	}

	for _, r := range ransomNote {
		alphaCount[r-'a']--
		if alphaCount[r-'a'] < 0 {
			return false
		}
	}

	return true
}
