package leetcode0007

func reverse(x int) int {
	maxInt := 1<<31 - 1
	maxRemainder := maxInt % 10
	maxQuotient := maxInt / 10
	isNeg := 1
	if x < 0 {
		isNeg = -1
		x = isNeg * x
	}

	var reverseX = 0
	for x > 0 {
		r := x % 10
		x = x / 10
		if (reverseX > maxQuotient) || (reverseX == maxQuotient && r > maxRemainder) {
			return 0
		}

		reverseX = 10*reverseX + r
	}

	return isNeg * reverseX
}
