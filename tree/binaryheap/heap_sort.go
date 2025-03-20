package binaryheap

func HeapSort(nums []int) {

	// heapify
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyCursively(nums, n, i)
	}

	// swap and restore the heap property
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		// sink from the `head` location of array. Similar to the `Pop` operation in PQ
		heapifyCursively(nums, i, 0)
	}
}

/*
Iterative version of heapify

`Sink` operation for each sub-tree
*/
func heapify(nums []int, size, index int) {

	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	for left < size && right < size {
		if nums[smallest] > nums[left] {
			smallest = left
		}
		if nums[smallest] > nums[right] {
			smallest = right
		}
		if smallest == index {
			break
		}
		nums[smallest], nums[index] = nums[index], nums[smallest]

		index = smallest
		left = index*2 + 1
		right = index*2 + 2
	}

	if left < size && nums[left] < nums[index] {
		nums[index], nums[left] = nums[left], nums[index]
	}
}

/*
Cursive version heapify
*/
func heapifyCursively(nums []int, size, index int) {

	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left < size && nums[left] < nums[smallest] {
		smallest = left
	}
	if right < size && nums[right] < nums[smallest] {
		smallest = right
	}

	if smallest != index {
		nums[smallest], nums[index] = nums[index], nums[smallest]
		heapifyCursively(nums, size, smallest)
	}
}
