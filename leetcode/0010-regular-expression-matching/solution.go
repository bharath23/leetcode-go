package leetcode0010

/*
Time and space complexity is hard to express.
*/
func isMatchV0(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	match := len(s) != 0
	match = match && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return isMatchV0(s, p[2:]) || (match && isMatchV0(s[1:], p))
	} else {
		return match && isMatchV0(s[1:], p[1:])
	}
}

/*
Every call to dynamicIsMatch(i, j) is really evaluated only once. If length of
string is S and length of pattern is P
Time complexity: O(T*P)
Space complexity: O(T*P) boolean entries for memoization cache
*/
func isMatchV1(s string, p string) bool {
	type key struct {
		i int
		j int
	}
	memo := make(map[key]bool)
	var dynamicIsMatch func(int, int) bool
	dynamicIsMatch = func(i, j int) bool {
		k := key{i: i, j: j}
		m, ok := memo[k]
		if ok {
			return m
		}

		if len(p) == j {
			m = len(s) == i
		} else {
			m = len(s) > i
			m = m && (s[i] == p[j] || p[j] == '.')
			if len(p) > j+1 && p[j+1] == '*' {
				m = dynamicIsMatch(i, j+2) || (m && dynamicIsMatch(i+1, j))
			} else {
				m = m && dynamicIsMatch(i+1, j+1)
			}
		}

		memo[k] = m
		return m
	}

	return dynamicIsMatch(0, 0)
}
