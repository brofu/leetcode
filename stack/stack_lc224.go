package stack

import (
	"unicode"

	"github.com/brofu/leetcode/common"
)

/**

Sub Problems
	1.	How to convert a continues stream to a number?
	2.	How to handle ' ', in the string?
	3.	How to handle the parentheses?
*/

func calculate(s string) int {

	var helper func(string) int

	helper = func(s string) int {
		result, number := 0, 0
		op := '+'
		stack := NewIntStack()

		for len(s) > 0 {
			o := rune(s[0])
			s = s[1:]
			if unicode.IsDigit(o) {
				number = number*10 + int(o-'0')
			}

			if o == '(' {
				number = helper(s)
			}

			if !unicode.IsDigit(o) && o != ' ' || len(s) == 0 { // o is NOT digital

				if op == '-' {
					number = -number
				}

				stack.Push(number)
				op = o     // update op
				number = 0 // clear number
			}

			if o == ')' {
				break
			}
		}

		common.Debug(stack.Debug, "flag1", s)
		for stack.Size() != 0 {
			result += stack.Pop()
		}

		return result
	}

	return helper(s)
}
