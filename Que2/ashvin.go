package que2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := maxDepth(root.Left)
	rightH := maxDepth(root.Right)

	return max(leftH, rightH) + 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}
