package sort

func HeapSort(nums []int) {

	length := len(nums)

	// construct the max heap
	for index := length/2 - 1; index >= 0; index -= 1 {
		heapify(nums, length, index)
	}

	// extract from the max heap
	for index := length - 1; index >= 0; index -= 1 {
		if nums[0] > nums[index] { // make it stable
			nums[0], nums[index] = nums[index], nums[0]
		}
		heapify(nums, index, 0)
	}
}

func heapify(nums []int, length, index int) {
	largest := index

	left := index*2 + 1
	right := index*2 + 2

	if left < length && nums[left] > nums[largest] {
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}

	if largest != index {
		nums[index], nums[largest] = nums[largest], nums[index]
		heapify(nums, length, largest)
	}
}

/*
KP:
	1. 2 steps similar as Priority Queue. Take `increasing sort` as an example
		a.	Construct maxHeap, by `swim` operation. (swich all the numbers one by one)
		b.	Get the `increasing sort`, by `sink` operation. (swap nums[0] with nums[i-1], and sink nums[0])
*/

func HeapSortIterativeV2(nums []int) {

	// construct the heap. swim the max heap
	for i := 0; i < len(nums); i++ {
		heapMaxSwim(nums, i)
	}

	// sink the max heap
	heapSize := len(nums)
	for heapSize > 0 {
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		heapMaxSink(nums, 0, heapSize)
	}
}

func heapMaxSwim(nums []int, node int) {
	for parent := (node - 1) / 2; parent >= 0 && nums[node] > nums[parent]; {
		nums[node], nums[parent] = nums[parent], nums[node]
		node = parent
		parent = (parent - 1) / 2
	}
}
func heapMaxSink(nums []int, node, heapSize int) {
	for 2*node+1 < heapSize || 2*node+2 < heapSize {
		maxM := node
		if 2*node+1 < heapSize && nums[maxM] < nums[2*node+1] {
			maxM = 2*node + 1
		}
		if 2*node+2 < heapSize && nums[maxM] < nums[2*node+2] {
			maxM = 2*node + 2
		}
		if maxM == node {
			break
		}
		nums[node], nums[maxM] = nums[maxM], nums[node]
		node = maxM
	}
}

/*
KP:
	1. 对于一个二叉堆来说，其左右子堆（子树）也是一个二叉堆。
	2. 如果有两个二叉堆，和一个二叉堆节点，那么可以把这个节点作为堆顶（根节点），两个二叉堆作为左右子堆（子树），构建出一棵新的二叉堆（二叉树）
	   但是有个问题，构建出的这个新二叉堆，它的左右子堆肯定都符合堆的性质，但这个新的根节点的值可能不符合堆的性质，怎么办？
		sink 方法不就是专门针对这种情况的吗？只要对根节点调用一次 sink 方法，就能让这个新的二叉堆符合堆的性质了。
	3. 所以，一个优化的建堆方案
		a. 从最后一个非叶子节点开始，向前，逐个调用sink操作
		b. 叶子节点自身，肯定是二叉堆
		c. 保证每一个非叶子节点，和它下面的二叉堆，新组成的，也是一个二叉堆
*/
func HeapSortIterativeV3(nums []int) {

	// construct the heap. swim the max heap
	heapSize := len(nums)
	for i := (heapSize - 1) / 2; i >= 0; i-- {
		heapMaxSink(nums, i, heapSize)
	}

	// sink the max heap
	for heapSize > 0 {
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		heapMaxSink(nums, 0, heapSize)
	}
}

func HeapSortIterative(nums []int) {
	length := len(nums)

	// setup max heap
	for index := length - 1; index >= 0; index -= 1 {
		// swim
		for parent := (index - 1) / 2; ; parent = (parent - 1) / 2 {
			if nums[index] > nums[parent] {
				nums[index], nums[parent] = nums[parent], nums[index]
			}
			if parent == 0 {
				break
			}
		}
	}

	for index := length - 1; index >= 0; index -= 1 {
		if nums[0] <= nums[index] { // make it stable
			continue
		}
		nums[0], nums[index] = nums[index], nums[0]
		// sink
		root := 0
		for left, right := 1, 2; left < index; left, right = root*2+1, root*2+2 {
			maxIndex := left
			if right < index && nums[maxIndex] < nums[right] {
				maxIndex = right
			}
			if nums[root] >= nums[maxIndex] {
				break
			}
			nums[root], nums[maxIndex] = nums[maxIndex], nums[root]
			root = maxIndex
		}
	}
}

func HeapSortPV1(nums []int) {

	length := len(nums)

	// construct max heap
	// why should loop from 1/2 largest indext to 0?
	for i := (length - 1) / 2; i >= 0; i -= 1 {
		heapifyPV1(nums, length, i)
	}

	// extract from max heap
	for i := length - 1; i >= 0; i -= 1 {
		if nums[0] > nums[i] {
			nums[0], nums[i] = nums[i], nums[0]
		}
		heapifyPV1(nums, i, 0)
	}
}

func heapifyPV1(nums []int, length, index int) {
	largest := index

	left := 2*index + 1
	right := 2*index + 2

	if left < length && nums[left] > nums[largest] {
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}

	if largest != index {
		nums[largest], nums[index] = nums[index], nums[largest]
		heapify(nums, length, largest)
	}
}

func HeapSortPV2(nums []int) {

	for i := (len(nums) - 1) / 2; i >= 0; i-- { // construct max heap, sink from index (len(nums)-1)/2
		heapifyPV2(nums, len(nums), i)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i] // swap the largest to nums[i]
		heapifyPV2(nums, i, 0)
	}

}

func heapifyPV2(nums []int, length, index int) {

	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < length && nums[left] > nums[largest] {
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}

	if largest != index {
		nums[largest], nums[index] = nums[index], nums[largest]
		heapifyPV2(nums, length, largest)
	}
}

func MergeSort(nums []int) {

	temp := make([]int, len(nums))

	var mergeSort func([]int, int, int)
	mergeSort = func(nums []int, low, high int) {
		if low == high {
			return
		}
		mid := (low + high) / 2

		// merge left part
		mergeSort(nums, low, mid)
		// merge right part
		mergeSort(nums, mid+1, high)

		// merge the whole part
		for index := low; index <= high; index++ {
			temp[index] = nums[index]
		}

		index, i, j := low, low, mid+1
		for ; i <= mid && j <= high; index += 1 {
			if temp[i] <= temp[j] {
				nums[index] = temp[i]
				i += 1
			} else {
				nums[index] = temp[j]
				j += 1
			}
		}
		for ; i <= mid; index, i = index+1, i+1 {
			nums[index] = temp[i]
		}
		for ; j <= high; index, j = index+1, j+1 {
			nums[index] = temp[j]
		}

	}

	mergeSort(nums, 0, len(nums)-1)
}
