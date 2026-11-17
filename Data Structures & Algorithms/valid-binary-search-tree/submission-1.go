/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt, math.MaxInt)
}


func check(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	} else if root.Left != nil && root.Left.Val >= root.Val {
		return false
	} else if root.Right != nil && root.Right.Val <= root.Val {
		return false
	} else if root.Val >= maxVal {
		return false
	} else if root.Val <= minVal {
		return false
	}
    return check(root.Left, minVal, root.Val) && check(root.Right, root.Val, maxVal)
}