package array

func shipWithinDays(weights []int, days int) int {
	left, right := 1, 0

	for _, w := range weights {
		right += w
		if left < w {
			left = w
		}
	}

	for left <= right {
		mid := left + (right-left)/2
		d := daysConsumed(weights, mid)

		if d == days {
			right = mid - 1
			continue
		}

		if d < days {
			right = mid - 1
			continue
		}

		if d > days {
			left = mid + 1
		}
	}

	return left
}

func daysConsumed(weights []int, capability int) int {
	result := 0

	for i := 0; i < len(weights); {
		c := capability
		for i < len(weights) {
			if weights[i] > c {
				break
			}
			c -= weights[i]
			i++
		}
		result += 1
	}

	return result
}
