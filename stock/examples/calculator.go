package examples

import (
	"fmt"
	"regexp"

	"strconv"
)

type StockString struct {
	elements []string
	length   int
}

func (s *StockString) Push(c string) {
	s.length++
	s.elements = append(s.elements, c)
}

func (s *StockString) Pop() string {
	if s.length == 0 {
		panic("empty")
	}
	c := s.elements[s.length-1]
	s.length--
	s.elements = s.elements[0:s.length]
	return c
}

func (s *StockString) Debug() {
	fmt.Println("flag", s.elements)
}

func NewStockString() *StockString {
	return &StockString{elements: []string{}}
}

func IsNumber(str string) (bool, int64) {
	if number, err := strconv.ParseInt(str, 10, 64); err == nil {
		return true, number
	} else {
		return false, 0
	}
}

func CalculateExpression(expr string) int64 {

	re := regexp.MustCompile(`[\d.]|[\+\-\*\/\(\)]`)
	stock := NewStockString()

	tokens := re.FindStringSubmatch(expr)

	stock.Push(tokens[0])

	for index := 1; index < len(tokens); index++ {
		stock.Debug()

		previousToken := tokens[index-1]
		token := tokens[index]
		nextToken := ""
		if index < len(tokens) {
			nextToken = tokens[index+1]
		}

		if token == "(" {
			stock.Push(token)
			continue
		}

		if token == ")" {
			continue
		}

		if isNumber, number := IsNumber(token); isNumber {

			if previousToken == "(" {
				stock.Push(token)
				continue
			}

			if nextToken == "*" || nextToken == "/" {
				stock.Push(token)
				continue
			}

			if previousToken == "*" || previousToken == "/" || previousToken == "+" || previousToken == "-" {
				stock.Pop()
				operandOne := stock.Pop()
				sourceNumber, _ := strconv.ParseInt(operandOne, 10, 64)

				switch previousToken {
				case "*":
					stock.Push(fmt.Sprintf("%d", sourceNumber*number))
				case "/":
					stock.Push(fmt.Sprintf("%d", sourceNumber/number))
				case "+":
					stock.Push(fmt.Sprintf("%d", sourceNumber+number))
				case "-":
					stock.Push(fmt.Sprintf("%d", sourceNumber-number))
				}
				continue
			}

			stock.Push(token)
			continue
		}

		// Push operators
		stock.Push(token)
	}

	result, _ := strconv.ParseInt(stock.Pop(), 10, 64)
	return result
}
