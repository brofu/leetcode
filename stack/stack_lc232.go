package stack

type MyQueue struct {
	s1, s2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	num := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return num
}

/*
KP:
	1. How to combine a queue with 2 stacks?
*/
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.s2) == 0 {
		for size := len(this.s1); size > 0; size-- {
			num := this.s1[size-1]
			this.s2 = append(this.s2, num)
			this.s1 = this.s1[:size-1]
		}
	}
	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}
