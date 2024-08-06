package others

import "fmt"

type dataNode struct {
	key      int
	val      int
	previous *dataNode
	next     *dataNode
}

type LRUCacheV2 struct {
	size    int
	dataMap map[int]*dataNode
	head    *dataNode
	tail    *dataNode
}

func ConstructorV2(capacity int) LRUCacheV2 {
	lc := LRUCacheV2{
		size:    capacity,
		dataMap: make(map[int]*dataNode, capacity),
		head:    &dataNode{},
		tail:    &dataNode{},
	}

	lc.head.next = lc.tail
	lc.tail.previous = lc.head
	return lc
}

func (this *LRUCacheV2) Get(key int) int {

	node, ok := this.dataMap[key]
	if !ok {
		return -1
	}

	this.adjustNode(node)

	return node.val
}

func (this *LRUCacheV2) Put(key int, value int) {

	node, ok := this.dataMap[key]
	if ok { // exists
		node.val = value
		this.adjustNode(node)
		return
	}

	// if key not exists yet
	if len(this.dataMap) == this.size { // full

		// retire the last one

		// remove the retired one
		oldNode := this.tail.previous
		this.tail.previous = oldNode.previous
		oldNode.next = this.tail
		delete(this.dataMap, oldNode.key)

	}

	// setup a new node
	node = &dataNode{
		key:      key,
		val:      value,
		previous: this.head,
		next:     this.head.next,
	}

	this.head.next.previous = node
	this.head.next = node

	this.dataMap[key] = node
}

func (this *LRUCacheV2) adjustNode(node *dataNode) {

	if node.previous != this.head {
		node.previous.next = node.next
		node.next.previous = node.previous

		node.next = this.head.next
		node.previous = this.head
		this.head.next.previous = node
		this.head.next = node
	}
}

func (this *LRUCacheV2) Debug(message string) {

	fmt.Println(message, "dataMap, ", this.dataMap)

	p := this.head
	for p != nil {
		fmt.Println(message, p.key, p.val)
		p = p.next
	}
}
