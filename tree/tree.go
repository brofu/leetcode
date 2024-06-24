package tree

import (
	"fmt"
	"sync"
)

var _ int = test(1)

func init() {
	test(2)
}

func test(n int) int {
	var once sync.Once
	once.Do(func() {})
	fmt.Println("I am here", n)
	return n
}
