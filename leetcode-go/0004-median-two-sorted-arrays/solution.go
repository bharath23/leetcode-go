package leetcode0004

func findMedianSortedArraysv0(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	k := 0
	for (i < len(nums1)) && (j < len(nums2)) {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		nums[k] = nums1[i]
		k++
		i++
	}

	for j < len(nums2) {
		nums[k] = nums2[j]
		k++
		j++
	}

	var median float64
	if len(nums)%2 == 0 {
		median = float64(nums[len(nums)/2-1] + nums[len(nums)/2])
		median = median / 2.0
	} else {
		median = float64(nums[len(nums)/2])
	}

	return median
}
