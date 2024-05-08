package array

import (
	"math"
	"math/rand"

	"github.com/brofu/leetcode/sort"
)

/**
KP
	1. QuickSelect.
	2. T O(n) on average, M O(1)
	3. https://stackoverflow.com/questions/56940793/quickselect-time-complexity-explained
*/

func findKthLargest(nums []int, k int) int {

	low, high := 0, len(nums)-1
	//pivot := (low + high) / 2
	//rand.Seed(time.Now().UnixNano())
	pivot := low + rand.Intn(high-low+1)

	for {
		pivot = quickSelectHelper(nums, pivot, low, high)
		if pivot == len(nums)-k {
			break
		}

		if pivot > len(nums)-k { // the result locates on the left of pivot
			high = pivot - 1
		} else { // the result locates on the right of the pivot
			low = pivot + 1
		}

		//pivot = (low + high) / 2
		pivot = low + rand.Intn(high-low+1)

	}

	return nums[pivot]
}

/**
KP.
	1. How to make sure all the smaller ones on the left and the bigger ones on the right?
*/
func quickSelectHelper(array []int, pivot, low, high int) int {

	val := array[pivot]
	array[high], array[pivot] = array[pivot], array[high]
	c := low

	for i := low; i < high; i += 1 {
		if array[i] < val {
			array[i], array[c] = array[c], array[i]
			c += 1
		}
	}

	array[high], array[c] = array[c], array[high]

	return c
}

/**
KP
	1.	Utilize array index for sorting, and value as the occurred times
	2.	T O(N).
	3.	M O(?) depends on the value range of `max - min`
*/
func findKthLargestV2(nums []int, k int) int {

	min, max := math.MaxInt, math.MinInt

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	count := make([]int, max-min+1)

	for _, num := range nums {
		count[num-min] += 1
	}

	for index := len(count) - 1; index >= 0; index -= 1 {
		k -= count[index]
		if k <= 0 {
			return index + min
		}
	}

	return 0
}

/**
KP
	1. QuickSort
	2. T O(N*lgN)
	3. M O(1)
*/
func findKthLargestV3(nums []int, k int) int {

	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, low, high int) {

	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := low + rand.Intn(high-low+1)
	val := nums[pivot]
	coursor := low

	nums[pivot], nums[high] = nums[high], nums[pivot]

	for i := low; i < high; i += 1 {
		if nums[i] < val {
			nums[i], nums[coursor] = nums[coursor], nums[i]
			coursor += 1
		}
	}
	nums[coursor], nums[high] = nums[high], nums[coursor]

	return coursor
}

/**
KP
	1. Heap Sort
*/

func findKthLargestV4(nums []int, k int) int {
	sort.HeapSort(nums)
	return nums[len(nums)-k]
}
