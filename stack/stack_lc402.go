package stack

/*
KP:
	1. How to get the smallest?
		1. 先删除 num 中的若干数字，使得 num 从左到右每一位都单调递增。比如 14329 转化成 129，这需要使用到单调栈技巧。
		2. num 中的每一位变成单调递增的之后，如果 k 还大于 0（还可以继续删除）的话，则删除尾部的数字，比如 129 删除成 12。
	2. Pay attention to edged cases
*/

func removeKdigits(num string, k int) string {

	st := make([]rune, 0)

	for _, n := range num {

		for len(st) > 0 && st[len(st)-1] > n && k > 0 {
			st = st[:len(st)-1]
			k--
		}
		if len(st) == 0 && n == '0' { // in case 0 occurs at the beginning of st
			continue
		}
		st = append(st, n)
	}

	for k > 0 && len(st) > 0 {
		st = st[:len(st)-1]
		k--
	}

	if len(st) == 0 {
		return "0"
	}

	return string(st)
}

/*
KP:
	1. A faster version.
	2. Why?
		```
		if len(st) == 0 && n == '0' {
			    continue
		}
		```
		这意味着在扫描每个字符时都要判断一次：
		len(st) == 0（访问 slice header）
		n == '0'（比较）
		以及可能走 continue,continue 会改变“栈增长节奏”，间接影响 append 的行为
		遇到前导 0 会跳过 append，导致 st 的长度增长更不规则。虽然总体分配不一定更多，但在 Go 的运行时/编译器优化层面，这种“不稳定路径”通常更难被优化。

		```
		i := 0
		for ; i < len(st) && st[i] == '0'; i++ {}
		```
		V2 把“去前导 0”变成一次线性扫描，且只做一次切片
*/
func removeKdigitsV2(num string, k int) string {

	st := make([]rune, 0)

	for _, n := range num {

		for len(st) > 0 && st[len(st)-1] > n && k > 0 {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, n)
	}

	for k > 0 && len(st) > 0 {
		st = st[:len(st)-1]
		k--
	}

	i := 0
	for ; i < len(st) && st[i] == '0'; i++ {
	}
	if len(st[i:]) == 0 {
		return "0"
	}

	return string(st[i:])
}
