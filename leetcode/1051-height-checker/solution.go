package leetcode1051

import "fmt"
import "sort"

/*
Simple one pass solution. Time complexity is O(n logn) + O(n) + O(n), which is
O(n). Space complexity is O(n) for the copy of the array.

Complexity:
Time complexity: O(n logn)
Space complexity: O(n)
*/

func heightCheckerV0(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0
	for i, v := range expected {
		if heights[i] != v {
			count++
		}
	}

	return count
}

const minHeight = 1
const maxHeight = 100

/*
Simple one pass solution. Time complexity is O(n) + O(n), which is O(n). Space
complexity is O(n) for the frequency.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func heightCheckerV1(heights []int) int {
	frequency := make([]int, maxHeight+1)
	for _, v := range heights {
		frequency[v]++
	}

	count := 0
	k := 0
	fmt.Printf("frequency: %#v\n", frequency)
	for i := minHeight; i <= maxHeight; i++ {
		for j := frequency[i]; j > 0; j-- {
			if heights[k] != i {
				count++
			}
			k++
		}
	}

	return count
}
