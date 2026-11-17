/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const Nil = "nil"

type Codec struct {
    Output []string
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.traverse(root)

	str := strings.Join(this.Output, ",")
	fmt.Printf("Result =%s\n", str)
	return str
}

func (this *Codec) traverse(root *TreeNode) {
	if root == nil {
		this.Output = append(this.Output, Nil)
	} else {
		strval := strconv.Itoa(root.Val)
		this.Output = append(this.Output, strval)
		this.traverse(root.Left)
		this.traverse(root.Right)
	}
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	this.Output = vals
	// fmt.Printf("Result =%s\n", vals)
    return this.createTree()
}

func (this *Codec) createTree() *TreeNode {
	vals := this.Output
	if len(vals) == 0 {
		return nil
	}
	if vals[0] == Nil {
		this.Output = vals[1:]
		return nil
	}

	intVal, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}

	node := &TreeNode{Val: intVal}
	this.Output = vals[1:]
	node.Left = this.createTree()
	node.Right = this.createTree()
	return node
}

