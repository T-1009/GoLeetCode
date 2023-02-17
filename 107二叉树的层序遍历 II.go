package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrderBottom(root *TreeNode) [][]int {
    var res, ans [][]int
    levelOrder(root, 0, &res)
    for _, val := range res {
        ans = append([][]int{val}, ans...)
    }
    return ans
}

func levelOrder(root *TreeNode, depth int, res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) == depth {
        *res = append(*res, []int{root.Val})
    } else { // 放入右子树
        (*res)[depth] = append((*res)[depth], root.Val)
    }
    levelOrder(root.Left, depth+1, res)
    levelOrder(root.Right, depth+1, res)
}