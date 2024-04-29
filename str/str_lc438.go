package str

/*
KP
	1. Is there duplicate letters in `p`?
*/
func findAnagrams(s string, p string) []int {

	start, count, match := 0, 0, 0
	var recording, need [26]int
	var result []int

	for _, o := range p {
		need[o-'a'] += 1
	}

	for _, i := range need {
		if i > 0 {
			count += 1
		}
	}

	for end, o := range s {

		// expand window
		index := o - 'a'
		if need[index] > 0 {
			recording[index] += 1
			if recording[index] == need[index] {
				match += 1
			}
		}

		for match == count {
			if end-start+1 == len(p) {
				result = append(result, start)
			}

			p := s[start]
			index := p - 'a'
			if need[index] > 0 {
				if need[index] == recording[index] {
					match -= 1
				}
				recording[index] -= 1
			}
			start += 1
		}
	}

	return result
}

func findAnagramsV2(s string, p string) []int {
	var result []int

	start, count, match := 0, 0, 0

	var need, recording [26]int

	for _, o := range p {
		need[o-'a'] += 1
	}

	for _, i := range need {
		if i > 0 {
			count += 1
		}
	}

	for end, o := range s {
		index := o - 'a'

		if need[index] > 0 {
			recording[index] += 1
			if recording[index] == need[index] {
				match += 1
			}
		}

		for end-start+1 >= len(p) {
			if match == count {
				result = append(result, start)
			}
			q := s[start]
			index := q - 'a'
			if need[index] > 0 {
				if recording[index] == need[index] {
					match -= 1
				}
				recording[index] -= 1
			}
			start += 1
		}

	}
	return result

}
