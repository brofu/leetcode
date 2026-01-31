package array

/*
KP.
	1. Cycle Array.
	2. Control `start` and `end`, with % method
	2. Use `campability` and `size` control the `full` and `empty`
*/

type MyCircularDeque struct {
	data            []int
	start, end      int
	size, capbility int
}

func Constructor641(k int) MyCircularDeque {
	d := make([]int, k)
	return MyCircularDeque{
		data:      d,
		start:     0,
		end:       0,
		size:      0,
		capbility: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.start = (this.start - 1 + this.capbility) % this.capbility
	this.data[this.start] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.end] = value
	this.end = (this.end + 1) % this.capbility
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % this.capbility
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.end = (this.end - 1 + this.capbility) % this.capbility
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.start]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.end-1+this.capbility)%this.capbility]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capbility
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
