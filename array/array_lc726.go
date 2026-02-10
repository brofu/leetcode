package array

import (
	"fmt"
	"sort"
)

/*
KP:
	1. How to handle the parathes? Stock
*/
func countOfAtoms(formula string) string {

	st := make([]int, 0)
	pidx := make(map[int]int)
	for idx, v := range formula {
		if v == '(' {
			st = append(st, idx)
		}
		if v == ')' {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			pidx[p] = idx
		}
	}
	resultMap := countOfAtomsHelper(formula, 0, len(formula)-1, pidx)

	fmt.Println("flag", resultMap)
	result := ""

	order := make([]string, 0)
	for k := range resultMap {
		order = append(order, k)
	}
	sort.Strings(order)

	for _, s := range order {
		count := resultMap[s]
		if count == 0 {
			continue
		}
		if count == 1 {
			result += s
			continue
		}
		result += fmt.Sprintf("%s%d", s, count)
	}

	return result
}

func countOfAtomsHelper(formula string, start, end int, pidx map[int]int) map[string]int {

	result := make(map[string]int)
	for i := start; i <= end; {
		if formula[i] == '(' {
			peidx := pidx[i]
			t := countOfAtomsHelper(formula, i+1, peidx-1, pidx)
			j, num := peidx+1, 0
			for ; j <= end && formula[j] >= '0' && formula[j] <= '9'; j++ {
				num = num*10 + int(formula[j]-48)
			}
			if num == 0 {
				num = 1
			}
			for k, v := range t {
				result[k] += num * v
			}
			i = j
			continue
		}
		if formula[i] == ')' || (formula[i] >= '2' && formula[i] <= '9') {
			continue
		}
		j, num := i+1, 0
		for ; j <= end && (formula[j] >= 'a' && formula[j] <= 'z'); j++ {
		}
		letter := formula[i:j]

		for ; j <= end && formula[j] >= '0' && formula[j] <= '9'; j++ {
			num = num*10 + int(formula[j]-48)
		}
		if num == 0 {
			num = 1
		}
		result[letter] += num
		i = j
	}

	return result
}
