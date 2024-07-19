package array

/**
KP.
	1.	The function of `extend and shrunk window` maybe different problem by problem.
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {

	result := 0
	records := make(map[byte]struct{})
	left, right := 0, 0

	for ; right < len(s); right++ {

		// extend the window
		r := s[right]
		if len(records) < 2 {
			records[r] = struct{}{}
			continue
		}
		if _, ok := records[r]; ok {
			records[r] = struct{}{}
			continue
		}

		// len(record) == 2 && r not in records.
		// we have a valid result
		length := right - left
		if result < length { // a better result
			result = length
		}
		records[r] = struct{}{}

		// shrunk the window
		left = right - 1
		for o := s[right-1]; s[left] == o && left >= 0; left-- {
		}
		delete(records, s[left])
		left++
	}

	// KP. The edged case
	if result < right-left {
		result = right - left
	}
	return result
}
