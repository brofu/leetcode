package stack

import "sort"

/*
KP:
	1. Get the main idea:
	是否能够形成车队，取决于下:如果车 x 排在 车 y 后面，且 x 到达终点所需时间比 y 少，则 x 必然会被 y 卡住，形成车队。
*/
func carFleet(target int, position []int, speed []int) int {

	orderP := make([][]int, len(position))
	timing := make([]float64, len(orderP))
	for idx, p := range position {
		orderP[idx] = []int{p, speed[idx]}
	}
	sort.Slice(orderP, func(x, y int) bool {
		return orderP[x][0] < orderP[y][0]
	})

	for idx, p := range orderP {
		timing[idx] = float64((target - p[0])) / float64(p[1])
	}

	st := make([]float64, 0)

	for i := 0; i < len(timing); i++ {
		for len(st) > 0 && timing[i] >= st[len(st)-1] {
			st = st[:len(st)-1]
		}
		st = append(st, timing[i])
	}

	return len(st)
}

/*
KP
	1. A version No need sort and monotonic stack
	2. Use the [target]int, to setup the order
*/

func carFleetV2(target int, position []int, speed []int) int {

	timing := make([]float64, target)

	for idx, p := range position {
		timing[p] = float64((target - p)) / float64(speed[idx])
	}

	result := 0
	previous := 0.0
	for i := len(timing) - 1; i >= 0; i-- {
		if timing[i] > previous {
			result++
			previous = timing[i]
		}
	}
	return result
}
