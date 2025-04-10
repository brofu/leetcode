package mst

import (
	"container/heap"
	"fmt"
	"sort"

	userHeap "github.com/brofu/leetcode/common/heap"
	uflib "github.com/brofu/leetcode/graph/union_find"
	"github.com/brofu/leetcode/tree/binaryheap"
)

func minimumCost(n int, connections [][]int) int {

	uf := uflib.NewPathCompactedUF(n + 1)

	sum := 0
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	for _, connection := range connections {
		if !uf.Connected(connection[0], connection[1]) {
			uf.Connect(connection[0], connection[1])
			sum += connection[2]
		}
	}

	if uf.Count() != 2 {
		return -1
	}
	return sum
}

/*

Prime Algorithm

1. How to `Cut`?
2. How to check the least weight? PQ

Key Points
1. The complexity of PQ

*/

func minimumCostPrim(n int, connections [][]int) int {

	sidesCount := 0
	// build graph
	graph := make([][][]int, n+1)
	for _, connection := range connections {
		from, to, weight := connection[0], connection[1], connection[2]
		graph[from] = append(graph[from], []int{to, weight})
		graph[to] = append(graph[to], []int{from, weight})
		sidesCount += 2
	}

	p := &prim{
		inMST: make([]bool, n+1),
		graph: graph,
		pq: binaryheap.NewPQBasedOnBH(sidesCount, func(data []interface{}, i, j int) bool {
			return data[i].([]int)[2] < data[j].([]int)[2]
		})}

	sum := 0
	// cut the first node
	p.inMST[1] = true
	p.cut(1)

	for p.pq.Size() > 0 {
		d := p.pq.Pop().([]int)
		to, weight := d[1], d[2]
		if p.inMST[to] {
			continue
		}
		p.inMST[to] = true
		p.cut(to)
		sum += weight
		fmt.Println("flag", d, sum)
	}

	for _, in := range p.inMST[1:] {
		if !in {
			return -1
		}
	}
	return sum
}

type prim struct {
	inMST []bool
	graph [][][]int
	pq    binaryheap.PriorityQueue
}

// cut
// get all the sides, and push the nodes which are not cut yet into the pq
func (p *prim) cut(n int) {
	for _, data := range p.graph[n] {
		to, weight := data[0], data[1]
		if p.inMST[to] {
			continue
		}
		in := []int{n, to, weight}
		p.pq.Push(in)
	}
}

/*
Version 2. With built-in heap
*/

func minimumCostPrimV2(n int, connections [][]int) int {

	sum := 0

	// build graph
	graph := make([][][]int, n+1)
	for _, connection := range connections {
		from, to, weight := connection[0], connection[1], connection[2]
		graph[from] = append(graph[from], []int{to, weight})
		graph[to] = append(graph[to], []int{from, weight})
	}

	pq := userHeap.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].([]int)[2] < data[j].([]int)[2]
	})
	inMST := make([]bool, n+1)

	var cut = func(n int) {
		for _, data := range graph[n] {
			to, weight := data[0], data[1]
			if inMST[to] {
				continue
			}
			in := []int{n, to, weight}
			heap.Push(pq, in)
		}
	}

	// cut the first node
	inMST[1] = true
	cut(1)

	for pq.Len() > 0 {

		d := heap.Pop(pq).([]int)
		to, weight := d[1], d[2]
		if inMST[to] {
			continue
		}
		inMST[to] = true
		cut(to)
		sum += weight
	}

	for _, in := range inMST[1:] {
		if !in {
			return -1
		}
	}

	return sum
}
