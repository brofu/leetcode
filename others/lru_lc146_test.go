package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		{
			name: "lc case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lrc := Constructor(1)
			lrc.Put(2, 1)
			val := lrc.Get(1)
			assert.Equal(t, -1, val)

			lrc = Constructor(2)

			lrc.Put(1, 1)
			lrc.Put(2, 2)
			val = lrc.Get(1)
			assert.Equal(t, val, 1)

			lrc.Put(3, 3)
			val = lrc.Get(2)
			assert.Equal(t, -1, val)

			lrc.Put(4, 4)
			val = lrc.Get(1)
			assert.Equal(t, val, -1)

			val = lrc.Get(3)
			assert.Equal(t, val, 3)

			val = lrc.Get(4)
			assert.Equal(t, val, 4)
		})
	}
}
