package array

type ProductOfNumbers struct {
	preM []int
}

func Constructor1352() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.preM = []int{1}
	}
	this.preM = append(this.preM, this.preM[len(this.preM)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.preM)
	if n-1 < k {
		return 0
	}
	return this.preM[n-1] / this.preM[n-1-k]
}
