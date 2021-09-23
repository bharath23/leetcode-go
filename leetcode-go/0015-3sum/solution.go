package leetcode0015

import "sort"

/*
We sort the array as then we can use twoSum for sorted array with complexity
O(n). Since we do this for each element total time complexity would be O(n^2).
Sorting complexity is O(nlogn).

Complexity:
Time complexity: O(n^2), actual complexity is O(nlogn) + O(n^2)
Space complexity: O(1), we ignore space complexity for sorting
*/

func threeSumV0(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	twoSum := func(i int) {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for (j < k) && (nums[j] == nums[j-1]) {
					j++
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSum(i)
		}
	}

	return result
}

/*
We sort the array as then we use twoSum with hashmap with complexity O(n).
Since we do this for each element total time complexity would be O(n^2).
Sorting complexity is O(nlogn)

Complexity:
Time complexity: O(n^2), actaul complexity is O(nlogn) + O(n^2)
Space complexity: O(n), space is needed to store the hash map
*/
func threeSumV1(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	twoSum := func(i int) {
		isPresent := make(map[int]bool, len(nums))
		for j := i + 1; j < len(nums); j++ {
			complement := -nums[i] - nums[j]
			if isPresent[complement] {
				result = append(result, []int{nums[i], complement, nums[j]})
				for (j+1 < len(nums)) && (nums[j] == nums[j+1]) {
					j++
				}
			}
			isPresent[nums[j]] = true
		}
	}

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSum(i)
		}
	}

	return result
}

/*
We run two loops so O(n2^2) time complexity. We need additional space for the
hash maps for duplicate and complement which is O(n).

Complexity:
Time complexity: O(n^2)
Space complexity: O(n)
*/

func threeSumV2(nums []int) [][]int {
	result := [][]int{}
	inResult := make(map[[3]int]bool)
	duplicate := make(map[int]bool, len(nums))
	isComplement := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		isDuplicate := duplicate[nums[i]]
		if !isDuplicate {
			duplicate[nums[i]] = true
			for j := i + 1; j < len(nums); j++ {
				complement := -nums[i] - nums[j]
				k, ok := isComplement[complement]
				if ok && k == i {
					set := []int{nums[i], complement, nums[j]}
					sort.Ints(set)
					key := [3]int{}
					copy(key[:], set)
					if !inResult[key] {
						inResult[key] = true
						result = append(result, set)
					}
				}
				isComplement[nums[j]] = i
			}
		}
	}

	return result
}
