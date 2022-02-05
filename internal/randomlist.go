package internal

type RandomList struct {
	Val    int
	Next   *RandomList
	Random *RandomList
}

func NewRandomListFromSlice(input [][]int) *RandomList {
	var head, tail *RandomList

	nodeMap := make(map[int]*RandomList, len(input))
	for i, val := range input {
		n := &RandomList{Val: val[0]}
		if head == nil {
			head = n
		}

		if tail != nil {
			tail.Next = n
		}

		tail = n
		nodeMap[i] = n
	}

	cur := head
	for _, val := range input {
		if val[1] != -1 {
			cur.Random = nodeMap[val[1]]
		}

		cur = cur.Next
	}

	return head
}

func RandomListToSlice(head *RandomList) [][]int {
	s := [][]int{}
	nodeMap := make(map[*RandomList]int)
	idx := 0
	for cur := head; cur != nil; cur = cur.Next {
		nodeMap[cur] = idx
		idx++
	}

	for cur := head; cur != nil; cur = cur.Next {
		e := []int{cur.Val}
		if cur.Random != nil {
			e = append(e, nodeMap[cur.Random])
		} else {
			e = append(e, -1)
		}
		s = append(s, e)
	}

	return s
}
