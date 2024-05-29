package stack

import (
	"fmt"
	"strconv"
	"unicode"
)

/**
KP.
	1.	Evaluate the expression at the next coming charactor. so, need to a padded charactor at the end of the string. Here we used `@`
	2.	The charactor is pushed into stack only there is `parentheses`
		a.	So, when the `)` is entercounted, just evaluate all `operands` in the stack, until there is an `operator`
	3.	Pay attention to the stack type. If use `rune`, the middle result may equal to operators `+`, `-`, `*`, `/`, which would affect the result,
		So, we use string here.
*/
func calculateWithParenthese(s string) int {

	result, currentOperand := 0, 0
	stack := NewStringStack()
	previousOperator := '+'
	s1 := s + "@"

	evaluate := func(operator rune, source, dest int) int {
		switch operator {
		case '-':
			return 0 - dest
		case '*':
			return source * dest
		case '/':
			return source / dest
		default:
			return source + dest
		}
	}

	for _, o := range s1 {

		if o == ' ' {
			continue
		}

		if unicode.IsDigit(o) {
			currentOperand = currentOperand*10 + int(o-48)
			continue
		}

		if o == '(' {
			stack.Push(string(previousOperator))
			previousOperator = '+'
			continue
		}

		switch previousOperator {
		case '-':
			currentOperand = evaluate(previousOperator, 0, currentOperand)
		case '*', '/':
			previousOperand, _ := strconv.Atoi(stack.Pop())
			currentOperand = evaluate(previousOperator, previousOperand, currentOperand)
		}

		stack.Push(fmt.Sprintf("%d", currentOperand))
		previousOperator = o
		currentOperand = 0

		if o == ')' {
			for stack.Size() > 0 {
				stacked := stack.Pop()
				if stacked == "+" || stacked == "-" || stacked == "*" || stacked == "/" {
					previousOperator = rune(stacked[0])
					break
				}
				stackedOperand, _ := strconv.Atoi(stacked)
				currentOperand += stackedOperand
			}
		}

	}

	for stack.Size() > 0 {
		stackedOperand, _ := strconv.Atoi(stack.Pop())
		result += stackedOperand
	}

	return result
}
