package hashmap

import "math/rand"

/*
ArrayHashMap
  - Provide the ability to return a random (uniform random) key from the hashmap with O(1)
  - Can not utilize the under layer array of hash map directly? why?
    -- key distribution in the array is NOT continuous. If there is NO key on some index, how to handle?
    -- probing to the left or right? Not O(1), Not random.
    -- more complex for Separate Chaining...
*/

type ArrayNode struct {
	key   string
	value int
	index int
}

type ArrayHashMap struct {
	data  map[string]*ArrayNode
	array []*ArrayNode
}

func NewArrayHashMap() *ArrayHashMap {
	return &ArrayHashMap{
		data:  make(map[string]*ArrayNode),
		array: []*ArrayNode{},
	}
}

func (hm *ArrayHashMap) Get(key string) int {
	node, exists := hm.data[key]
	if !exists {
		return -1
	}
	return node.value
}

func (hm *ArrayHashMap) Put(key string, value int) {
	node, exists := hm.data[key]
	if exists {
		node.value = value
		return
	}
	node = &ArrayNode{
		key:   key,
		value: value,
		index: len(hm.array),
	}
	hm.data[key] = node
	hm.array = append(hm.array, node)
}

// Remove
// swap the target node with the TAIL node in the array.
// This would break the insert order of the keys
func (hm *ArrayHashMap) Remove(key string) {
	node, exists := hm.data[key]
	if !exists {
		return
	}

	// delete from map
	delete(hm.data, key)

	// delete from array
	if node.index == len(hm.array)-1 { // the last one in the array
		hm.array[len(hm.array)-1] = nil
		return
	}

	// not the last array
	tailNode := hm.array[len(hm.array)-1]
	hm.array[node.index], hm.array[len(hm.array)-1] = tailNode, nil
	tailNode.index = node.index
}

func (hm *ArrayHashMap) randomKey() string {
	index := rand.Intn(len(hm.array))
	return hm.array[index].key
}
