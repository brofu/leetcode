package parentheses

func minInsertions(s string) int {
	needRight, needLeft := 0, 0

	for _, o := range s {

		if o == '(' {
			needRight += 2
			if needRight%2 == 1 { //? hmm why?
				needLeft++
				needRight--
			}
		}

		if o == ')' {
			needRight--
			if needRight == -1 {
				needRight = 1 // clear needRight and
				needLeft += 1 // add needLeft
			}
		}
	}

	return needLeft + needRight

}
