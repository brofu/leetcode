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

	heapifyIterative(nums, length)

	for index := length - 1; index >= 0; index -= 1 {
		if nums[0] > nums[index] { // make it stable
			nums[0], nums[index] = nums[index], nums[0]
		}
		heapifyIterative(nums, index)
	}

}

func heapifyIterative(nums []int, length int) {

	for deep := 0; ; deep += 1 {

		left, right := deep*2+1, deep*2+2

		if left < length && nums[0] < nums[left] {
			nums[0], nums[left] = nums[left], nums[0]
		}
		if right < length && nums[0] < nums[right] {
			nums[0], nums[right] = nums[right], nums[0]
		}

		if left >= length || right >= length {
			break
		}
	}
}
