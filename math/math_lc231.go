package math

func singleNumber(nums []int) int {

	res := 0
	for _, x := range nums {
		res = res ^ x
	}

	return res
}
