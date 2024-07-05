package tree

import (
	"sync"
)

var _ int = test(1)

func init() {
	test(2)
}

func test(n int) int {
	var once sync.Once
	once.Do(func() {})
	return n
}
