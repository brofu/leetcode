package math

import "math"

/**
KP.
	1.	Only the lights which are pressed odd number times is ON.
	2.	Which lights would be pressed ODD number times?
		a.	Only the ones which are the square of another number, for example, 1, 4, 9, 16...
		b.	Because all the other numbers would be pressed by EVEN number times. For example,
			a.	2 = 1 * 2, would be pressed by the 1st, 2nd round.
			b.	3 = 1 * 3, would be pressed by the 1st, 3rd round.
			c.  6 = 1 * 6 = 2 * 3, would be pressed by the 1st, 2nd, 3rd, and 6th round.
			d.	And so on.
	3.	According to point 2, the problem is actually about the number of different factors of a number
*/

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func bulbSwitchV2(n int) int {
	result := 0
	for i := 1; i*i <= n; i++ {
		result++
	}
	return result
}
