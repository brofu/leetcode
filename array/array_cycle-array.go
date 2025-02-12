package array

type CycleArrayInt struct {
	data  []int
	start int
	end   int
	count int
	size  int
}

func NewCycleArray(size int) CycleArrayInt {
	return CycleArrayInt{
		data:  make([]int, size),
		start: 0,
		end:   0,
		count: 0,
		size:  0,
	}
}

func (ca *CycleArrayInt) AddAtFirst(value int) {
	if ca.isFull() {
		ca.extend()
	}
	ca.start = ((ca.start - 1) + ca.size) % ca.size
	ca.data[ca.start] = value
	ca.count++
}

func (ca *CycleArrayInt) AddAtEnd(value int) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	ca.data[ca.end] = value
	ca.end++
	ca.count++
}

func (ca *CycleArrayInt) DeleteAtFirst(value int) {
	if ca.isEmpty() {
		return
	}
	ca.data[ca.start] = 0
	ca.start = ((ca.start + 1) + ca.size) % ca.size
	ca.count--
	if ca.count <= ca.size {
		ca.resize(ca.size / 2)
	}
}

func (ca *CycleArrayInt) DeleteAnEnd() {

	if ca.isEmpty() {
		return
	}

	ca.end = ((ca.end - 1) + ca.size) % ca.size
	ca.data[ca.end] = 0
	ca.count--

	if ca.count <= ca.size {
		ca.resize(ca.size / 2)
	}
}

func (ca *CycleArrayInt) isEmpty() bool {
	return ca.count == 0
}

func (ca *CycleArrayInt) isFull() bool {
	return ca.size == ca.count
}

func (ca *CycleArrayInt) extend() {
	data := make([]int, 2*ca.size)
	for idx, d := range ca.data {
		data[idx] = d
	}
	ca.data = data
}

func (ca *CycleArrayInt) resize(newSize int) {
	data := make([]int, newSize)
	for idx := 0; idx < ca.count; idx++ {
		data[idx] = ca.data[(ca.start+idx)%ca.count] //key point
	}
	ca.data = data
	ca.start = 0
	ca.end = ca.count
	ca.size = newSize
}
