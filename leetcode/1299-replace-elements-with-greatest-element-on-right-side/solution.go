package leetcode1299

/*
Simple one pass solution. We set the maximum to -1 and walk the array in
reverse. If the current element is greater than the maximum we swap the
value of max and the current element, else we set the element to the maximum.
Time complexity is O(n) as we traverse the array once. The space complexity is
O(1) as no additional storage is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			t := max
			max = arr[i]
			arr[i] = t
		} else {
			arr[i] = max
		}
	}

	return arr
}
