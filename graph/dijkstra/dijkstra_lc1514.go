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
