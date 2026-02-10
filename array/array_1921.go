package array

import (
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {

	timeConsuming := make([]float32, len(dist))
	for idx, d := range dist {
		timeConsuming[idx] = float32(d) / float32(speed[idx])
	}

	sort.Slice(timeConsuming, func(x, y int) bool {
		return timeConsuming[x] < timeConsuming[y]
	})

	count := 0
	for timeNeed, timeArrive := range timeConsuming {
		if float32(timeNeed) < timeArrive {
			count++
		} else {
			break
		}
	}
	return count
}
