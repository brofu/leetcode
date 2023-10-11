package sort

// InsertionSort
// T O(n) S O(1)
// Skill: reduce the set value

func InsertionSort(nums []int) {
	for i, val := range nums {

		j := i - 1

		for ; j >= 0 && nums[j] > val; j -= 1 {
			nums[j+1] = nums[j]
		}

		nums[j+1] = val

	}
}
