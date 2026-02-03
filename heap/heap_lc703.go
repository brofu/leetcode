package heap

type KthLargest struct {
	heap []int
	k    int
}

func Constructor703(k int, nums []int) KthLargest {

	instance := KthLargest{
		k: k,
	}
	for i := 0; i < len(nums); i++ {
		instance.Push(nums[i])
		if i >= k {
			instance.Pop()
		}
	}
	return instance

}

func (this *KthLargest) Add(val int) int {

	this.Push(val)
	if len(this.heap) > this.k {
		this.Pop()
	}

	return this.heap[0]
}

func (this *KthLargest) Push(val int) {
	this.heap = append(this.heap, val)
	node := len(this.heap) - 1
	for parent := (node - 1) / 2; node > 0 && parent >= 0; {
		if this.heap[node] > this.heap[parent] {
			break
		}
		this.heap[node], this.heap[parent] = this.heap[parent], this.heap[node]
		node = parent
		parent = (parent - 1) / 2
	}
}

func (this *KthLargest) Pop() {
	if len(this.heap) == 0 {
		return
	}
	heapSize := len(this.heap)
	this.heap[0], this.heap[heapSize-1] = this.heap[heapSize-1], this.heap[0]
	heapSize--

	node := 0
	for left, right := 2*node+1, 2*node+2; left < heapSize || right < heapSize; {
		mi := node
		if left < heapSize && this.heap[mi] > this.heap[left] {
			mi = left
		}
		if right < heapSize && this.heap[mi] > this.heap[right] {
			mi = right
		}
		if node == mi {
			break
		}
		this.heap[mi], this.heap[node] = this.heap[node], this.heap[mi]
		node = mi
		left = 2*node + 1
		right = 2*node + 2
	}
	this.heap = this.heap[:heapSize]
}
