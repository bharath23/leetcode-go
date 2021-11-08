package leetcode0238

/*
A solution that uses a temporary storage. First copy all non-zero elements to
the temporary array. Zero out the rest of the array and then copy back to nums.
Time complexity is O(n) for copying the values and O(n) to copy back to
original array. Time complexity is O(n), space complexity is O(n).

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func moveZeroesV0(nums []int) {
	temp := make([]int, len(nums))
	i := 0
	for _, v := range nums {
		if v != 0 {
			temp[i] = v
			i++
		}
	}

	for i < len(nums) {
		temp[i] = 0
		i++
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}
}

/*
A single pass, two pointer solution. We have a reader and writer pointers.
If the value of the element for the current read pointer is non-zero write to
current write pointer. Once we finish writing all non-zero values fill the
rest of the array with zero. Time complexity is O(n) for copying non-zero
values is O(n) filling with rest of the array with zerios is O(n). Time
complexity O(n), space complexity O(1) as no additional space is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func moveZeroesV1(nums []int) {
	w := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[w] = nums[r]
			w++
		}
	}

	for w < len(nums) {
		nums[w] = 0
		w++
	}
}

/*
Single pass, two pointer solution. We have a reader and writer pointers. If
the value of element for the current reader is non-zero we swap with the value
pointed to the write pointer. We go through the array once so time complexity
is O(n), space complexity is O(1) as no additional space is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func moveZeroesV2(nums []int) {
	w := 0
	for r, v := range nums {
		if v != 0 {
			t := nums[w]
			nums[w] = v
			nums[r] = t
			w++
		}
	}
}
