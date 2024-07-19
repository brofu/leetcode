package array

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

func minWindowPV1(s string, t string) string {

	need, window := make(map[byte]int), make(map[byte]int)

	for _, r := range t {
		need[byte(r)] += 1
	}

	start := 0
	length := math.MaxInt

	valid := 0
	left := 0
	for right := 0; right < len(s); right++ {

		// expand the window
		r := s[right]
		if count, ok := need[r]; ok {
			window[r] += 1
			if haveCount, _ := window[r]; haveCount == count {
				valid += 1
			}
		}

		// check if we have an valid result
		// if yes, shrunk the window
		for len(need) == valid {

			if right-left < length { // get a better valid result
				start = left
				length = right - left + 1
			}

			// shrunk the window
			o := s[left]
			if count, ok := need[o]; ok {
				if haveCount, _ := window[o]; haveCount == count {
					valid -= 1
				}
				window[o] -= 1
			}
			left += 1
		}
	}

	if length == math.MaxInt {
		return ""
	}

	return s[start : start+length]
}
