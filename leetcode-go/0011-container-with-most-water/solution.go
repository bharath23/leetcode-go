package leetcode0011

/*
Simple brute for approach.
Time complexity: O(n^2)
Space complexity: O(1), constant space required
*/
func maxAreaV0(height []int) int {
	maxArea := 0
	numHeight := len(height)
	for i := 0; i < numHeight-1; i++ {
		for j := i + 1; j < numHeight; j++ {
			l := height[i]
			r := height[j]
			h := l
			if r < l {
				h = r
			}
			area := h * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

/*
One pass solution.
Time complexity: O(n)
Space complexity: O(1), constant space required
*/
func maxAreaV1(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := height[l]
		if height[r] < height[l] {
			h = height[r]
		}
		area := h * (r - l)
		if area > maxArea {
			maxArea = area
		}
		if height[l] < height[r] {
			l = l + 1
		} else {
			r = r - 1
		}
	}

	return maxArea
}
