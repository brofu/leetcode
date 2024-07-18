package array

func similarPairs(words []string) int {

	result := 0
	data := make(map[int]int)

	for _, word := range words {
		data[toBits(word)]++
	}
	for _, v := range data {
		result += v * (v - 1) / 2
	}

	return result
}

func toBits(word string) int {
	res := 0
	for _, r := range word {
		res |= 1 << (r - 'a')
	}
	return res
}
