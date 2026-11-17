/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true // the trees are equal
	} else if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val {
		// Possible subtree
		if isEqual(root, subRoot) {
			return true
		}
	}
	if isSubtree(root.Left, subRoot) {
		return true
	}
	if isSubtree(root.Right, subRoot) {
		return true
	}
	
	return false
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if !isEqual(a.Left, b.Left) {
		return false
	}
	if !isEqual(a.Right, b.Right) {
		return false
	}
	return true
}