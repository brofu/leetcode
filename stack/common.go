package stack

type ByteStack struct {
	bytes []byte
}

func NewByteStack() *ByteStack {
	return &ByteStack{
		bytes: make([]byte, 0),
	}
}

func (this *ByteStack) Pop() byte {
	o := this.bytes[len(this.bytes)-1]
	this.bytes = this.bytes[:len(this.bytes)-1]
	return o
}

func (this *ByteStack) Push(o byte) {
	this.bytes = append(this.bytes, o)
}

func (this *ByteStack) Size() int {
	return len(this.bytes)
}

type IntStack struct {
	nums []int
}

func NewIntStack() *IntStack {
	return &IntStack{
		nums: make([]int, 0),
	}
}

func (this *IntStack) Pop() int {
	o := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return o
}

func (this *IntStack) Push(o int) {
	this.nums = append(this.nums, o)
}

func (this *IntStack) Size() int {
	return len(this.nums)
}
