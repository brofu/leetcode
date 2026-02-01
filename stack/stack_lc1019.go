package stack

func nextLargerNodes(head *ListNode) []int {

	nums := make([]int, 0)
	st := make([]int, 0)

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = 0
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, nums[i])
	}

	return res
}
