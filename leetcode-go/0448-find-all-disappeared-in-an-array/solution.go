package leetcode0448

/*
Simple two pass solution. First pass make a set of all unique numbers present in
the array. Second pass check if each index exists in the unique set. If not add
to the result set. Time complexity is O(n), space complexity is also O(n).

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func findDisappearedNumbersV0(nums []int) []int {
	result := []int{}
	numbersPresent := make(map[int]bool)
	for _, n := range nums {
		numbersPresent[n] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !numbersPresent[i] {
			result = append(result, i)
		}
	}

	return result
}

/*
Another two pass solution. In the first pass mark the value stored in array as
negative for the index pointed by the value in each element of the array. In
then second pass any index that stores a positive value is missing, add it to
the result set. Time complexity is O(n), space complexity is O(1).

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func findDisappearedNumbersV1(nums []int) []int {
	for _, n := range nums {
		nums[abs(n)-1] = -1 * abs(nums[abs(n)-1])
	}

	result := []int{}
	for i, n := range nums {
		if n > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
