package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeFromList(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	q := NewQueue()
	root := &TreeNode{Val: input[0]}
	q.Enqueue(root)
	i := 1
	for i < len(input) {
		cur := q.Dequeue().(*TreeNode)
		node := &TreeNode{Val: input[i]}
		cur.Left = node
		q.Enqueue(node)
		i = i + 1
		if i < len(input) {
			node := &TreeNode{Val: input[i]}
			cur.Right = node
			q.Enqueue(node)
			i = i + 1
		}
	}

	return root
}
