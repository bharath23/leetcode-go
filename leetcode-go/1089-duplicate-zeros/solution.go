package leetcode1089

/*
Trivial one pass solution. When you encounter a '0', move all elements to its
right by one and then duplicate '0'. Since we have an inner loop time worst
case time complexity is O(n^2)

Complexity:
Time complexity: O(n^2)
Space complexity: O(1), no additional space is required
*/

func duplicateZerosV0(arr []int) {
	lenArr := len(arr)
	for i := 0; i < lenArr; i++ {
		if arr[i] != 0 {
			continue
		}

		for j := lenArr - 1; j > i; j-- {
			arr[j] = arr[j-1]
		}
		i++
	}

	return
}

/*
Two pass solution. First pass we count the number of zeros that we can
duplicate. If we find a zero that cannot be duplicated we copy that that as
the last element and next pass duplicate the zeros. Since we go through the
array twice the time complexity is O(n).

Complexity:
Time complexity: O(n)
Space complexity: O(1), no additional space is required.
*/
func duplicateZerosV1(arr []int) {
	lenArr := len(arr)
	numZeros := 0
	for i := 0; i < lenArr-numZeros; i++ {
		if arr[i] == 0 {
			if i == lenArr-numZeros-1 {
				arr[lenArr-1] = arr[i]
				lenArr--
				break
			}
			numZeros++
		}
	}

	for i := lenArr - numZeros - 1; i >= 0; i-- {
		arr[i+numZeros] = arr[i]
		if arr[i] == 0 {
			numZeros--
			arr[i+numZeros] = 0
		}
	}

	return
}
