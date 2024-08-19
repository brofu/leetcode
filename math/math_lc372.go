package math

func superPow(a int, b []int) int {

	power := func(x, n, base int) int { //KP. Mod base in case of `overflow`
		x %= base
		mul := 1
		for i := 0; i < n; i++ {
			mul *= x
			mul %= base
		}
		return mul
	}

	var helper func(a int, b []int, base int) int

	helper = func(a int, b []int, base int) int {
		if len(b) == 0 {
			return 1
		}

		if a == 1 {
			return 1
		}

		tmp := b[len(b)-1]
		b = b[:len(b)-1]

		num1 := power(a, tmp, base)
		num2 := power(helper(a, b, base), 10, base)
		//num1 := fastPow(a, tmp, base)
		//num2 := fastPow(helper(a, b, base), 10, base)

		return (num1 % base * num2 % base) % base
	}

	return helper(a, b, 1337)
}

// fastPow fast version of power. Utilize the bit operation
func fastPow(a, b, base int) int {
	if a == 1 {
		return 1
	}

	a %= base
	res := 1

	for b != 0 { // KP. Utilize `bit operation` for power calculation.
		if b&1 != 0 {
			res = (res * a) % base
		}

		a = (a * a) % base
		b = b >> 1
	}

	return res
}

/**
The iteration version.
*/
func superPowV2(a int, b []int) int {

	res, base := 1, 1337

	if len(b) == 0 {
		return res
	}

	for i := 0; i < len(b); i++ { // KP. The `Iteration` version to calculate Power
		res = fastPow(res, 10, base)
		if b[i] > 0 {
			res = res * fastPow(a, b[i], base) % base
		}
	}

	return res
}
