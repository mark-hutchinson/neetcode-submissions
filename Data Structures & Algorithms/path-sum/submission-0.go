/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	target := targetSum - root.Val

	// are we a leaf node?
	if root.Left == nil && root.Right == nil {
		return target == 0
	}
	if root.Left != nil && hasPathSum(root.Left, target){
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, target) {
		return true
	}
	return false
}
