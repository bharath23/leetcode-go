package leetcode0412

import "strconv"

func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		switch {
		case i % 15 == 0:
			result = append(result, "FizzBuzz")
		case i % 5 == 0:
			result = append(result, "Buzz")
		case i % 3 == 0:
			result = append(result, "Fizz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
