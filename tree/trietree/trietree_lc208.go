package trietree

type Trie struct {
	tset *TrieSet
}

func Constructor() Trie {
	return Trie{
		tset: NewTrieSet(),
	}
}

func (this *Trie) Insert(word string) {
	this.tset.Put(word)
}

func (this *Trie) Search(word string) bool {
	return this.tset.ContainsKey(word)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.tset.HasKeyWithPrefix(prefix)
}
