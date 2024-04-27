package str

import (
	"math"
)

func minWindow(s string, t string) string {

	left, right := 0, 0
	start, length := 0, math.MaxInt32
	match := 0

	records := make(map[byte]int)
	need := make(map[byte]int)

	for _, temp := range t {
		need[byte(temp)] += 1
	}

	for right < len(s) {

		// expand the window
		d := s[right]
		right += 1

		if count, ok := need[d]; ok { // we need this byte
			records[d] += 1
			if records[d] == count {
				match += 1
			}
		}

		for match == len(need) { // all the letters in T are matched requirements
			// record current (start, length) tuple

			if right-left < length {
				start = left
				length = right - left
			}

			// shrink the records
			d := s[left]
			if count, ok := need[d]; ok {
				if records[d] == count {
					match -= 1
				}
				records[d] -= 1
			}
			left += 1
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}
