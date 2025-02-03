package array

type MyLinkedListV2 struct {
	data []int
}

func Constructor2() MyLinkedListV2 {
	return MyLinkedListV2{
		data: make([]int, 0, 4),
	}
}

func (this *MyLinkedListV2) Get(index int) int {
	if index < 0 || index > len(this.data)-1 {
		return -1
	}
	return this.data[index]
}

func (this *MyLinkedListV2) AddAtHead(val int) {
	this.data = append([]int{val}, this.data...)
}

func (this *MyLinkedListV2) AddAtTail(val int) {
	this.data = append(this.data, val)
}

func (this *MyLinkedListV2) AddAtIndex(index int, val int) {
	if index < 0 || index > len(this.data) {
		return
	}
	if index == len(this.data) {
		this.data = append(this.data, val)
		return
	}
	this.data = append(this.data[:index], append([]int{val}, this.data[index:]...)...)
}

func (this *MyLinkedListV2) DeleteAtIndex(index int) {
	if index < 0 || index >= len(this.data) {
		return
	}

	this.data = append(this.data[:index], this.data[index+1:]...)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/**
Version 2

Maintain the slice info at application level, including
	*	size
	*	extend the capability
	*	move the elements

*/
type MyLinkedList struct {
	data []int
	size int
}

func ConstructorV3() MyLinkedList {
	return MyLinkedList{
		data: make([]int, 4, 4),
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	return this.data[index]
}

func (this *MyLinkedList) extend() {
	data := this.data
	this.data = make([]int, 2*len(data), 2*cap(data))
	for idx, v := range data {
		this.data[idx] = v
	}
}
func (this *MyLinkedList) makeSureCap() {
	if this.size == cap(this.data) {
		this.extend()
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.makeSureCap()
	for idx := this.size; idx > 0; idx-- {
		this.data[idx] = this.data[idx-1]
	}
	this.size++
	this.data[0] = val
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.makeSureCap()
	this.data[this.size] = val
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	this.makeSureCap()
	for idx := this.size; idx > index; idx-- {
		this.data[idx] = this.data[idx-1]
	}
	this.data[index] = val
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	for idx := index; idx < this.size-1; idx++ {
		this.data[idx] = this.data[idx+1]
	}
	this.size--
}
