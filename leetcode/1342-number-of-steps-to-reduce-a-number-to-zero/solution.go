package leetcode1342

func numberOfStepsV0(num int) int {
	steps := 0
	for num > 0 {
		steps++
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
	}

	return steps
}

func numberOfStepsV1(num int) int {
	steps := 0
	for num > 0 {
		steps++
		if num&1 == 0 {
			num = num >> 1
		} else {
			num = num - 1
		}
	}

	return steps
}
