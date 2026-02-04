package math

/*
KP:
	1. How to make sure a number is prime?
	2. 素数筛选法
	3. 边界的控制
*/
func countPrimes(n int) int {

	if n == 2 {
		return 1
	}
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ { // check all the factors of n
		if !isPrime[2] {
			continue
		}
		for m := i * i; m <= n; m += i { // why start from i*i?
			isPrime[m] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
