package array

// pivotArray
// Problem 2161
// O(N), O(N)
// Key Points
// 1. extra list to store the info
//
func pivotArray(nums []int, pivot int) []int {

	smaller, larger := []int{}, []int{}

	equalNumber := 0

	for _, val := range nums {

		if val < pivot {
			smaller = append(smaller, val)
		}
		if val == pivot {
			equalNumber += 1
		}
		if val > pivot {
			larger = append(larger, val)
		}
	}

	for index, val := range smaller {
		nums[index] = val
	}
	offset := len(smaller)

	for n := 0; n < equalNumber; n++ {
		nums[offset+n] = pivot
	}

	offset += equalNumber

	for index, val := range larger {
		nums[offset+index] = val
	}

	return nums
}

// pivotArray
// Problem 2161
// Key Points
// reduce time
// T: 83.8 M: 75
// 1. reduce the time of slice extending capbality?
func pivotArrayV2(nums []int, pivot int) []int {

	total := len(nums)
	rs := make([]int, total)

	smaller, equal, larger := 0, 0, 0

	for _, val := range nums {
		if val < pivot {
			smaller += 1
		} else if val == pivot {
			equal += 1
		}
	}

	// adjust the index
	larger = smaller + equal
	equal = smaller
	smaller = 0

	for _, val := range nums {

		if val < pivot {
			rs[smaller] = val
			smaller += 1
		} else if val > pivot {
			rs[larger] = val
			larger += 1
		} else {
			rs[equal] = pivot
			equal += 1
		}
	}

	return rs
}

// pivotArrayV3
// Problem 2161
// T O(N^2) S O(1)
// Key point. Insertion Sort  idea, to reduce memory usage
// A sub case of Insertion Sort
// Question, can use other sort algo?
func pivotArrayV3(nums []int, pivot int) []int {

	for i, val := range nums {
		if val <= pivot {
			j := i
			for ; j > 0 && (val == pivot && nums[j-1] > pivot || val < pivot && nums[j-1] >= pivot); j -= 1 {
				nums[j] = nums[j-1]
			}
			nums[j] = val
		}
	}

	return nums
}
