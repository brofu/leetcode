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