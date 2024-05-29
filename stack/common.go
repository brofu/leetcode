package stack

import "fmt"

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

func (this *IntStack) Debug() {
	fmt.Println(this.nums)
}

func (this *IntStack) DebugWthMessage(msg string) {
	fmt.Println(msg, this.nums)
}

type RuneStack struct {
	runes []rune
}

func NewRuneStack() *RuneStack {
	return &RuneStack{
		runes: make([]rune, 0),
	}
}

func (this *RuneStack) Pop() rune {
	o := this.runes[len(this.runes)-1]
	this.runes = this.runes[:len(this.runes)-1]
	return o
}

func (this *RuneStack) Push(o rune) {
	this.runes = append(this.runes, o)
}

func (this *RuneStack) Size() int {
	return len(this.runes)
}

func (this *RuneStack) DebugWthMessage(msg string) {
	fmt.Println(msg, this.runes)
}

type StringStack struct {
	strings []string
}

func NewStringStack() *StringStack {
	return &StringStack{
		strings: make([]string, 0),
	}
}

func (this *StringStack) Pop() string {
	o := this.strings[len(this.strings)-1]
	this.strings = this.strings[:len(this.strings)-1]
	return o
}

func (this *StringStack) Push(o string) {
	this.strings = append(this.strings, o)
}

func (this *StringStack) Size() int {
	return len(this.strings)
}

func (this *StringStack) DebugWthMessage(msg string) {
	fmt.Println(msg, this.strings)
}
