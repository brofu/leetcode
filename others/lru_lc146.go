package others

type LRUCache struct {
	head *Node
	tail *Node
	size int
	maps map[int]*Node
}

func Constructor(capacity int) LRUCache {
	o := LRUCache{
		head: &Node{},
		tail: &Node{},
		size: capacity,
		maps: make(map[int]*Node, capacity),
	}

	o.head.Next = o.tail
	o.tail.Pre = o.head
	return o
}

func (this *LRUCache) Get(key int) int {
	result := -1
	if node, ok := this.maps[key]; ok {
		result = node.Val
		this.adjustLocation(node)
	}
	return result
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.maps[key]; ok {
		// adjust the value
		node.Val = value
		// adjust the location
		this.adjustLocation(node)
	} else {
		node := &Node{
			Key: key,
			Val: value,
		}
		if len(this.maps) == this.size { // delete the
			this.removeNode(this.tail.Pre)
		}
		this.addNode(node)
	}
}

// assume the node is NOT head or tail
func (this *LRUCache) removeNode(node *Node) {
	delete(this.maps, node.Key)
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// assume the node is NOT head or tail
// always add after the head
func (this *LRUCache) addNode(node *Node) {
	this.maps[node.Key] = node
	node.Next = this.head.Next
	this.head.Next.Pre = node
	this.head.Next = node
	node.Pre = this.head
}

func (this *LRUCache) adjustLocation(node *Node) {
	if node == this.head.Next { // already was the least used one
		return
	}
	this.removeNode(node)
	this.addNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
