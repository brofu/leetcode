package stack

/*
KP:
	The idea of monotonic stack

TC:
	O(N): N = len(nums1) + len(nums2)
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	m := make(map[int]int)
	for idx, num := range nums2 {
		m[num] = idx
	}
	temp := make([]int, len(nums2))
	st := []int{}

	for i := len(nums2) - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] <= nums2[i] { //peak
			st = st[:len(st)-1] //pop
		}
		if len(st) == 0 {
			temp[i] = -1
		} else {
			temp[i] = st[len(st)-1] // peak
		}
		st = append(st, nums2[i]) // push
	}

	res := make([]int, len(nums1))
	for idx, num := range nums1 {
		res[idx] = temp[m[num]]
	}

	return res
}

/*
KP:
	Better SC
*/
func nextGreaterElementV2(nums1 []int, nums2 []int) []int {

	m := make(map[int]int)
	st := []int{}

	for i := len(nums2) - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] <= nums2[i] { //peak
			st = st[:len(st)-1] //pop
		}
		if len(st) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = st[len(st)-1] // peak
		}
		st = append(st, nums2[i]) // push
	}

	res := make([]int, len(nums1))
	for idx, num := range nums1 {
		res[idx] = m[num]
	}

	return res
}
