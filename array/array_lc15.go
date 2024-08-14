package array

import (
	"github.com/brofu/leetcode/common"
)

func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i <= j; {
		formatI, okI := common.CheckAndFormatAlphanumeric(s[i])
		if !okI {
			i++
		}
		formatJ, okJ := common.CheckAndFormatAlphanumeric(s[j])
		if !okJ {
			j--
		}

		if okI && okJ {
			if formatI != formatJ {
				return false
			}
			i++
			j--
		}
	}

	return true
}
