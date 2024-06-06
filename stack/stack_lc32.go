package stack

/**
Wrong version
*/

func longestValidParentheses(s string) int {

	result, solution := 0, 0

	stack := make([]rune, 0)

	for _, str := range s {

		if str == '(' {
			stack = append(stack, str)
			continue
		}
		if len(stack) == 0 {
			continue
		}

		// then, str should be ')'
		pre := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre == '(' {
			solution += 2
			if result < solution {
				result = solution
			}
		} else {
			solution = 0
		}
	}

	return result
}

/**
Wrong version. Same as original version
*/

func longestValidParenthesesV2(s string) int {

	result, solution := 0, 0
	stack := NewByteStack()

	for _, o := range s {

		if o == '(' {
			stack.Push(byte(o))
			continue
		}

		if stack.Size() == 0 {
			continue
		}

		pre := stack.Pop()

		if pre == '(' {
			solution += 2
			if result < solution {
				result = solution
			}
		} else {
			solution = 0
		}
	}

	return result
}

/**
Tricky version. Hard to have this idea
KP
	1.	T: O(n)
	2.	S: O(n)
*/

func longestValidParenthesesV3(s string) int {

	result := 0

	stack := NewIntStack()
	stack.Push(-1)
	for i, o := range s {

		if o == '(' {
			stack.Push(i)
			continue
		}

		stack.Pop()

		if stack.Size() == 0 {
			stack.Push(i)
		} else {
			pre := stack.Pop()
			newSolution := i - pre
			if result < newSolution {
				result = newSolution
			}
			stack.Push(pre)
		}
	}
	return result
}

/**
Tricky version. More Harder to have this idea than V3
KP
	1.	T: O(n)
	2.	S: O(1), better than V3. With only 2 vars used.
*/

func longestValidParenthesesV4(s string) int {

	result := 0

	left, right := 0, 0

	for _, o := range s {
		if o == '(' {
			left += 1
		} else {
			right += 1
		}

		if left == right && result < 2*left {
			result = 2 * left
		}

		if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0

	for i := len(s) - 1; i >= 0; i -= 1 {

		o := s[i]

		if o == ')' {
			right += 1
		} else {
			left += 1
		}

		if left == right && result < 2*right {
			result = 2 * right
		}

		if left > right {
			left, right = 0, 0
		}
	}
	return result
}

/**
Normal Version. Dynamic Programming. (Similar as backtrack)
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

func longestValidParenthesesV3PV1(s string) int {
	result := 0
	stack := NewIntStack()
	stack.Push(-1)

	for i, o := range s {

		if o == '(' {
			stack.Push(i)
			continue
		}

		stack.Pop()

		if stack.Size() == 0 {
			stack.Push(i)
		} else {
			pre := stack.Pop()
			newAnswer := i - pre
			if result < newAnswer {
				result = newAnswer
			}
			stack.Push(pre)
		}
	}

	return result
}
