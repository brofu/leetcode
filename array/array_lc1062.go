package array

import (
	"fmt"

	"github.com/brofu/leetcode/common"
)

func longestRepeatingSubstring(s string) int {

	result := 0
	data := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			data[s[i:j+1]] += 1
		}
	}

	for k, v := range data {
		if v > 1 {
			result = common.MaxInt(result, len(k))
		}
	}

	return result
}

/*
*
KP.
 1. Idea of use possible length as the 1st loop.
 2. The edged case
 3. Break the loop in advance
*/
func longestRepeatingSubstringV2(s string) int {

	result := 0
	data := make(map[string]struct{})

	for i := 2; i < len(s); i++ { // KP. i means the possible lengths
		for j := 0; j <= len(s)-i; j++ { // KP. the bounder of j
			sub := s[j : j+i]
			fmt.Println("flag", sub)
			if _, ok := data[sub]; ok {
				result = i
				break
			}
			data[sub] = struct{}{}
		}

		if result != i { // KP. Stop the loop in advance
			break
		}
	}
	return result
}

/*
*
KP.
 1. Check the possible length with binary search idea
*/
func longestRepeatingSubstringV3(s string) int {

	data := make(map[string]struct{})

	search := func(length int) bool {
		for i := 0; i <= len(s)-length; i++ {
			sub := s[i : i+length]
			if _, ok := data[sub]; ok {
				return true
			} else {
				data[sub] = struct{}{}
			}
		}
		return false
	}

	low, high := 1, len(s)

	for low <= high { // KP. User binary search ideas for the length checking
		mid := (low + high) / 2
		if search(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low - 1
}

// wrong
func longestRepeatingSubstringV4(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 1
	result := 0

	for ; right < len(s); right++ {

		if s[right] == s[right-1] {
			continue
		}

		if result < right-left {
			result = right - left
		}

		left = right
	}

	return result
}
