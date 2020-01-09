package leetcode0004

/*
Naive approach, we merge the array and then find the median and return it

Complexity:
Time complexity: O(n), we need to merge the array by visiting each element once
Space complexity: O(n), we need additional space to store the merged array
*/
func findMedianSortedArraysV0(nums1 []int, nums2 []int) float64 {
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

/*
We do not merge the array but directly find the element that will be the
median or contribute to the median.

Complexity:
Time complexity: O(log(n)), we use binary search to find the elements
Space complexity: O(1), we do not need any additional space
*/
func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	var a, b []int
	if len(nums1) > len(nums2) {
		a = nums2
		b = nums1
	} else {
		a = nums1
		b = nums2
	}

	aLen := len(a)
	bLen := len(b)
	aMin := 0
	aMax := aLen
	medianLen := (aLen + bLen + 1) / 2
	var leftHalfEnd, rightHalfStart int
	for aMin <= aMax {
		aCount := aMin + (aMax-aMin)/2
		bCount := medianLen - aCount
		if (aCount > 0) && (a[aCount-1] > b[bCount]) {
			aMax = aCount - 1
		} else if (aCount < aLen) && (b[bCount-1] > a[aCount]) {
			aMin = aCount + 1
		} else {
			if aCount == 0 {
				leftHalfEnd = b[bCount-1]
			} else if bCount == 0 {
				leftHalfEnd = a[aCount-1]
			} else {
				if a[aCount-1] > b[bCount-1] {
					leftHalfEnd = a[aCount-1]
				} else {
					leftHalfEnd = b[bCount-1]
				}
			}
			if (aLen+bLen)%2 == 1 {
				return float64(leftHalfEnd)
			}
			if aCount == aLen {
				rightHalfStart = b[bCount]
			} else if bCount == bLen {
				rightHalfStart = a[aCount]
			} else {
				if a[aCount] < b[bCount] {
					rightHalfStart = a[aCount]
				} else {
					rightHalfStart = b[bCount]
				}
			}

			return float64(leftHalfEnd+rightHalfStart) / 2.0
		}
	}

	return -1.0
}
