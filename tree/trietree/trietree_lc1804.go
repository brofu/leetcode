package trietree

import "fmt"

type TrieV2 struct {
	tm *TrieMap
}

func ConstructorV2() TrieV2 {
	return TrieV2{
		tm: NewTrieMap(),
	}
}

func (this *TrieV2) Insert(word string) {
	if this.tm.ContainsKey(word) {
		this.tm.Put(word, this.tm.Get(word).(int)+1)
	} else {
		this.tm.Put(word, 1)
	}
}

func (this *TrieV2) CountWordsEqualTo(word string) int {
	val := this.tm.Get(word)
	if val == nil {
		return 0
	}
	return val.(int)
}

func (this *TrieV2) CountWordsStartingWith(prefix string) int {
	result := 0
	words := this.tm.KeysWithPrefix(prefix)
	fmt.Println("flag", words)
	for _, word := range words {
		result += this.tm.Get(word).(int)
	}
	return result
}

func (this *TrieV2) Erase(word string) {
	result := this.tm.Get(word)
	if result == nil {
		return
	}
	if result.(int) == 1 {
		this.tm.Remove(word)
	} else {
		this.tm.Put(word, this.tm.Get(word).(int)-1)
	}

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */
