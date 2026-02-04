package math

import "github.com/brofu/leetcode/common"

/*
KP
	1. Construct the urg array and util it at the same time.
	2. Key Observe:
		if x is an ugly number, then, x*2, x*3, and x*5 would be also ugly number
	3. The idea of merge 3 linked list

*/

func nthUglyNumber(n int) int {

	ugly := make([]int, n+1)

	p := 1
	p2, p3, p5 := 1, 1, 1
	pd2, pd3, pd5 := 1, 1, 1

	for p <= n {
		m := common.MinIntThree(pd2, pd3, pd5)
		ugly[p] = m
		p++

		if m == pd2 {
			pd2 = ugly[p2] * 2
			p2++
		}

		if m == pd3 {
			pd3 = ugly[p3] * 3
			p3++
		}

		if m == pd5 {
			pd5 = ugly[p5] * 5
			p5++
		}

	}

	return ugly[n]
}
