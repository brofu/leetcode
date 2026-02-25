package dijkstra

import (
	"container/heap"

	myHeap "github.com/brofu/leetcode/common/heap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	type probS struct {
		to   int
		prob float64
	}

	// build graph
	graph := make([][]probS, n)
	for index, record := range edges {
		graph[record[0]] = append(graph[record[0]], probS{record[1], succProb[index]})
		graph[record[1]] = append(graph[record[1]], probS{record[0], succProb[index]})
	}

	// initial the record
	probToNode := make([]float64, n)
	for index := range probToNode {
		probToNode[index] = 0
	}

	// initial the heap
	hp := myHeap.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].(probS).prob > data[j].(probS).prob
	})

	// start from start_node
	probToNode[start_node] = 1
	heap.Push(hp, probS{start_node, 1})

	for hp.Len() > 0 {
		current := heap.Pop(hp).(probS)
		currentNode, currentProb := current.to, current.prob

		if currentProb < probToNode[currentNode] {
			continue
		}

		if currentNode == end_node {
			return currentProb
		}

		for _, node := range graph[currentNode] { // check neighbors
			nextNode, neighborProb := node.to, node.prob
			nextProb := currentProb * neighborProb
			if nextProb > probToNode[nextNode] {
				probToNode[nextNode] = nextProb
				heap.Push(hp, probS{nextNode, nextProb})
			}
		}
	}
	return 0
}

type state1514 struct {
	node int
	prob float64
}
type heapState1514 []state1514

func (h *heapState1514) Len() int {
	return len(*h)
}

func (h *heapState1514) Less(x, y int) bool {
	return (*h)[x].prob > (*h)[y].prob
}

func (h *heapState1514) Swap(x, y int) {
	(*h)[x], (*h)[y] = (*h)[y], (*h)[x]
}

func (h *heapState1514) Push(x any) {
	*h = append(*h, x.(state1514))
}

func (h *heapState1514) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func maxProbabilityPV1(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	graph := make([][]state1514, n)
	for idx, edge := range edges {
		from := edge[0]
		graph[from] = append(graph[from], state1514{edge[1], succProb[idx]})
	}
	probTo := make([]float64, n)
	probTo[start_node] = 1

	var pq heapState1514 = make([]state1514, 0)

	heap.Push(&pq, state1514{start_node, 1})
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(state1514)
		if current.node == end_node {
			return current.prob
		}
		if probTo[current.node] > current.prob {
			continue
		}
		for _, next := range graph[current.node] {
			nextProb := current.prob * next.prob
			if probTo[next.node] >= nextProb {
				continue
			}
			probTo[next.node] = nextProb
			heap.Push(&pq, state1514{next.node, nextProb})
		}
	}

	return 0
}
