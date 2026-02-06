package math

func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n = n >> 1
	}
	return res
}
