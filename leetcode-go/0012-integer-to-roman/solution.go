package leetcode0012

/*
Time complexity: O(1), since worst case is for 3999 which is known
Space complexity: O(1)
*/
func intToRomanV0(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i, v := range values {
		for v <= num {
			roman += symbol[i]
			num -= v
		}
	}

	return roman
}

/*
Look up table based solution
Time complexity: O(1) regardless of number it takes same time
Space complexity: O(1)
*/
func intToRomanV1(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	roman := thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + ones[(num%10)]
	return roman
}
