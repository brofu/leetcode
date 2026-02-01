package stack

func finalPrices(prices []int) []int {

	reslut := make([]int, len(prices))
	st := make([]int, 0)

	for i := len(prices) - 1; i >= 0; i-- {
		for len(st) > 0 && prices[i] < st[len(st)-1] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			reslut[i] = prices[i]
		} else {
			reslut[i] = prices[i] - st[len(st)-1]
		}
		st = append(st, prices[i])
	}

	return reslut
}
