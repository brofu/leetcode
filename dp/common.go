package dp

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinIntTrix(a, b, c int) int {
	if a > b {
		return MinInt(b, c)
	}
	return MinInt(a, c)
}
