package stack

import "strconv"

func evalRPN(tokens []string) int {

	st := make([]int, 0)

	for _, t := range tokens {
		if t == "+" {
			o1, o2 := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, o1+o2)
			continue
		}
		if t == "-" {
			o1, o2 := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, o1-o2)
			continue
		}
		if t == "*" {
			o1, o2 := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, o1*o2)
			continue
		}
		if t == "/" {
			o1, o2 := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			st = append(st, o1/o2)
			continue
		}

		num, _ := strconv.Atoi(t)
		st = append(st, num)
	}

	return st[0]
}
