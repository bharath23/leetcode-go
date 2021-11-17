package leetcode0414

import (
	"math"
	"sort"
)

/*
This is a simple solution based on sorting. Sort the numbers in reverse and
find the third maximum. Time complexity is O(n logn) for sorting and O(n) for
finding the maximum. Total time complexity is O(n logn). Space complexity is
O(1) as no additional storage is required.

Complexity:
Time complexity: O(n logn)
Space complexity: O(1)
*/

func thirdMaxV0(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	max := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			max = nums[i]
			j++
		}

		if j == 3 {
			break
		}
	}

	if j != 3 {
		max = nums[0]
	}

	return max
}

/*
4 pass solution. First pass, create the set of unique numbers. Second pass
find the maximum in the unique numbers. Third pass find the second maximum
and forth pass find the third maximum. Time complexity is 4 * O(n) which is
O(n). Space complexity is O(n) for map and O(n) for the unique numbers, which
is still O(n).

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func thirdMaxV1(nums []int) int {
	unique := make(map[int]bool)
	for _, n := range nums {
		unique[n] = true
	}

	uniqueNums := []int{}
	for n := range unique {
		uniqueNums = append(uniqueNums, n)
	}

	m1 := math.MinInt
	for _, n := range uniqueNums {
		if n > m1 {
			m1 = n
		}
	}

	if len(uniqueNums) < 3 {
		return m1
	}

	m2 := math.MinInt
	for _, n := range uniqueNums {
		if n > m2 && n != m1 {
			m2 = n
		}
	}

	m3 := math.MinInt
	for _, n := range uniqueNums {
		if n > m3 && n != m1 && n != m2 {
			m3 = n
		}
	}

	return m3
}

/*
Three pass solution. In each pass we find a maximum that hasnt been seen
earlier. Then we return the smallest of the three maximums. Time complexity
is 3 * O(n) which is still O(n). Space complexity is O(1) as amount of
additional memory required is always the same regardless of n.

Complexity:
Time complexity: O(n)
Space complexity: O(1)

*/

func thirdMaxV2(nums []int) int {
	maximums := []int{}
	isSeenMax := func(n int) bool {
		for _, m := range maximums {
			if n == m {
				return true
			}
		}

		return false
	}

	findNextMaximum := func() int {
		max := math.MinInt
		for _, n := range nums {
			if n > max && !isSeenMax(n) {
				max = n
			}
		}

		return max
	}

	for i := 0; i < 3; i++ {
		max := findNextMaximum()
		if max == math.MinInt {
			break
		}
		maximums = append(maximums, max)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(maximums)))
	if len(maximums) < 3 {
		return maximums[0]
	}

	return maximums[2]
}

/*
One pass solution. We try to find the three maximum numbers. At the end we
try to find the three maximum numbers. At the end we return the third maximum,
Time complexity is is O(n) as we go through the list once. Space complexity is
O(1).

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func thirdMaxV3(nums []int) int {
	f, s, t := math.MinInt, math.MinInt, math.MinInt
	for _, n := range nums {
		if n == f || n == s || n == t {
			continue
		}
		switch {
		case n > f:
			f, s, t = n, f, s
		case n > s:
			s, t = n, s
		case n > t:
			t = n
		}
	}

	if t == math.MinInt {
		return f
	}

	return t
}
