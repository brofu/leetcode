package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructorV2(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCacheV2
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lrc := ConstructorV2(2)
			lrc.Put(1, 1)
			get := lrc.Get(1)
			assert.Equal(t, 1, get)
			lrc.Put(2, 2)
			lrc.Put(3, 3)
			get = lrc.Get(1)
			assert.Equal(t, -1, get)
			get = lrc.Get(2)
			assert.Equal(t, 2, get)
			get = lrc.Get(3)
			assert.Equal(t, 3, get)
		})
	}
}
