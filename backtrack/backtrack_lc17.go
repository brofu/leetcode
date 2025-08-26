package backtrack

import "strings"

/*
Key Points

TC:
1. The recursive deepth is N
2. At each recursive there is 3 choices. So, totally, there is 3^N
3. For each result, we need to join the letters together, which is around N
4. So, over all the TC is 3^N*N

SC:
1. The recursive deepth is N
2. For each result, we need a N-length string to store the record, totally 3^N*N space for this
3. We need around 26 to store the d map
4. Over all the SC is 3^N*N
*/

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	d := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var bt func(int, []string)

	bt = func(idx int, track []string) {

		// base case
		if idx == len(digits) {
			temp := strings.Join(track, "")
			result = append(result, temp)
			return
		}

		for _, str := range d[rune(digits[idx])] {
			// choose
			// next layer
			bt(idx+1, append(track, str))
			// cancel
		}
	}

	bt(0, []string{})
	return result
}
