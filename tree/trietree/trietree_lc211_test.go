package trietree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary_Search(t *testing.T) {
	type fields struct {
		ts *TrieSet
	}
	type args struct {
		searchWord []string
		words      []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []bool
	}{
		{
			name: "lc case 1",
			args: args{
				words: []string{"bad", "dad", "mad"},

				searchWord: []string{"pad", "bad", ".ad", "b.."},
			},
			want: []bool{false, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := ConstructorWordDictionary()
			for _, p := range tt.args.searchWord {
				got := this.Search(p)
				assert.Equal(t, false, got)
			}

			for _, w := range tt.args.words {
				this.AddWord(w)
			}

			for index, p := range tt.args.searchWord {
				got := this.Search(p)
				assert.Equal(t, tt.want[index], got)
			}
		})
	}
}
