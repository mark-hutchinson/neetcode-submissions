type PrefixTree struct {
	Root *Node
}

type Node struct {
	Children  [26]*Node
	Value     rune
	IsWordEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		Root : &Node{},
	}
}

func (this *PrefixTree) Insert(word string) {
	node := this.Root
	for _, c := range word {
		index := c - 'a'
		if node.Children[index] == nil {
			node.Children[index] = &Node{
				Value: c,
			}
		}
		node = node.Children[index]
	}
	// marke the last node as the end of a word
	node.IsWordEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.Root
	for _, c := range word {
		index := c - 'a'
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return node.IsWordEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.Root
	for _, c := range prefix {
		index := c - 'a'
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return true
}
