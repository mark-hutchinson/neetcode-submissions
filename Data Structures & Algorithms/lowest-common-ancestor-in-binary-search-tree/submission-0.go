/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// both numbers less than -> traverse left
	if p.Val < root.Val && q.Val < root.Val {
		// common ancestor is to the left
		return lowestCommonAncestor(root.Left, p, q)
	}
	// both numbers greater than root -> traverse right
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	
	// one number lte , other gte, we found it
	return root
	
}
