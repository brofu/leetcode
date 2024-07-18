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
