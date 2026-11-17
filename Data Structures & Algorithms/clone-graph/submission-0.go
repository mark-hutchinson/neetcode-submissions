/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}
	_, ok := visited[node]
	if ok {
		return visited[node]
	}

	newNode := &Node{
		Val: node.Val,
	}
	visited[node] = newNode

	for _, child := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(child, visited))
	}
	return newNode
}
