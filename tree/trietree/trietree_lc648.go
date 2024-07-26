package trietree

import "strings"

func replaceWords(dictionary []string, sentence string) string {

	words := strings.Split(sentence, " ")
	newWords := make([]string, len(words))

	tm := NewTrieSet()
	for _, w := range dictionary {
		tm.Put(w)
	}
	for index, w := range words {
		prefix := tm.ShortestPrefix(w)
		if prefix == "" {
			prefix = w
		}
		newWords[index] = prefix
	}

	return strings.Join(newWords, " ")
}
