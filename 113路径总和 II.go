package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) [][]int {
    var res [][]int
    var stack []int
    search(root, targetSum, &res, stack)
    return res
}

func search(root *TreeNode, targetSum int, res *[][]int, stack []int) {
    if root == nil {
        return
    }
    targetSum -= root.Val
    stack = append(stack, root.Val)
    if targetSum == 0 && root.Left == nil && root.Right == nil {
        *res = append(*res, append([]int{}, stack...))
        stack = stack[:len(stack)-1]
    }
    search(root.Left, targetSum, res, stack)
    search(root.Right, targetSum, res, stack)
}