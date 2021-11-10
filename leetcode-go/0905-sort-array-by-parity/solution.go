package leetcode0905

import "sort"

/*
Simple sorting based solution. Sort the array such that even numbers are
considered smaller than odd numbers. Time complexity is O(n logn) for sorting
and space complexity depends on the algorithm.

Complexity:
Time complexity: O(n logn)
Space complexity: O(1) or O(n)
*/

func sortArrayByParityV0(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return (nums[i] % 2) < (nums[j] % 2) })
	return nums
}

/*
Two pass solution. We scan the array first and store all the even elements in
the result array, followed by scanning the array again and storing the odd
elements. Time complexity is O(n) + O(n) which is equivalent to O(n). Space
complexity is O(n) to store the result.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func sortArrayByParityV1(nums []int) []int {
	result := make([]int, len(nums))
	i := 0
	for _, v := range nums {
		if v%2 == 0 {
			result[i] = v
			i++
		}
	}

	for _, v := range nums {
		if v%2 == 1 {
			result[i] = v
			i++
		}
	}

	return result
}

/*
One pass, two pointer solution. We scan the array front and back. We compare
the elements at the current pointer and swap them if front pointer points to
on odd number. We move the pointers accordingly and keep checking. Time
complexity is O(n) since we scan the array just once. Space complexity is O(1)
as we do not need any additional storage.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func sortArrayByParityV2(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 > nums[j]%2 {
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
		}

		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}

	return nums
}
