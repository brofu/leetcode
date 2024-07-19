package array

// adopt the sliding window framework
// https://labuladong.online/algo/essential-technique/sliding-window-framework/#%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%A1%86%E6%9E%B6%E6%A6%82%E8%A7%88
func lengthOfLongestSubstring(s string) int {

	left := 0
	res := 0
	window := make(map[rune]int)

	for right, r := range s {
		window[r] += 1

		for window[r] > 1 {
			t := s[left]
			window[rune(t)] -= 1
			left += 1
		}

		newRes := right - left + 1
		if res < newRes {
			res = newRes
		}
	}

	return res
}

// left, right mark the window

func lengthOfLongestSubstringV2(s string) int {

	left := 0
	res := 0
	checkingMap := make(map[rune]int)

	for right, str := range s {
		if index, existing := checkingMap[str]; existing && index >= left {
			left = index + 1
		}
		checkingMap[str] = right

		newRes := right - left + 1

		if res < newRes {
			res = newRes
		}
	}

	return res
}

/**
KP.
	1.	This version can show the standard workflows of `Sliding Window`
*/
func lengthOfLongestSubstringPV1(s string) int {

	records := make(map[byte]int)
	result := 0
	left, right := 0, 0

	for ; right < len(s); right++ {

		// extend the window
		c := s[right]
		index, ok := records[c]
		if !ok {
			records[c] = right
			continue
		}

		// a valid result
		if result < right-left {
			result = right - left
		}
		records[c] = right

		// shrunk the window
		if left <= index {
			left = index + 1
		}
	}

	if result < right-left {
		result = right - left
	}
	return result
}
