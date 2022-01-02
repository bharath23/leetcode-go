package leetcode0027

//import "fmt"

/*
One pass solution. Time complexity would be O(n) since we visit all the nodes
once. Since we do everything in place space complexity is O(1)

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func removeElementV0(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}

	return n
}

/*
One pass solution. Since, we do not have to maintain order of the elements
we can copy the elements from the back to replace the elements which are
equal. This reduces the number of copies. Time complexity O(n) and space
complexity is still O(1) as not additional space is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func removeElementV1(nums []int, val int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}
