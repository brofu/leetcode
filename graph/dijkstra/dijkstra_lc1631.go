package dijkstra

import (
	"container/heap"
	"math"

	"github.com/brofu/leetcode/common"
	myHeap "github.com/brofu/leetcode/common/heap"
)

func minimumEffortPath(heights [][]int) int {

	m, n := len(heights), len(heights[0])

	effortToNode := make([][]int, m)
	for i := 0; i < m; i++ {
		effortToNode[i] = make([]int, n)
		for j := 0; j < n; j++ {
			effortToNode[i][j] = math.MaxInt
		}
	}

	effortToNode[0][0] = 0

	hp := myHeap.NewInterfacePriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].([]int)[2] < data[j].([]int)[2]
	})

	heap.Push(hp, []int{0, 0, 0})

	for hp.Len() > 0 {
		current := heap.Pop(hp).([]int)

		currentX, currentY, currentDist := current[0], current[1], current[2]

		if currentDist > effortToNode[currentX][currentY] { // exists a smaller distance
			continue
		}

		effortToNode[currentX][currentY] = currentDist

		if currentX == m-1 && currentY == n-1 { // already arrived at the bottom-right node
			continue
		}

		// check neighbors
		for _, neighbor := range [][]int{{currentX - 1, currentY}, {currentX, currentY - 1}, {currentX + 1, currentY}, {currentX, currentY + 1}} {
			neighborX, neighborY := neighbor[0], neighbor[1]
			if neighborX < 0 || neighborX == m || neighborY < 0 || neighborY == n { //invalid neighbors
				continue
			}
			// KEY POINT
			// The definition and update logic of `effortToNode`
			effortToNeighbor := common.MaxInt(effortToNode[currentX][currentY], common.AbsIntSub(heights[currentX][currentY], heights[neighborX][neighborY]))
			if effortToNeighbor < effortToNode[neighborX][neighborY] {
				effortToNode[neighborX][neighborY] = effortToNeighbor
				heap.Push(hp, []int{neighborX, neighborY, effortToNeighbor})
			}
		}
	}

	return effortToNode[m-1][n-1]
}

type DKState1631 struct {
	X, Y   int
	Effort int
}

type HeapDKState1631 []DKState1631

func (this *HeapDKState1631) Push(x any) {
	*this = append(*this, x.(DKState1631))
}

func (this *HeapDKState1631) Pop() any {
	x := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return x
}

func (this *HeapDKState1631) Len() int {
	return len(*this)
}

func (this *HeapDKState1631) Less(x, y int) bool {
	return (*this)[x].Effort < (*this)[y].Effort
}

func (this *HeapDKState1631) Swap(x, y int) {
	(*this)[x], (*this)[y] = (*this)[y], (*this)[x]
}

type Neighboor struct {
	x, y int
}

func minimumEffortPathPV1(heights [][]int) int {

	m, n := len(heights), len(heights[0])

	effortTo := make([][]int, m)
	for idx := range effortTo {
		effortTo[idx] = make([]int, n)
		for j := range effortTo[idx] {
			effortTo[idx][j] = 1000000
		}
	}
	effortTo[0][0] = 0

	var pq HeapDKState1631 = make([]DKState1631, 0)
	heap.Push(&pq, DKState1631{
		X:      0,
		Y:      0,
		Effort: 0,
	})

	for pq.Len() > 0 {
		current, _ := heap.Pop(&pq).(DKState1631)

		if current.X == m-1 && current.Y == n-1 {
			return current.Effort
		}

		if effortTo[current.X][current.Y] < current.Effort {
			continue
		}

		for _, neighbor := range getNeighboor(m, n, current.X, current.Y) {
			nextEffort := common.MaxInt(current.Effort, common.AbsIntSub(heights[current.X][current.Y], heights[neighbor.x][neighbor.y]))
			if effortTo[neighbor.x][neighbor.y] <= nextEffort {
				continue
			}
			effortTo[neighbor.x][neighbor.y] = nextEffort
			heap.Push(&pq, DKState1631{neighbor.x, neighbor.y, nextEffort})
		}
	}
	return effortTo[m-1][n-1]
}

func getNeighboor(m, n, x, y int) (neighboors []Neighboor) {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, direction := range directions {
		nx := x + direction[0]
		ny := y + direction[1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		neighboors = append(neighboors, Neighboor{nx, ny})
	}
	return
}
