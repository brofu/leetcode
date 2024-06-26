package stack

/**
KP
	1.	It's invalid if current is `)`, but stack.Size() == 0
	2.	Need to check if all the elements in stack are matched
*/

func isValid(s string) bool {

	stack := NewByteStack()

	preof := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, o := range s {
		if o == '{' || o == '[' || o == '(' {
			stack.Push(byte(o))
			continue
		}

		if stack.Size() == 0 { // current is right parentheses, but NO left one
			return false
		}

		pre := stack.Pop()
		expectedPre := preof[byte(o)]
		if pre != expectedPre {
			return false
		}
	}

	// check if all the left parentheses are matched
	return stack.Size() == 0
}
