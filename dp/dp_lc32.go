package dp

/**
Normal Version. Dynamic Programming. (Similar as backtrack)

Reference: https://leetcode.com/problems/longest-valid-parentheses/editorial/

KP
	1.	T:	O(n)
	2.	S:	O(n), with DP slice

*/

func longestValidParenthesesV5(s string) int {
	dp := make([]int, len(s))
	result := 0

	for i := 1; i < len(s); i++ {
		o := s[i]

		if o == '(' {
			continue
		}

		// s[i] == ')'
		// s[i-1] == '('
		if s[i-1] == '(' {
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else { // i==1
				dp[i] = 2
			}
		} else { // s[i-1] == ')'
			// like ((..))
			if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 { // like .((..))
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else { // like ((..))
					dp[i] = dp[i-1] + 2
				}
			}
		}

		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}
