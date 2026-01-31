package stack

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))
	s := [][]int{}

	for i := len(temperatures) - 1; i >= 0; i-- {

		for len(s) > 0 && s[len(s)-1][0] <= temperatures[i] {
			s = s[:len(s)-1] // pop
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1][1] - i // peak
		}
		// push
		s = append(s, []int{temperatures[i], i})
	}

	return res
}

/*
KP:
	SC better.
	TC same O(N), but faster.

	Why?
	```
	s = append(s, []int{temperatures[i], i})
	```
	每次循环都会创建一个新的 []int（长度 2 的 slice），通常会逃逸到堆上 → 大量小对象分配 → GC 频繁。
	whie
	```
	s = append(s, i)
	```
	只是往一个 []int 里追加一个 int（通常只在扩容时分配）→ 大部分迭代 0 分配。

	Why there would be escape and why there would GC pressure?
	* 在 Go 里，“逃逸到堆上”不是因为对象小不小，而是因为 它的生命周期必须超过当前函数栈帧（或当前迭代），编译器就只能把它放到 heap，保证引用一直有效。
	* s = append(s, []int{temperatures[i], i}) ==> 这里发生了两件事：
		a. []int{a, b} 会创建一个 slice（header）+ 一个 长度为 2 的底层数组 来存 a,b。 底层数组在哪里分配，取决于它会不会“逃逸”。
		b. 把这个 slice 存进了 s [][]int，生命周期跨越了当前迭代. s 是函数内的变量，但它在整个循环过程中都要保留所有 push 进去的元素。
		   也就是说：你这次迭代创建的那块 [2]int 底层数组，必须在下一次迭代、甚至循环结束后仍然存在，因为 s 里存着指向它的引用。
		   如果编译器把这块底层数组放在当前迭代的栈上，那么迭代结束（更准确：栈帧里的临时对象生命周期结束）后，这块内存就可能被复用/覆盖，s 里就会变成悬空引用 —— 这是不允许的。
		   所以编译器会判断：这个底层数组被保存到了一个“可能活得更久”的地方（s 里），因此它 escapes，必须分配到 heap。
	* 为什么会GC造成压力？		
		因为循环 n 次就会创建 n 个“小 slice + 小数组”的堆对象（很多时候是两类对象）：
		每次 []int{...} 的底层数组（很可能在堆上）,还会产生大量短命对象 → GC 要扫描/标记/回收，成本就上来了. 即使每个对象都很小，对象数量一大，GC 也会明显变慢。
	* 对比：为什么 []int 存下标就不会这样？
	  s := []int{}
	  s = append(s, i)
	  这里只有一个 []int 的 backing array（s 自己的那块连续内存）, 它扩容次数很少（按 cap 成倍增长） 
	  不会每次循环都创建一个新的“小对象”, 
	  所以堆分配和 GC 压力都大幅降低。

	When there would escape?
	只要你在 Go 里做了类似事情，基本就会逃逸：
	把某个临时对象（或它的底层数据）的引用保存到一个“更长寿”的容器里（slice/map/channel/返回值/闭包捕获等）
	编译器就必须让它活得足够久 → 逃逸到堆
*/

func dailyTemperaturesV2(temperatures []int) []int {

	res := make([]int, len(temperatures))
	s := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {

		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1] // pop
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i // peak
		}
		// push
		s = append(s, i)
	}

	return res
}
