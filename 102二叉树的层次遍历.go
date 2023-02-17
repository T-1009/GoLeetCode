package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var res [][]int
	dfsLevel(root, 0, &res)
	return res
}

func dfsLevel(root *TreeNode, level int, res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) == level {
        *res = append(*res, []int{root.Val})
    } else {
        (*res)[level] = append((*res)[level], root.Val)
    }
    dfsLevel(root.Left, level+1, res)
    dfsLevel(root.Right, level+1, res)
}

