package math

/*
KP:
	https://labuladong.online/zh/algo/frequency-interview/factorial-problems/

*/
func trailingZeroes(n int) int {

	res := 0
	for divisor := 5; divisor <= n; {
		res += n / divisor
		divisor *= 5
	}

	return res
}
