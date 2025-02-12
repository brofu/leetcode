package hashmap

type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}

type LinkedHashMap struct {
	data map[string]*Node
	head *Node
	tail *Node
}

func Constructor() LinkedHashMap {

	head := &Node{}
	tail := &Node{
		prev: head,
	}
	head.next = tail

	hm := LinkedHashMap{
		data: make(map[string]*Node),
		head: head,
		tail: tail,
	}

	return hm
}

func (hm *LinkedHashMap) Get(key string) int {
	if node, exists := hm.data[key]; exists {
		return node.value
	}
	return -1
}

func (hm *LinkedHashMap) Put(key string, value int) {
	node, exists := hm.data[key]

	if exists {
		node.value = value
		return
	}

	node = &Node{
		key:   key,
		value: value,
	}
	node.prev = hm.tail.prev
	node.next = hm.tail
	hm.tail.prev.next = node
	hm.tail.prev = node

	hm.data[key] = node
}

func (hm *LinkedHashMap) Remove(key string) {

	node, exists := hm.data[key]

	if !exists {
		return
	}

	delete(hm.data, key)

	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (hm *LinkedHashMap) ContainsKey(key string) bool {
	_, exists := hm.data[key]
	return exists
}

func (hm *LinkedHashMap) Keys() []string {
	keys := []string{}
	if hm.head.next == hm.tail {
		return keys
	}
	for p := hm.head.next; p != nil; p = p.next {
		keys = append(keys, p.key)
	}
	return keys
}
