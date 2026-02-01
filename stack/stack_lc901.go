package stack

type StockSpanner struct {
	prices []int
	st     []int
}

func Constructor906() StockSpanner {
	return StockSpanner{}
}

/*
SC:
	1. O(N). 调用N次next, O(N^2)
*/
func (this *StockSpanner) Next(price int) int {

	this.prices = append(this.prices, price)

	count := 0
	for i := 0; i < len(this.prices); i++ {
		if this.prices[i] <= price {
			count++
			continue
		}
		count = 0
	}

	return count
}

/*
KP:
	1. Monotonic Stock

SC:
	1. 均摊O(1), 调用N次next, O(N)
	2. How to understand this?
		单次 Next 不保证最坏是 O(1)。 单次最坏：O(n)（比如这次 price 很大，把栈全弹空）
		但摊还/均摊：O(1)（做 n 次调用，总时间 O(n)）
		这就是单调栈经典的“摊还 O(1)”场景。

		为什么能摊还到 O(1)？
		关键观察：每个元素最多被 push 一次、pop 一次。 每次 Next(price) 都会把 price 入栈一次（push）, while 循环里，每弹出一个旧 price，它以后就不会再回到栈里（pop 只能发生一次）
		所以，假设总共调用了 T 次 Next：
		push 总次数 = T, pop 总次数 ≤ T
		while 循环虽然在某一次调用里可能弹很多个，但把所有调用的弹栈次数加起来最多也就 T 次。
		因此总操作次数是线性的：
		Total=O(T)
*/

func (this *StockSpanner) NextV2(price int) int {

	a := 1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		a += this.st[len(this.st)-1]
		this.prices = this.prices[:len(this.prices)-1]
		this.st = this.st[:len(this.st)-1]
	}
	this.prices = append(this.prices, price)
	this.st = append(this.st, a)
	return a
}
