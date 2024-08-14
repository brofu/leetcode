package str

func index(a, b int) int {
	result := 1
	if b == 0 {
		return result
	}
	for i := 0; i < b; i++ {
		result *= a
	}

	return result
}

func WhatCentury(year string) string {

	strMap := map[int]string{
		1: "st",
		2: "nd",
		3: "rd",
	}
	yearNumber := 0
	for i, r := range year {
		yearNumber += index(10, 3-i) * int(r-48)
	}

	fs, ls := "", "th"

	f, l := yearNumber/100, yearNumber%100
	if l != 0 {
		f += 1
	}

	if f < 10 {
		fs = string(rune(f + 48))
		if v, ok := strMap[f]; ok {
			ls = v
		}
	}

	if f >= 10 && f < 100 {
		f1, f2 := f/10, f%10
		if f1 == 1 {
			ls = "th"
		} else {
			if v, ok := strMap[f2]; ok {
				ls = v
			}
		}
		fs = string(rune(f1+48)) + string(rune(f2+48))
	}

	if f == 100 {
		return "100th"
	}

	return fs + ls
}
