package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
    var res [][]int
    var ans []int
    search(root, 0, &res)
    for i:=0; i<len(res); i++ {
        ans = append(ans, res[i][len(res[i])-1])
    }
    return ans
}

func search(root *TreeNode, level int ,res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) == level {
        *res = append(*res, []int{root.Val})
    } else {
        (*res)[level] = append((*res)[level], root.Val)
    }
    search(root.Left, level+1, res)
    search(root.Right, level+1, res)
}