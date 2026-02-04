package math

// quick calculate a^b%k
// what if b < 0?
func QuickPow(a, b, k int) int64 {
	res := int64(1)

	for b > 0 {
		if b&1 == 1 {
			res = (res * int64(a)) % int64(k)
		}
		a = (a * a) % k
		b >>= 1
	}

	return res
}

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return gcd(a, b)
}

func gcd(a, b int) int {

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func GCDV2(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

/*
KP:
	1. lcm(a, b) = a * b / gcd(a, b)
	2. Take care of number flow
*/
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
