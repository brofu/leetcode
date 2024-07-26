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
	if !tm.ContainsKey(key) {
		tm.size++
	}
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

func (tm *TrieMap) LongestPrefix(key string) string {
	p := tm.root
	i := 0
	maxLen := 0

	for ; i < len(key); i++ {
		if p == nil {
			break
		}
		if p.Val != nil { // find a prefix
			maxLen = i
		}
		c := getIndexLowerCase(key[i])
		p = p.Children[c]
	}
	if p != nil && p.Val != nil {
		return key
	}
	return key[:maxLen]
}

/**
LongestPrefixV2 returns the longest prefix of all the keys

KP. Map has better formance for `Children`
*/
func (tm *TrieMap) LongestPrefixV2() string {
	p := tm.root
	result := ""
	for p != nil {
		count := 0
		var c int
		for index, node := range p.Children {
			if node != nil {
				c = index
				count++
				if count > 1 {
					break
				}
			}
		}
		if count != 1 {
			return result
		}
		result += string(byte(c + 'a'))
		p = p.Children[c]
	}
	return result
}

func (tm *TrieMap) HasKeyWithPattern(pattern string) bool {
	return tm.hasKeyWithPattern(tm.root, pattern, 0)
}

func (tm *TrieMap) KeysWithPrefix(prefix string) []string {
	result := []string{}
	//	for i := 0; i < len(prefix); i++ {
	//		if p == nil {
	//			return result
	//		}
	//		c := prefix[i]
	//		p = p.Children[getIndexLowerCase(c)]
	//	}
	// can result getNode

	p := tm.getNode(tm.root, prefix)
	if p == nil {
		return result
	}

	var dfs func(*TrieTreeNode, []byte)
	dfs = func(root *TrieTreeNode, bytes []byte) {
		if root == nil {
			return
		}

		if root.Val != nil { // get a key
			result = append(result, string(bytes))
			return
		}

		// backtrack framework
		for index, node := range root.Children {
			// choose, next layer and cancel
			c := getLowerCaseFromIndex(index)
			dfs(node, append(bytes, c))
		}
	}

	dfs(p, []byte(prefix))

	return result
}

// Remove remove a key
func (tm *TrieMap) Remove(key string) {
	tm.removeKey(tm.root, key, 0)
}

/**
removeKey removes a key from the tree

KP.
	1.	Cursively, handle the children nodes frist
	2.	Post-order to check 2 more condition (about if this node can be deleted)
		a.	if the node.val == nil
		b.	if there are children node
	2.  and if a node can be delete, return nil to parent node
*/

func (tm *TrieMap) removeKey(node *TrieTreeNode, key string, index int) *TrieTreeNode {
	if node == nil { // the node is nil. means the key doesn't exist
		return nil
	}

	// get the key, remove the value
	if index == len(key) { // node is the last node of the key
		node.Val = nil // but not the time to delete the node yet, why? (it may have children.)
	} else { // delete the child cursively.
		c := getIndexLowerCase(key[index])
		node.Children[c] = tm.removeKey(node.Children[c], key, index+1) // similar as add new key
	}

	if node.Val != nil { // the node has val. it is a prefix of key
		return node
	}

	for _, next := range node.Children { // this node has children
		if next != nil {
			return node
		}
	}

	return nil // the node delete itself
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
	if c != '.' { // Not "."
		node = node.Children[getIndexLowerCase(c)]
		return tm.hasKeyWithPattern(node, pattern, index+1)
	}
	// Is "."
	for i := 0; i < 26; i++ { // any one is ok
		if tm.hasKeyWithPattern(node.Children[i], pattern, index+1) {
			return true
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

func getLowerCaseFromIndex(index int) byte {
	return byte(index) + 'a'
}
