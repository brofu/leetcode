package array

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j, p := m-1, n-1, len(nums1)-1

	for ; i >= 0 && j >= 0; p-- {
		if nums1[i] <= nums2[j] {
			nums1[p] = nums2[j]
			j--
		} else {
			nums1[p] = nums1[i]
			i--
		}
	}

	for ; i >= 0 && p >= 0; i, p = i-1, p-1 {
		nums1[p] = nums1[i]
	}

	for ; j >= 0 && p >= 0; j, p = j-1, p-1 {
		nums1[p] = nums2[j]
	}
}
