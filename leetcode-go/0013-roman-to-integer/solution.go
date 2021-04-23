package leetcode0013

/*
Simple one pass solution left to right
Time complexity: O(1), since worst case is for 3999
Space complexity: O(1), since storage is same regardless of string
*/

func romanToIntV0(s string) int {
	symbolToValue := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0
	prev := 0
	for _, c := range s {
		cur := symbolToValue[c]
		if prev < cur {
			num = num - 2*prev
		}
		num += cur
		prev = cur
	}

	return num
}

/*
One pass solution left to right
Time complexity: O(1), since worst case is for 3999
Space complexity: O(1), storage is same regardless of string
*/

func romanToIntV1(s string) int {
	symbolToValue := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	num := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			val, ok := symbolToValue[string(s[i:i+2])]
			if ok {
				num += val
				i = i + 2
				continue
			}
		}

		val := symbolToValue[string(s[i])]
		num += val
		i = i + 1
	}

	return num
}
