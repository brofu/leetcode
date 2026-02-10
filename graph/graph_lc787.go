package graph

import (
	"container/heap"
	"math"
)

type state787 struct {
	current        int
	distFromStart  int
	stepsFromStart int
}

type hp []state787

func (this *hp) Len() int {
	return len(*this)
}
func (this *hp) Push(x any) {
	(*this) = append(*this, x.(state787))
}

func (this *hp) Pop() any {
	val := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return val
}

func (this *hp) Less(i, j int) bool {
	return (*this)[i].distFromStart < (*this)[j].distFromStart
}

func (this *hp) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

/*
KP:
	1. dijkstra with constraints
	2. What would happen, if we set graph := make([][][]int, n)?
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	type edge struct {
		to    int
		price int
	}
	graph := make([][]edge, n)
	for _, flight := range flights {
		from := flight[0]
		graph[from] = append(graph[from], edge{to: flight[1], price: flight[2]})
	}

	// dist[i][j] means, the shorted path, from start, to node i, with j steps
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, k+2)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	for j := 0; j < len(dist[src]); j++ {
		dist[src][j] = 0
	}

	h := &hp{}
	heap.Init(h)

	heap.Push(h, state787{
		current:        src,
		distFromStart:  0,
		stepsFromStart: 0,
	})
	for h.Len() > 0 {
		current := heap.Pop(h).(state787)

		if current.current == dst {
			if current.stepsFromStart <= k+1 {
				return current.distFromStart
			}
			return -1
		}

		if dist[current.current][current.stepsFromStart] < current.distFromStart {
			continue
		}

		for _, next := range graph[current.current] {
			nextNode := next.to
			price := next.price
			nextDist := current.distFromStart + price
			nextStep := current.stepsFromStart + 1
			if nextStep > k+1 || dist[nextNode][nextStep] <= nextDist {
				continue
			}
			dist[nextNode][nextStep] = nextDist
			heap.Push(h, state787{
				current:        nextNode,
				distFromStart:  nextDist,
				stepsFromStart: nextStep,
			})
		}
	}

	return -1
}
