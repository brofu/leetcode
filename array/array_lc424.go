package array

import "github.com/brofu/leetcode/common"

func characterReplacement(s string, k int) int {

	result := 0
	unique := make(map[byte]int)

	type subMax struct {
		count int
		rs    map[byte]struct{}
	}

	subMaxR := subMax{0, make(map[byte]struct{})}
	subSum := 0

	for left, right := 0, 0; right < len(s); right++ {

		// extend the window
		r := s[right]
		subSum++
		unique[r]++

		if subMaxR.count < unique[r] {
			subMaxR.count = unique[r]
			subMaxR.rs = make(map[byte]struct{}) // a new record
			subMaxR.rs[r] = struct{}{}
		}

		if subMaxR.count == unique[r] {
			subMaxR.rs[r] = struct{}{}
		}

		// shrunk window
		for subSum-subMaxR.count > k {
			o := s[left]
			subSum--
			if unique[o] == subMaxR.count { // o also has max count
				delete(subMaxR.rs, o)
				if len(subMaxR.rs) == 0 { // search the new max value
					subMaxR.count = 0
					for k, v := range unique {
						if subMaxR.count < v {
							subMaxR.count = v
							subMaxR.rs[k] = struct{}{}
						}
						if subMaxR.count == v {
							subMaxR.rs[k] = struct{}{}
						}
					}
				}
			}
			unique[o]--
			left++
		}

		if result < right-left+1 {
			result = right - left + 1
		}
	}

	return result
}

// The brief version
/**
KP
	*	Utilize the `index` of the character in s. For example,
		*	the sum of the characters in the window,
		*	the count of other character except the one which as largest frequency.
*/
func characterReplacementV2(s string, k int) int {
	result := 0

	maxFrequency := 0
	frequency := make(map[byte]int) // KP. array has better performance.
	for left, right := 0, 0; right < len(s); right++ {

		// extend the window
		r := s[right]
		frequency[r]++
		maxFrequency = common.MaxInt(maxFrequency, frequency[r])

		// shrunk the window
		if right-left+1-maxFrequency > k {
			o := s[left]
			frequency[o]--
			left++
		}

		result = common.MaxInt(result, right-left+1)
	}

	return result
}

func characterReplacementPV3(s string, k int) int {
	result := 0

	maxFrequency := 0
	frequency := make(map[byte]int) // KP. array has better performance.
	for left, right := 0, 0; right < len(s); right++ {

		// extend the window
		r := s[right]
		frequency[r]++
		maxFrequency = common.MaxInt(maxFrequency, frequency[r])

		if right-left+1-maxFrequency <= k {
			result = common.MaxInt(result, right-left+1)
			continue
		}

		// shrunk the window
		o := s[left]
		frequency[o]--
		left++

		//result = common.MaxInt(result, right-left+1)
	}

	return result
}
