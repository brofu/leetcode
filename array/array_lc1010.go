package array

func numPairsDivisibleBy60(time []int) int {

	result := 0

	data := make(map[int]int)
	for index, val := range time {
		time[index] = val % 60
	}
	for _, val := range time {
		data[val] += 1
	}

	for k, v := range data {
		if 0 < k && k < 30 {
			other, _ := data[60-k]
			result += v * other
		}
		if k == 30 || k == 0 {
			result += v * (v - 1) / 2
		}
	}

	return result
}
