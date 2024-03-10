package leetcode2236

import (
	"github.com/bharath23/leetcode-go/internal"
)

func checkTree(root *internal.TreeNode) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
