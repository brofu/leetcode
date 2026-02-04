package common

type BitSet []uint64

func NewBitSet(n int) BitSet {
	return make([]uint64, (n+63)>>6)
}

// Set set the nth bit
func (b BitSet) Set(n int) {
	w := n / 64
	offset := uint(n % 64)
	mask := uint64(1) << offset
	b[w] |= mask
}

func (b BitSet) Get(n int) bool {
	w := n / 64
	offset := uint(n % 64)
	mask := uint64(1) << offset
	return (b[w] & mask) != 0
}
