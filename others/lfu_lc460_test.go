package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		fun     func()
	}{
		{
			name: "lc case 1",
			fun: func() {
				c := ConstructorLFU(2)

				c.Put(1, 1)
				c.Put(2, 2)
				val := c.Get(1)
				assert.Equal(t, 1, val)

				c.Put(3, 3)
				val = c.Get(2)
				assert.Equal(t, -1, val)

				val = c.Get(3)
				assert.Equal(t, 3, val)

				c.Put(4, 4)
				val = c.Get(1)
				assert.Equal(t, -1, val)

				val = c.Get(3)
				assert.Equal(t, 3, val)

				val = c.Get(4)
				assert.Equal(t, 4, val)

				val = c.Get(3)
				val = c.Get(4)

				val = c.Get(3)
				val = c.Get(4)

				val = c.Get(3)
				val = c.Get(4)

				c.Debug()
				c.Put(5, 5)
				c.Debug()
				c.Put(6, 6)
				c.Debug()
			},
		},
		{
			name:    "lc case 2",
			enabled: true,
			fun: func() {
				c := ConstructorLFU(2)
				c.Put(3, 1)
				c.Put(2, 1)
				c.Put(2, 2)
				c.Debug()
				c.Put(4, 4)
				c.Debug()
				val := c.Get(2)
				c.Debug()
				assert.Equal(t, 2, val)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enabled {
				tt.fun()
			}
		})
	}
}
