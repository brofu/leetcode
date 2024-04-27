package str

/**
1. O(n) space complexity is necessary, because need to check
	a. If one letter occurs in s1 and s2
	b. The occur times
*/
func checkInclusion(s1 string, s2 string) bool {

	recording := make(map[byte]int)
	needs := make(map[byte]int)

	match := 0
	left, right := 0, 0

	for _, s := range s1 {
		needs[byte(s)] += 1
	}

	for right < len(s2) {

		// expend the window
		o := s2[right]
		right += 1

		if count, ok := needs[o]; ok {
			recording[o] += 1
			if recording[o] == count { // means we get all the numbers of `o` so far
				match += 1
			}
		}

		// shrink the window
		for right-left >= len(s1) { // the condition we need to shrink the window

			if match == len(needs) { // hit a possible result
				return true
			}

			// shrink window
			o := s2[left]
			left += 1
			if count, ok := needs[o]; ok { // the removed one also in s1
				if count == recording[o] {
					match -= 1
				}
				recording[o] -= 1
			}
		}

	}

	return false
}

func checkInclusionV2(s1 string, s2 string) bool {

	recording := make(map[byte]int)
	needs := make(map[byte]int)

	match := 0
	left, right := 0, 0

	for _, s := range s1 {
		needs[byte(s)] += 1
	}

	for right < len(s2) {

		// expend the window
		o := s2[right]
		right += 1

		if count, ok := needs[o]; ok {
			recording[o] += 1
			if recording[o] == count { // means we get all the numbers of `o` so far
				match += 1
			}
		}

		// shrink the window
		for match == len(needs) {

			if right-left == len(s1) { // a possible result
				return true
			}

			o := s2[left]
			left += 1
			if count, ok := needs[o]; ok { // the removed one also in s1
				if count == recording[o] {
					match -= 1
				}
				recording[o] -= 1
			}

		}

	}

	return false
}
