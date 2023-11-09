package dp

func fib(n int) int {
	mem := make([]int, n+1)
	return fibdp(n, mem)
}

func fibdp(n int, mem []int) int {
	if n == 0 || n == 1 {
		return n
	}

	number := mem[n]
	if number == 0 {
		number = fibdp(n-1, mem) + fibdp(n-2, mem)
		mem[n] = number
	}
	return number
}
