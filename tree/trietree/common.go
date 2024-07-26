package trietree

type TrieTreeNode struct {
	Val      interface{}
	Children [26]*TrieTreeNode // only support 26 letters
}

// TrieSet is implemented based on `TrieMap`
type TrieSet struct {
	tmap *TrieMap
}

func NewTrieSet() *TrieSet {
	return &TrieSet{
		tmap: NewTrieMap(),
	}
}

func (ts *TrieSet) Put(key string) {
	ts.tmap.Put(key, struct{}{})
}

func (ts *TrieSet) ContainsKey(key string) bool {
	return ts.tmap.ContainsKey(key)
}

func (ts *TrieSet) HasKeyWithPrefix(prefix string) bool {
	return ts.tmap.HasKeyWithPrefix(prefix)
}

type TrieMap struct {
	size int
	root *TrieTreeNode
}

func NewTrieMap() *TrieMap {
	return &TrieMap{
		root: &TrieTreeNode{
			Children: [26]*TrieTreeNode{},
		},
	}
}

func (tm *TrieMap) Get(key string) interface{} {
	node := tm.getNode(tm.root, key)
	if node == nil {
		return nil
	}
	return node.Val
}

func (tm *TrieMap) Put(key string, val interface{}) {
	// check if contains
	// if contains, just return for now
	if tm.ContainsKey(key) {
		return
	}
	tm.size++
	tm.root = tm.putNode(tm.root, key, val, 0)
}

func (tm *TrieMap) ContainsKey(key string) bool {
	return tm.Get(key) != nil
}

func (tm *TrieMap) HasKeyWithPrefix(prefix string) bool {
	return tm.getNode(tm.root, prefix) != nil
}

// putNode put (key, val) from the root node, and return the inserted root node
func (tm *TrieMap) putNode(root *TrieTreeNode, key string, val interface{}, index int) *TrieTreeNode {

	if root == nil { // create a new one
		root = &TrieTreeNode{
			Children: [26]*TrieTreeNode{},
		}
	}

	if index == len(key) {
		root.Val = val
		return root
	}

	c := getIndexLowerCase(key[index])
	root.Children[c] = tm.putNode(root.Children[c], key, val, index+1)

	return root
}

func (tm *TrieMap) getNode(root *TrieTreeNode, key string) *TrieTreeNode {
	p := root
	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}
		c := getIndexLowerCase(key[i])
		p = p.Children[c]
	}
	return p
}

func getIndexLowerCase(r byte) int {
	return int(r - 'a')
}
