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

func fibRecurrence(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	mem := make([]int, n+1)
	mem[1] = 1
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}
