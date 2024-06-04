package parentheses

/**
KP
	1.	when got a `(`, need 2 more `)`. And, if the number of `)` is odd, need another one of it, insert it and reduce 1 from the need number of it
	2.	when got a `)`, reduce 1 from the number of `)`, and if the number of it is -1, it means there is 1 `)`, which is NOT mached. then, we need another `)` and another `(`
*/
func minInsertions(s string) int {
	insert, need := 0, 0

	for _, o := range s {

		if o == '(' {
			insert += 2
			if insert%2 == 1 { //? hmm why?
				need++
				insert--
			}
		}

		if o == ')' {
			insert--
			if insert == -1 {
				insert = 1 // clear needRight and
				need += 1  // add needLeft
			}
		}
	}

	return need + insert

}

func minInsertionsPV1(s string) int {
	insert, need := 0, 0

	for _, o := range s {
		if o == '(' {
			need += 2
			if need%2 == 1 {
				insert += 1 // insert a ")"
				need -= 1   // and reeduce the need number
			}
		}

		if o == ')' {
			need -= 1
			if need == -1 {
				need = 1
				insert += 1 // insert a "("
			}
		}
	}

	return insert + need
}
