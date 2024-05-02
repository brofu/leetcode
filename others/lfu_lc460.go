package others

import "fmt"

type NodeSet struct {
	nodes      map[int]*Node
	head, tail *Node
}

func (this *NodeSet) Debug() {
	fmt.Println(this.nodes, this.head, this.head.Next, this.tail.Pre, this.tail)
}

func (this *NodeSet) RemoveTail() (int, bool) { //TODO: better name
	if this.tail.Pre != this.head {
		key := this.tail.Pre.Key
		this.RemoveKey(key)
		return key, true
	}
	return 0, false
}

func (this *NodeSet) RemoveKey(key int) {
	if node, ok := this.nodes[key]; ok {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		node.Pre = nil
		node.Next = nil
		delete(this.nodes, key)
	}
}

func (this *NodeSet) AddKey(key int) {
	node := &Node{
		Key:  key,
		Pre:  this.head,
		Next: this.head.Next,
	}
	this.head.Next.Pre = node
	this.head.Next = node
	this.nodes[key] = node
}

func (this *NodeSet) Size() int {
	return len(this.nodes)
}

func NewNodeSet() *NodeSet {
	o := &NodeSet{
		nodes: make(map[int]*Node),
		head:  &Node{},
		tail:  &Node{},
	}
	o.head.Next = o.tail
	o.tail.Pre = o.head
	return o
}

type LFUCache struct {
	size            int
	minFrequency    int
	keyToValue      map[int]int
	keyToFrequency  map[int]int
	frequencyToKeys map[int]*NodeSet
}

func ConstructorLFU(capacity int) LFUCache {
	o := LFUCache{
		size:            capacity,
		keyToValue:      make(map[int]int),
		keyToFrequency:  make(map[int]int),
		frequencyToKeys: make(map[int]*NodeSet),
	}
	return o
}

func (this *LFUCache) Get(key int) int {
	result := -1
	if val, ok := this.keyToValue[key]; ok {
		result = val
		this.increaseFrequency(key)
	}
	return result
}

func (this *LFUCache) Put(key int, value int) {

	if _, ok := this.keyToValue[key]; ok { // already exists
		this.keyToValue[key] = value
		this.increaseFrequency(key)
	} else {
		if this.size == len(this.keyToValue) { // full, remove the min frequency key
			this.removeMin()
		}
		minFrequency := 1
		this.keyToValue[key] = value
		this.keyToFrequency[key] = minFrequency
		ns := this.makeSureFrequencyToKeySetExist(minFrequency)
		ns.AddKey(key)
		this.minFrequency = minFrequency
	}
}

func (this *LFUCache) removeMin() {
	o := this.frequencyToKeys[this.minFrequency]
	key, success := o.RemoveTail()
	if !success {
		panic("Warning. No node deleted")
	}
	frequency := this.keyToFrequency[key]
	delete(this.keyToValue, key)
	delete(this.keyToFrequency, key)
	this.makeSureFrequencyToKeySetCleaned(frequency, false)
}

func (this *LFUCache) increaseFrequency(key int) {
	frequency, _ := this.keyToFrequency[key]
	this.keyToFrequency[key] = frequency + 1
	this.frequencyToKeys[frequency].RemoveKey(key)
	this.makeSureFrequencyToKeySetCleaned(frequency, true)
	ns := this.makeSureFrequencyToKeySetExist(frequency + 1)
	ns.AddKey(key)
}

func (this *LFUCache) makeSureFrequencyToKeySetExist(frequency int) *NodeSet {
	nodeSet, ok := this.frequencyToKeys[frequency]
	if !ok {
		nodeSet = NewNodeSet()
		this.frequencyToKeys[frequency] = nodeSet
	}
	return nodeSet
}

func (this *LFUCache) makeSureFrequencyToKeySetCleaned(frequency int, needUpdateMinFrequency bool) {
	if this.frequencyToKeys[frequency].Size() == 0 {
		delete(this.frequencyToKeys, frequency)
		if this.minFrequency == frequency && needUpdateMinFrequency {
			this.minFrequency += 1 // is this really correct?
		}
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func (this *LFUCache) Debug() {
	fmt.Println(this.minFrequency, this.keyToValue, this.keyToFrequency, this.frequencyToKeys)
}
