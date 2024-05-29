package stack

import (
	"strconv"
	"unicode"
)

func calculateBasic(s string) int {

	result := 0
	lastOperation := '+'
	numberStartIndex := -1
	stack := NewIntStack()

	for index := 0; index < len(s); index++ {

		o := s[index]

		if unicode.IsDigit(rune(o)) { // if operand

			if numberStartIndex == -1 {
				numberStartIndex = index
			}
			if index < len(s)-1 && unicode.IsDigit(rune(s[index+1])) {
				continue
			}

			operand, _ := strconv.Atoi(s[numberStartIndex : index+1])

			switch lastOperation {
			case '-':
				operand = -operand
			case '*':
				lastOperand := stack.Pop()
				operand *= lastOperand
			case '/':
				lastOperand := stack.Pop()
				operand = lastOperand / operand
			}

			stack.Push(operand)

			numberStartIndex = -1 // restore the index marker
		} else {

			if o == '+' || o == '-' || o == '*' || o == '/' {
				lastOperation = rune(o)
			}
		}
	}

	for stack.Size() > 0 {
		result += stack.Pop()
	}

	return result
}

/**
KP.
	1.	Handle the logic at next character after the operand.
		a.	Need pay attention to the end of the string.
			1.	If the end of the string is ' ',
			2.	If the end of the string is digit.
		b.	Golang is different from Python
*/
func calculateBasicV2(s string) int {

	currentOperand, result := 0, 0
	lasteOperator := '+'
	stack := NewIntStack()

	for index, o := range s {

		if o == ' ' && index < len(s)-1 {
			continue
		}

		if unicode.IsDigit(o) {
			currentOperand = currentOperand*10 + int(o-48)
			if index < len(s)-1 {
				continue
			}
		}

		switch lasteOperator {
		case '-':
			currentOperand = -currentOperand
		case '*':
			lastOperand := stack.Pop()
			currentOperand *= lastOperand
		case '/':
			lastOperand := stack.Pop()
			currentOperand = lastOperand / currentOperand
		}
		stack.Push(currentOperand)

		lasteOperator = o
		currentOperand = 0
	}

	for stack.Size() > 0 {
		result += stack.Pop()
	}

	return result
}
