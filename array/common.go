package array

type DifferArray struct {
	data []int
}

func (d *DifferArray) Increate(i, j int, val int) {
	d.data[i] += val
	if j+1 < len(d.data) {
		d.data[j+1] -= val
	}
}

func (d *DifferArray) Result() []int {
	res := make([]int, len(d.data))
	res[0] = d.data[0]
	for idx := 1; idx < len(res); idx++ {
		res[idx] = res[idx-1] + d.data[idx]
	}
	return res
}

func NewDifferArray(input []int) *DifferArray {

	d := make([]int, len(input))
	d[0] = input[0]
	for i := 1; i < len(input); i++ {
		d[i] = input[i] - input[i-1]
	}
	return &DifferArray{
		data: d,
	}
}
