package GoLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    lenTree := len(nums)
    if lenTree == 0 {
        return nil
    }
    return &TreeNode{Val: nums[lenTree/2], 
                    Left: sortedArrayToBST(nums[:lenTree/2]),
                    Right: sortedArrayToBST(nums[lenTree/2+1:])}
}