package stack

import "github.com/brofu/leetcode/common"

/*
KP:
	1. Get the idea of `record the min value when each new is pushed into the stack`

*/
type MinStack struct {
	st    []int
	minSt []int
}

func Constructor155() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.st) == 0 {
		this.st = append(this.st, val)
		this.minSt = append(this.minSt, val)
		return
	}
	this.st = append(this.st, val)
	minV := common.MinInt(val, this.minSt[len(this.minSt)-1])
	this.minSt = append(this.minSt, minV)
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
  * obj := Constructor();
   * obj.Push(val);
    * obj.Pop();
	 * param_3 := obj.Top();
	  * param_4 := obj.GetMin();
*/
