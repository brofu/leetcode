package math

import (
	"container/heap"

	myHeap "github.com/brofu/leetcode/heap"
)

func nthSuperUglyNumber(n int, primes []int) int {

	ugly := make([]int, n+1)

	p := 1

	h := myHeap.NewHeap(len(primes))
	heap.Init(&h)

	pds := make([]int, len(primes))
	ps := make([]int, len(primes))
	for i := range pds {
		pds[i] = 1
		ps[i] = 1
		heap.Push(&h, pds[i])
	}

	for p <= n {
		m := heap.Pop(&h).(int)
		ugly[p] = m
		p++
		for m == h.Peek() {
			heap.Pop(&h)
		}
		for idx, pd := range pds {
			if m == pd {
				pds[idx] = primes[idx] * ugly[ps[idx]]
				ps[idx]++
				heap.Push(&h, pds[idx])
			}
		}
	}
	return ugly[n]
}
