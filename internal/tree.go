package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewListFromTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	q := NewQueue()
	q.Enqueue(root)
	for !q.IsEmpty() {
		node := q.Dequeue().(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}

		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}

	return result
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
