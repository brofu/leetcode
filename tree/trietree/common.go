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

func (ts *TrieSet) ShortestPrefix(key string) string {
	return ts.tmap.ShortestPrefix(key)
}

func (ts *TrieSet) HasKeyWithPattern(pattern string) bool {
	return ts.tmap.HasKeyWithPattern(pattern)
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

// ShortestPrefix returns the key of short prefix
func (tm *TrieMap) ShortestPrefix(key string) string {

	p := tm.root

	for i := 0; i < len(key); i++ {
		if p == nil {
			return ""
		}
		//result += string(key[i]) // KP: why this is wrong?
		if p.Val != nil {
			//return result
			return key[:i]
		}
		c := getIndexLowerCase(key[i])
		p = p.Children[c]
	}

	if p != nil && p.Val != nil { // KP. why need to check this case?
		return key
	}

	return ""
}

func (tm *TrieMap) HasKeyWithPattern(pattern string) bool {
	return tm.hasKeyWithPattern(tm.root, pattern, 0)
}

// hasKeyWithPattern check if there is key with the pattern
// "." in the pattern means any letter
// KP. The logic to handle "." match

func (tm *TrieMap) hasKeyWithPattern(node *TrieTreeNode, pattern string, index int) bool {

	if node == nil {
		return false
	}

	if index == len(pattern) {
		return node.Val != nil // if the val != nil, means there is a key
	}

	c := pattern[index]

	// KP. logic to handle "."
	if c != '.' {
		node = node.Children[getIndexLowerCase(c)]
		return tm.hasKeyWithPattern(node, pattern, index+1)
	} else {
		for i := 0; i < 26; i++ {
			cNode := node.Children[i] // any one is ok
			if tm.hasKeyWithPattern(cNode, pattern, index+1) {
				return true
			}
		}
	}

	return false
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
