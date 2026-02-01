package stack

type FrontMiddleBackQueue struct {
}

func Constructor1670() FrontMiddleBackQueue {

	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {

}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {

}

func (this *FrontMiddleBackQueue) PushBack(val int) {

}

func (this *FrontMiddleBackQueue) PopFront() int {

	return 0
}

func (this *FrontMiddleBackQueue) PopMiddle() int {

	return 0
}

func (this *FrontMiddleBackQueue) PopBack() int {

	return 0
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
  * obj := Constructor();
   * obj.PushFront(val);
    * obj.PushMiddle(val);
	 * obj.PushBack(val);
	  * param_4 := obj.PopFront();
	   * param_5 := obj.PopMiddle();
	    * param_6 := obj.PopBack();
*/
