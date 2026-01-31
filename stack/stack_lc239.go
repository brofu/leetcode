package stack

func maxSlidingWindow(nums []int, k int) []int {

	var (
		result     []int
		q          = make([]int, 0, len(nums))
		start, end = 0, 0
	)

	for end < len(nums) {
		// extend the window
		if end-start < k {
			for len(q) > 0 && q[len(q)-1] < nums[end] {
				q = q[:len(q)-1]
			}
			q = append(q, nums[end])
			end++
		}
		if end-start == k {
			// collect result
			result = append(result, q[0])
			// shrink window
			if q[0] == nums[start] {
				q = q[1:]
			}
			start++
		}
	}
	return result
}

/*
KP:
	1. Use the `head` and `tail` pointer to control queue. To reduce the slice shrink/expand operation
		* 不改变 slice 的起始指针，append 扩容更少、更可控
		deq 的 ptr 永远指向同一个 backing array 的开头，你只是移动 head 这个整数。 
		append 的行为更稳定, 你可以把“逻辑上的队头”前面的空间看作可复用/可清理的
		* 可以在合适时机“重置/压缩”，真正释放/复用 
		```
		if head > 1024 {                 // 或 head > len(deq)/2
			copy(deq, deq[head:])        // 把有效区搬到前面
			deq = deq[:len(deq)-head]
			head = 0
		}
		```
	2. 工程参考：
		Use ring buffer, to control the memory mollac and GC pressure.
*/
func maxSlidingWindowV2(nums []int, k int) []int {

	var (
		n          = len(nums)
		result     []int
		q          = make([]int, n, n)
		start, end = 0, 0
		head, tail = 0, 0
	)

	for end < len(nums) {
		// extend the window
		if end-start < k {
			for tail > head && q[tail-1] < nums[end] {
				tail--
			}
			q[tail] = nums[end]
			tail++
			end++
		}
		if end-start == k {
			// collect result
			result = append(result, q[head])
			// shrink window
			//IMPORTANT: why use `head` pointer, instead of q = q[1:]?
			if q[head] == nums[start] {
				head++
			}
			start++
		}
	}
	return result
}
