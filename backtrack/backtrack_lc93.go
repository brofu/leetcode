package backtrack

import (
	"strconv"
	"strings"
)

/*
Key Points:

1. Backtrack problem.
2. Set problem. Put the "." into the correct location, and at most 3 "."

TC:
1. Possible cutting positions set: C(n-1, 3). There are totally `n-1` positions to put 3 `.`,
2. But the actually result would be much less that this, because of the pruning. Actually, it would be roughly 3^4, (for each section, we can have 1, 2, or 3 digits)
3. For each possible result, need to copy, as O(n)
4. Checking if the number are valid IP parts, O(1) at each recursive layer, with totally 4 recursive depth.
5. Over all, O(C(n-1, 3)*n)

SC:
1. Recursive call depth: 4
3. Store result O(C(n-1, 3)*n)
*/
func restoreIpAddresses(s string) []string {

	result := make([]string, 0)

	var bt func(int, []string)

	bt = func(start int, trace []string) {

		// base case
		if start == len(s) && len(trace) == 4 {
			result = append(result, strings.Join(trace, "."))
			return
		}

		for idx := start; idx < len(s); idx++ {
			// pruning
			if !isValidIPSection(s[start : idx+1]) {
				continue
			}
			// choose
			// next layer
			bt(idx+1, append(trace, s[start:idx+1]))
			// cancel
		}
	}

	bt(0, []string{})

	return result
}

func isValidIPSection(input string) bool {
	if len(input) > 1 && input[0] == '0' {
		return false
	}
	if len(input) > 3 {
		return false
	}
	number, _ := strconv.Atoi(input)
	if number > 255 {
		return false
	}
	return true
}
