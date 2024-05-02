package others

import "fmt"

type FrequencyNodeList struct {
	size       int
	head, tail *FrequencyNode
}

func NewFrequencyNodeList() *FrequencyNodeList {
	o := &FrequencyNodeList{
		head: &FrequencyNode{},
		tail: &FrequencyNode{},
	}
	o.head.Next = o.tail
	o.tail.Pre = o.head
	return o
}

func (this *FrequencyNodeList) AddNode(node *FrequencyNode) {
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next.Pre = node
	this.head.Next = node
	this.size += 1
}

func (this *FrequencyNodeList) RemoveTail() *FrequencyNode {
	if this.head.Next == this.tail {
		return nil
	}
	node := this.tail.Pre
	this.RemoveNode(this.tail.Pre)
	return node
}

func (this *FrequencyNodeList) RemoveNode(node *FrequencyNode) {
	if this.head.Next == this.tail {
		return
	}
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	this.size -= 1
}

func (this *FrequencyNodeList) Size() int {
	return this.size
}

type LFUCacheV2 struct {
	capbility       int
	minFrequency    int
	cache           map[int]*FrequencyNode
	frequencyToKeys map[int]*FrequencyNodeList
}

func ConstructorLFUV2(capacity int) LFUCacheV2 {
	o := LFUCacheV2{
		capbility:       capacity,
		cache:           make(map[int]*FrequencyNode),
		frequencyToKeys: make(map[int]*FrequencyNodeList),
	}
	return o
}

func (this *LFUCacheV2) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.increaseFrequency(node)
		return node.Val
	}
}

func (this *LFUCacheV2) Debug(debugMsg string, objs ...interface{}) {
	fmt.Println(debugMsg, this.minFrequency, this.cache, this.frequencyToKeys, objs)
}

func (this *LFUCacheV2) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.increaseFrequency(node)
		return
	}

	if len(this.cache) == this.capbility {
		this.removeMin()
	}

	node := &FrequencyNode{
		Node: Node{
			Key: key,
			Val: value},
		Frequency: 1,
	}
	this.cache[key] = node

	list := this.makeSureFrequencyToKeySetExist(node)
	list.AddNode(node)
	this.frequencyToKeys[node.Frequency] = list
	this.minFrequency = node.Frequency
}

func (this *LFUCacheV2) removeMin() {
	if list, ok := this.frequencyToKeys[this.minFrequency]; ok {
		node := list.RemoveTail()
		delete(this.cache, node.Key)
	}
}

func (this *LFUCacheV2) increaseFrequency(node *FrequencyNode) {

	this.frequencyToKeys[node.Frequency].RemoveNode(node)
	this.makeSureFrequencyToKeySetCleaned(node)

	node.Frequency += 1

	list := this.makeSureFrequencyToKeySetExist(node)
	list.AddNode(node)
	this.frequencyToKeys[node.Frequency] = list
}

func (this *LFUCacheV2) makeSureFrequencyToKeySetCleaned(node *FrequencyNode) {
	list := this.frequencyToKeys[node.Frequency]
	if list.Size() == 0 {
		delete(this.frequencyToKeys, node.Frequency)
		if node.Frequency == this.minFrequency {
			this.minFrequency += 1
		}
	}
}

func (this *LFUCacheV2) makeSureFrequencyToKeySetExist(node *FrequencyNode) *FrequencyNodeList {
	list, ok := this.frequencyToKeys[node.Frequency]
	if !ok {
		list = NewFrequencyNodeList()
	}
	return list
}
