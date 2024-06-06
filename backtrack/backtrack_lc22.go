package backtrack

/**
KP
	1.	Principles
		* 一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解。
		* 对于一个「合法」的括号字符串组合 p，必然对于任何 0 <= i < len(p) 都有：子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。
	2.	问题转换		
		现在有 2n 个位置，每个位置可以放置字符 ( 或者 )，组成的所有括号组合中，有多少个是合法的？
	3.	Skill
		Use integer `left`, `right` to mark the left number of `(` or `)`
*/

func generateParenthesis(n int) []string {

	result := []string{}
	track := ""

	var bt func(int, int, *string, *[]string)

	bt = func(left, right int, track *string, result *[]string) {
		if right < 0 || left < 0 {
			return
		}

		if right < left {
			return
		}

		// ending of choose
		if right == 0 && left == 0 {
			*result = append((*result), *track)
			return
		}

		// choose "("
		*track += "("
		// depth + 1
		bt(left-1, right, track, result)
		// cancel choose
		*track = (*track)[:len(*track)-1]

		// choose ")"
		*track += ")"
		// depth + 1
		bt(left, right-1, track, result)
		// cancel choose
		*track = (*track)[:len(*track)-1]

	}

	bt(n, n, &track, &result)
	return result
}
