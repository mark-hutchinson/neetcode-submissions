
const NumBuckets = 100

type Node struct {
	Key int
	Next *Node
}

type MyHashSet struct {
	// other idea - slice of slices
	Set []*Node
}

func Constructor() MyHashSet {
    hs := MyHashSet{}
	hs.Set = make([]*Node, NumBuckets)
	return hs
}

func (this *MyHashSet) Add(key int) {
    if this.Contains(key) {
		// There is a more optimal way - could avoid doing htis pass but fine for now
		return
	}
	node := &Node{
		Key: key,
	}
	index := hash(key)
	if this.Set[index] == nil {
		this.Set[index] = node
	} else {
		curr := this.Set[index]
		// find the last node
		for curr.Next != nil {
			curr = curr.Next
		}
		// add the new node
		curr.Next = node
	}
}

func (this *MyHashSet) Remove(key int) {
    index := hash(key)
	if this.Set[index] == nil {
		return
	}

	var prev *Node
	curr := this.Set[index]
	for curr != nil {
		if curr.Key == key {
			// Delete it
			if prev == nil {
				// first node needs deleting
				this.Set[index] = curr.Next
			} else {
				prev.Next = curr.Next
			}
			return
		}
		prev = curr
		curr = curr.Next
	}	
}

func (this *MyHashSet) Contains(key int) bool {
    index := hash(key)
	if this.Set[index] == nil {
		return false
	}
	curr := this.Set[index]
	for curr != nil {
		if curr.Key == key {
			return true
		}
		curr = curr.Next
	} 
	return false
}


func hash(key int)int {
	return key % NumBuckets;
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 