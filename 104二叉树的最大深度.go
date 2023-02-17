package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
    var maxDepth int
    maxDepth = childTreeDepth(root, 0)
    return maxDepth
}

func childTreeDepth(root *TreeNode, depth int) int {
    if root == nil {
        return 0
    }
    left := childTreeDepth(root.Left, depth+1)
    right := childTreeDepth(root.Right, depth+1)
    if left > right {
        return left + 1
    } else {
        return right + 1
    }
}