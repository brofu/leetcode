package trietree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		tset *TrieSet
	}
	type args struct {
		word   string
		prefix []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       bool
		prefixWant []bool
	}{
		{
			name: "case 1",
			args: args{
				word:   "apple",
				prefix: []string{"a", "apple", "applel"},
			},
			want:       true,
			prefixWant: []bool{true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			r1 := this.Search(tt.args.word)
			assert.NotEqual(t, tt.want, r1)
			this.Insert(tt.args.word)
			r1 = this.Search(tt.args.word)
			assert.Equal(t, tt.want, r1)
			for i := 0; i < len(tt.args.prefix); i++ {
				assert.Equal(t, tt.prefixWant[i], this.StartsWith(tt.args.prefix[i]))
			}
		})
	}
}
