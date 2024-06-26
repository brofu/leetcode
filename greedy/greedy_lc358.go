package greedy

import (
	"strings"
)

func rearrangeString(s string, k int) string {

	freqsMap := make(map[rune]int)
	maxFreqs := []rune{}
	secondFreqs := []rune{}

	for _, r := range s {
		if _, exists := freqsMap[r]; exists {
			freqsMap[r] += 1
		} else {
			freqsMap[r] += 1
		}
	}

	maxFreq := 0
	for _, freq := range freqsMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	for r, freq := range freqsMap {
		if freq == maxFreq {
			maxFreqs = append(maxFreqs, r)
			delete(freqsMap, r)
		}
		if freq == maxFreq-1 {
			secondFreqs = append(secondFreqs, r)
			delete(freqsMap, r)
		}
	}

	segments := make([]*strings.Builder, maxFreq)

	// set up the max and 2nd max frequency characters
	for i := 0; i < maxFreq; i++ {
		builder := &strings.Builder{}
		builder.Grow(k)
		segments[i] = builder

		// write the max freq s
		for _, r := range maxFreqs {
			segments[i].WriteRune(r)
		}

		// write the 2nd most freq
		if i < maxFreq-1 {
			for _, r := range secondFreqs {
				segments[i].WriteRune(r)
			}
		}
	}

	// distribution the left characters.
	segmentId := 0

	for r, freq := range freqsMap {
		for f := freq; f > 0; f-- {
			segments[segmentId].WriteRune(r)
			segmentId = (segmentId + 1) % (maxFreq - 1) // why mod (maxFreq - 1)?
		}
	}

	for i := 0; i < maxFreq-1; i++ {
		if segments[i].Len() < k {
			return ""
		}
	}

	str := make([]string, maxFreq)
	for i, builder := range segments {
		str[i] = builder.String()
	}

	return strings.Join(str, "")
}
