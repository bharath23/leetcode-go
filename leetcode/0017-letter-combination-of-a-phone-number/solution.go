package leetcode0017

/*
Simple iterative solution. If the maximum letters in the map is M and N is the
number of digits, total number of combinations is M^N. Time complexity would
be O(N M^N)

Complexity:
Time Complexity: O(N M^N), M is the maximum number of letters in the map and
N is the number of digits.
Space Complexity: O(N), N is the of digits
*/
func letterCombinationsV0(digits string) []string {
	digitToLetters := map[rune][]string{
		'1': []string{},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
		'0': []string{},
	}

	combinations := []string{}
	for _, d := range digits {
		letters := digitToLetters[d]
		if len(combinations) == 0 {
			combinations = append(combinations, letters...)
			continue
		}

		appendedCombinations := []string{}
		for _, c := range combinations {
			for _, l := range letters {
				appendedCombinations = append(appendedCombinations, c+l)
			}
		}

		combinations = appendedCombinations
	}

	return combinations
}

/*
Modified version of previous implementation.
*/
func letterCombinationsV1(digits string) []string {
	digitToLetters := map[rune][]string{
		'1': []string{},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
		'0': []string{},
	}

	combinations := []string{}
	for _, d := range digits {
		letters := digitToLetters[d]
		if len(combinations) == 0 {
			combinations = append(combinations, letters...)
			continue
		}

		numCombinations := len(combinations)
		for i := 0; i < numCombinations; i++ {
			c := combinations[i]
			for _, l := range letters {
				combinations = append(combinations, c+l)
			}
		}

		combinations = combinations[numCombinations:]
	}

	return combinations
}

/*
Recursive solution to generate combinations. Complexity would be similar to
iterative solution
*/
func letterCombinationsV2(digits string) []string {
	digitToLetters := map[byte][]string{
		'1': []string{},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
		'0': []string{},
	}

	combinations := []string{}
	if len(digits) == 0 {
		return combinations
	}

	var generate func(string, int)
	generate = func(str string, idx int) {
		if idx == len(digits) {
			combinations = append(combinations, str)
			return
		}

		letters := digitToLetters[digits[idx]]
		for _, l := range letters {
			generate(str+l, idx+1)
		}
	}

	generate("", 0)
	return combinations
}

/*
Iterative solution which generates each combination using math. If the maximum
letters in the map is M and N is the number of digit, total number of
combinations is M^N. Time complexity would be O(N * M^N)

Complexity:
Time complexity: O(N M^N)
Space complexity: O(N)
*/
func letterCombinationsV3(digits string) []string {
	digitToLetters := map[rune][]string{
		'1': []string{},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
		'0': []string{},
	}

	if len(digits) == 0 {
		return []string{}
	}

	numCombinations := 1
	letterRepeat := map[int]int{}
	letterCount := make([]int, len(digits))
	for i, d := range digits {
		letterCount[i] = len(digitToLetters[d])
		numCombinations *= len(digitToLetters[d])
	}

	for i := range digits {
		n := 1
		for j := i + 1; j < len(digits); j++ {
			n *= letterCount[j]
		}

		letterRepeat[i] = n
	}

	combinations := make([]string, numCombinations)
	for i := 0; i < numCombinations; i++ {
		str := ""
		for j, d := range digits {
			c := letterCount[j]
			m := letterRepeat[j]
			p := i / m
			q := p % c
			str += digitToLetters[d][q]
		}
		combinations[i] = str
	}
	return combinations
}
