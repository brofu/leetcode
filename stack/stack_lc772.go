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

/**
KP
	1.	How to control the cursive call-return point? Use the length of s. And need a pointer to s
	2.	Handl the edged case when ' ' or number is at the end of the string.
*/
func calculateWithParentheseRecursively(s string) int {

	var helper func(*string) int

	helper = func(s *string) int {
		result := 0
		stack := NewRuneStack()
		previousOperator := '+'
		currentOperand := 0
		*s = *s + "@"

		for len(*s) > 0 {

			o := (*s)[0]
			*s = (*s)[1:]

			if o == ' ' {
				continue
			}

			if unicode.IsDigit(rune(o)) {
				currentOperand = currentOperand*10 + int(o-'0')
				continue
			}

			if o == '(' {
				currentOperand = helper(s)
				continue
			}

			switch previousOperator {
			case '-':
				currentOperand = -currentOperand
			case '*':
				previousOperand := stack.Pop()
				currentOperand *= int(previousOperand)
			case '/':
				previousOperand := stack.Pop()
				currentOperand = int(previousOperand) / currentOperand
			}
			stack.Push(rune(currentOperand))
			previousOperator = rune(o)
			currentOperand = 0

			if o == ')' {
				break
			}
		}
		for stack.Size() > 0 {
			result += int(stack.Pop())
		}

		return result
	}

	return helper(&s)
}
