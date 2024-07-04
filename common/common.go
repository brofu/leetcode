package common

import "fmt"

//Comparable is the interface which is comparable
type Comparable interface {
	// if a > b, a.CompareTo(b) return 1
	// if a == b, a.CompareTo(b) return 0
	// if a < b, a.CompareTo(b) return -1
	CompareTo(Comparable) int
}

// Merge Sort

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	middle := len(nums) / 2

	MergeSort(nums[:middle])
	MergeSort(nums[middle:])

	left := make([]int, middle)
	right := make([]int, len(nums)-middle)

	_ = copy(left, nums[:middle])
	_ = copy(right, nums[middle:])

	index, i, j := 0, 0, 0
	for ; i < middle && j < len(nums)-middle; index += 1 {
		if left[i] <= right[j] {
			nums[index] = left[i]
			i += 1
		} else {
			nums[index] = right[j]
			j += 1
		}
	}

	for ; i < middle; index, i = index+1, i+1 {
		nums[index] = left[i]
	}
	for ; j < len(nums)-middle; index, j = index+1, j+1 {
		nums[index] = right[j]
	}
}

func MergeSortPV1(nums []int) {
	if len(nums) <= 1 {
		return
	}

	middle := len(nums) / 2

	MergeSortPV1(nums[:middle])
	MergeSortPV1(nums[middle:])

	left := make([]int, len(nums[:middle]))
	right := make([]int, len(nums[middle:]))

	_ = copy(left, nums[:middle])
	_ = copy(right, nums[middle:])

	i, j := 0, 0
	index := 0

	for ; i < middle && j < len(nums[middle:]); index += 1 {
		if left[i] < right[j] {
			nums[index] = left[i]
			i += 1
		} else {
			nums[index] = right[j]
			j += 1
		}
	}

	for ; i < middle; i, index = i+1, index+1 {
		nums[index] = left[i]
	}

	for ; j < len(nums[middle:]); j, index = j+1, index+1 {
		nums[index] = right[j]
	}
}

func Debug(debug func(), msg string, args ...interface{}) {
	fmt.Println(msg)
	if debug != nil {
		debug()
	}
	fmt.Println(msg, args)
}

func DeepEqualIntSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, list := range a {
		if !CheckIntSlice(list, b) {
			return false
		}
	}
	return true
}

func CheckIntSlice(a []int, b [][]int) bool {
	result := false

OutLoop:
	for _, list := range b {

		if len(list) != len(a) {
			continue
		}
		mapA := make(map[int]int)
		for _, e := range a {
			if _, exists := mapA[e]; exists {
				mapA[e] += 1
			} else {
				mapA[e] = 1
			}
		}
		mapB := make(map[int]int)
		for _, e := range list {
			if _, exists := mapB[e]; exists {
				mapB[e] += 1
			} else {
				mapB[e] = 1
			}
		}

		for keyA, valA := range mapA {
			if valB, exists := mapB[keyA]; !exists {
				continue OutLoop
			} else {
				if valA != valB {
					continue OutLoop
				}
			}
		}

		result = true
	}

	return result
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsIntSub(a, b int) int {
	sub := a - b
	if sub >= 0 {
		return sub
	}
	return -sub
}
