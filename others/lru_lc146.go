package others

type LRUCacheV1 struct {
	head *Node
	tail *Node
	size int
	maps map[int]*Node
}

func ConstructorV1(capacity int) LRUCacheV1 {
	o := LRUCacheV1{
		head: &Node{},
		tail: &Node{},
		size: capacity,
		maps: make(map[int]*Node, capacity),
	}

	o.head.Next = o.tail
	o.tail.Pre = o.head
	return o
}

func (this *LRUCacheV1) Get(key int) int {
	result := -1
	if node, ok := this.maps[key]; ok {
		result = node.Val
		this.adjustLocation(node)
	}
	return result
}

func (this *LRUCacheV1) Put(key int, value int) {
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
func (this *LRUCacheV1) removeNode(node *Node) {
	delete(this.maps, node.Key)
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// assume the node is NOT head or tail
// always add after the head
func (this *LRUCacheV1) addNode(node *Node) {
	this.maps[node.Key] = node
	node.Next = this.head.Next
	this.head.Next.Pre = node
	this.head.Next = node
	node.Pre = this.head
}

func (this *LRUCacheV1) adjustLocation(node *Node) {
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

type LRUCache struct {
	head       *Node
	tail       *Node
	data       map[int]*Node
	capability int
}

func Constructor(capability int) *LRUCache {

	cache := &LRUCache{
		head:       &Node{},
		tail:       &Node{},
		data:       make(map[int]*Node),
		capability: capability,
	}
	cache.head.Next = cache.tail
	cache.tail.Pre = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.data[key]
	if !exists {
		return -1
	}
	this.promoteNode(node)
	return node.Val
}

func (this *LRUCache) Put(key, val int) {

	// key exists
	node, exists := this.data[key]
	if exists {
		node.Val = val
		this.promoteNode(node)
		return
	}

	node = &Node{
		Key: key,
		Val: val,
	}

	this.data[key] = node
	this.promoteNode(node)

	//key not exists, not exceed the capability
	// exceed the capability
	if len(this.data) > this.capability {
		tailNode := this.tail.Pre
		if tailNode == this.head {
			return
		}

		tailNode.Pre.Next = this.tail
		this.tail.Pre = tailNode.Pre

		delete(this.data, tailNode.Key)
	}

}

func (this *LRUCache) promoteNode(node *Node) {
	if this.head.Next == node {
		return
	}

	if node.Pre != nil {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}

	this.head.Next.Pre = node
	node.Next = this.head.Next
	this.head.Next = node
	node.Pre = this.head
}
