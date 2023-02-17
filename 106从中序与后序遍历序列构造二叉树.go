package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:pos])
			root.Right = buildTree(inorder[pos+1:], postorder[pos:])
		}
	}
	return root
}