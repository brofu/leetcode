package array

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	for index, val := range nums {
		sum[index+1] = val + sum[index]
	}
	return NumArray{preSum: sum}
}

// KP. the left number is included in the range
func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
