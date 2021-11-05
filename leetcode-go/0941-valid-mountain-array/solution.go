package leetcode0941

/*
One pass solution. Scan the array till we find the element that is not
greater. We then check if all the following elements less than the previous
element till the end of the array. Time complexity is O(n) as we scan the
array once. Space complexity is O(1) as we do not need additional space.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func validMountainArray(arr []int) bool {
	lenArr := len(arr)
	if lenArr < 3 {
		return false
	}

	i := 0
	for i < lenArr-1 {
		if arr[i] >= arr[i+1] {
			break
		}
		i++
	}

	if i == 0 || i == lenArr-1 {
		return false
	}

	for i < lenArr-1 {
		if arr[i] <= arr[i+1] {
			break
		}
		i++
	}

	return i == lenArr-1
}
