package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    return isBalanced(root.Left) && isBalanced(root.Right) && abs(leftDepth, rightDepth)<=1
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}

func abs(a int, b int) int {
    if a >= b {
        return a-b
    }
    return b-a
}