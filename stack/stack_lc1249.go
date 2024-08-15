package stack

import "bytes"

/**
KP.
	1.	What's valid Parentheses?
		a. Number same, '(' and ')' and,
		b. There must '(', before ')', so, no leading ')', and no open '('

	2. Utilized stack.
		a.	Empty stack + ')', means leading ')'
		b.	Non-Empty stack and the end, means open '('
*/
func minRemoveToMakeValid(s string) string {

	stack := []int{}
	record := make(map[int]struct{})

	for i, r := range s {

		if r == '(' {
			stack = append(stack, i)
			continue
		}

		if r == ')' {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			} else { // need to remove this ')'
				record[i] = struct{}{}
			}
		}
	}

	for i := 0; i < len(stack); i++ { // need to remove all the left '('
		record[stack[i]] = struct{}{}
	}

	b := bytes.NewBuffer([]byte{})
	for i, r := range s {
		if _, ok := record[i]; !ok { // no need to remove
			b.WriteRune(r)
		}
	}

	return b.String()
}
