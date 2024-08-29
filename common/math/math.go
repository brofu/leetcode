package math

import "math"

/**
Utilize `Bit Operation` to calculate `Power`

KP.
	1.	Utilize `Bit Operation` to reduce the time complexity. O(logN)
	2.	If `b < 0`, need to covert the calculation to `Power(1/a, -b)`. And,
		a.	An edged case is, `b = math.MinInt64`. For this situation,  there would be overflow of `-b`
*/

func MyPower(a float64, b int) float64 {

	if b < 0 {
		if b > math.MinInt64 {
			return fastPower(1/a, -b)
		}
		return fastPower(1/a, -(b+1)) / a
	}
	return fastPower(a, b)
}

func fastPower(a float64, b int) float64 {

	if a == 1 {
		return a
	}

	result := 1.0

	for b != 0 {
		if b&1 != 0 {
			result *= a
		}
		a *= a
		b >>= 1
	}

	return result
}

/**
Multi Version of `Power` implementation.

KP.
	1.	Reduce the complexity
*/

func MyPowerV2(a float64, b int) float64 {

	if a == 1 {
		return a
	}

	if b == 0 {
		return 1
	}

	if b == math.MinInt64 {
		return MyPowerV2(1/a, -(b+1)) / a
	}

	if b < 0 {
		return MyPowerV2(1/a, -b)
	}

	if b%2 == 1 {
		return MyPowerV2(a, b-1) * a
	} else {
		sub := MyPowerV2(a, b/2)
		return sub * sub
	}
}
