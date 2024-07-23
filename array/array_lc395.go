package array

import (
	"fmt"

	"github.com/brofu/leetcode/common"
)

func longestSubstring(s string, k int) int {

	result := 0

	helper := func(count int) int {

		result, valid := 0, 0
		left := 0

		records := make(map[byte]int)
		uniqueCount := 0

		for right := 0; right < len(s); right++ {

			// extend the window
			r := s[right]
			if _, ok := records[r]; !ok {
				uniqueCount++
			}
			records[r]++
			if records[r] == k {
				valid++
			}

			// shrunk the window
			for uniqueCount > count {
				o := s[left]
				if records[o] == k {
					valid--
				}
				records[o]--
				if records[o] == 0 {
					delete(records, o)
					uniqueCount--
				}
				left++
			}

			// check if we have a valid result
			if valid == uniqueCount {
				fmt.Println(count, right, left)
				result = common.MaxInt(result, right-left+1)
			}
		}
		return result
	}

	for i := 1; i <= 26; i++ {
		temp := helper(i)
		result = common.MaxInt(result, temp)
	}

	return result
}
