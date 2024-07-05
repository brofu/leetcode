package graph

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/brofu/leetcode/common"
)

func minimumEffortPath(heights [][]int) int {

	m, n := len(heights), len(heights[0])

	distanceTo := make([][]int, m)
	for i := 0; i < m; i++ {
		distanceTo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distanceTo[i][j] = math.MaxInt
		}
	}
	distanceTo[0][0] = 0
	fmt.Println("flag", distanceTo)
	pq := make(stateXYPriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &stateXY{0, 0, 0})

	for pq.Len() > 0 {
		currendNode := pq.Pop().(*stateXY)

		currentX, currentY := currendNode.x, currendNode.y

		if currentX == m-1 && currentY == n-1 {
			return currendNode.distanceFromStart
		}

		if currendNode.distanceFromStart > distanceTo[currentX][currentY] { // there is already a better path
			continue
		}

		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} { // possible neighbors of currentNode
			nextX := currentX + dir[0]
			nextY := currentY + dir[1]
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
				continue
			}
			nextDistance := common.MaxInt(distanceTo[currentX][currentY], common.AbsIntSub(heights[currentX][currentY], heights[nextX][nextY]))

			if nextDistance < distanceTo[nextX][nextY] {
				distanceTo[nextX][nextY] = nextDistance
				heap.Push(&pq, &stateXY{nextX, nextY, nextDistance})
			}
		}
	}

	return -1

}
