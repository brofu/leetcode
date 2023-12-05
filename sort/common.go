package sort

import "fmt"

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

//func HeapSortRecursive(nums []int) {
//	fmt.Println(nums)
//	heapifyRecursive(nums, 0, 1, 2)
//}
//
//func heapifyRecursive(nums []int, root, left, right int) {
//	fmt.Println(root, left, right)
//	if left >= len(nums) || right >= len(nums) {
//		return
//	}
//
//	fmt.Println(root, left, right, nums[root], nums[left], nums[right])
//	heapifyRecursive(nums, left, left*2+1, left*2+2)
//	heapifyRecursive(nums, right, right*2+1, right*2+2)
//
//	if nums[root] > nums[left] {
//		nums[root], nums[left] = nums[left], nums[root]
//	}
//
//	if nums[root] > nums[right] {
//		nums[root], nums[right] = nums[right], nums[root]
//	}
//}

func HeapSortIterative(nums []int) {
	length := len(nums)

	fmt.Println(nums)
	for index := 0; index <= length/2; index += 1 {
		heapifyIterative(nums, length, index)
	}

	for index := length - 1; index > 0; index -= 1 {
		if nums[0] > nums[index] { // make it stable
			nums[0], nums[index] = nums[index], nums[0]
		}
		heapifyIterative(nums, index, 0)
	}

}

func heapifyIterative(nums []int, length, index int) {

	for left := 2*index + 1; left < length; left = left*2 + 1 {
		fmt.Println("left", nums, index, left, nums[index], nums[left])
		if nums[index] < nums[left] {
			fmt.Println("left hit")
			nums[index], nums[left] = nums[left], nums[index]
		}
		fmt.Println("left after", nums)
	}

	for right := 2*index + 2; right < length; right = right*2 + 2 {
		fmt.Println("right", nums, index, right, nums[index], nums[right])
		if nums[index] < nums[right] {
			nums[index], nums[right] = nums[right], nums[index]
			fmt.Println("right hit")
		}
		fmt.Println("right after", nums)
	}

	fmt.Println(nums)
}
