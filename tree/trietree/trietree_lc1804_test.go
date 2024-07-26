package trietree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieV2_CountWordsStartingWith(t *testing.T) {
	type fields struct {
		tm *TrieMap
	}
	type args struct {
		prefix      string
		insertWords []string
		wordCount   string
	}
	tests := []struct {
		name          string
		enabled       bool
		fields        fields
		args          args
		want          int
		wordCountWant int
		f             func()
	}{
		{
			name: "lc case 1",
			f: func() {
				this := ConstructorV2()
				this.Insert("apple")
				this.Insert("apple")
				got := this.CountWordsEqualTo("apple")
				assert.Equal(t, 2, got)
				got = this.CountWordsStartingWith("app")
				t.Log("flag, got")
				assert.Equal(t, 2, got)
				this.Erase("apple")
				got = this.CountWordsEqualTo("apple")
				assert.Equal(t, 1, got)
				got = this.CountWordsStartingWith("app")
				assert.Equal(t, 1, got)
				this.Erase("apple")
				got = this.CountWordsEqualTo("apple")
				assert.Equal(t, 0, got)
				got = this.CountWordsStartingWith("apple")
				assert.Equal(t, 0, got)
			},
		},
		{
			name:    "lc case 2",
			enabled: true,
			f: func() {
				this := ConstructorV2()
				this.Insert("n")
				this.Insert("n")
				this.Insert("jvo")
				this.Erase("n")
				got := this.CountWordsStartingWith("n")
				assert.Equal(t, 1, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enabled {
				tt.f()
			}
		})
	}
}
