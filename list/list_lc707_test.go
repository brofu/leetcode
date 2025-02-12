package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList(t *testing.T) {

	t.Run("defalut", func(t *testing.T) {
		l := Constructor()
		l.AddAtHead(1)
		l.AddAtTail(3)
		val := l.Get(1)
		assert.Equal(t, 3, val)
		l.AddAtIndex(1, 2)
		val = l.Get(1)
		assert.Equal(t, 2, val)
	})
}
