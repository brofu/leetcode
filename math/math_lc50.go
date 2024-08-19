package math

import "math"

func myPow(x float64, n int) float64 {

	if n == 0 { // base case
		return 1.0
	}

	if n == math.MinInt32 { // KP. In case of overflow
		return myPow(1/x, -(n+1)) / x
	}

	if n < 0 {
		return myPow(1/x, -n)
	}

	if n%2 == 1 {
		return myPow(x, n-1) * x
	} else { // KP. Time complexity reduces to O(lgN), but the space complexity would be increased to O(lgN)
		sub := myPow(x, n/2)
		return sub * sub
	}
}
