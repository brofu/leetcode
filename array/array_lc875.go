package array

func minEatingSpeed(piles []int, h int) int {

	left, right := 1, 1000000000

	for left <= right {
		mid := left + (right-left)/2
		hour := timeConsumed(piles, mid)

		if hour == h {
			right = mid - 1 // shrunk the right part
			continue
		}

		if hour < h { // eat slower
			right = mid - 1
			continue
		}

		if hour > h { // eat faster
			left = mid + 1
		}
	}

	// no need check if left > 100000000
	// there must be answer
	return left
}

func timeConsumed(piles []int, speed int) int {
	result := 0
	for _, pile := range piles {
		result += pile / speed
		if pile%speed != 0 {
			result += 1
		}
	}
	return result
}
