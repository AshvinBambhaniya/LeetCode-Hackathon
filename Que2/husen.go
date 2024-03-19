/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    leftDepthOfTree := 1
    rightDepthOfTree := 1

    leftDepthOfTree += maxDepth(root.Left)
    rightDepthOfTree += maxDepth(root.Right)

    if leftDepthOfTree > rightDepthOfTree {
        return leftDepthOfTree
    }
    return rightDepthOfTree
}
