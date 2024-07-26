package trietree

type WordDictionary struct {
	ts *TrieSet
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{
		ts: NewTrieSet(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.ts.Put(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.ts.HasKeyWithPattern(word)
}
