/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// The node is always the first element of preorder
	val := preorder[0]
	node := &TreeNode{
		Val: val,
	}

	// Figure out which nodes are left and right, by lookig at the inorder array
	midpoint := 0
	for midpoint=0; midpoint<len(inorder);midpoint++ {
		if inorder[midpoint] == val {
			break
		}
	}
	
	// Split the inorder array at the midpoint we found
	leftInorder := inorder[:midpoint]
	rightInorder := inorder[midpoint+1:]
	// Split the preorder array to be the same size
	leftPreorder := preorder[1:len(leftInorder) + 1]
	var rightPreorder []int
	if len(preorder) > len(leftInorder) + 1 {
		rightPreorder = preorder[len(leftInorder) + 1:]
	}
	
	//fmt.Printf("val=%d, lp=%v rp=%v li=%v, ri=%v",val, leftPreorder, rightPreorder, leftInorder, rightInorder)

	// build left and right trees
	node.Left = buildTree(leftPreorder, leftInorder)
	node.Right = buildTree(rightPreorder, rightInorder)
	return node
}
