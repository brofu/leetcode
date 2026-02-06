package math

/*
KP:
	https://labuladong.online/zh/algo/frequency-interview/factorial-problems/

	1. Got the idear that, how many factors of `5` are there in the n!
*/
func trailingZeroes(n int) int {

	res := 0
	for divisor := 5; divisor <= n; {
		res += n / divisor
		divisor *= 5
	}

	return res
}
