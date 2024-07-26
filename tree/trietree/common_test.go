package trietree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieMap_Put(t *testing.T) {
	type fields struct {
		size int
		root *TrieTreeNode
	}
	type args struct {
		key string
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TrieMap{
				size: tt.fields.size,
				root: tt.fields.root,
			}
			tm.Put(tt.args.key, tt.args.val)
		})
	}
}

func TestTrieMap_putNode(t *testing.T) {
	type fields struct {
		size int
		root *TrieTreeNode
	}
	type args struct {
		root  *TrieTreeNode
		key   string
		val   interface{}
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TrieTreeNode
	}{
		{
			name: "case 1",
			fields: fields{
				size: 0,
				root: &TrieTreeNode{
					Children: [26]*TrieTreeNode{},
				},
			},
			args: args{
				key: "the",
				val: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TrieMap{
				tt.fields.size,
				tt.fields.root,
			}
			tm.putNode(tm.root, tt.args.key, tt.args.val, 0)
			p := tm.root
			for i := 0; i < len(tt.args.key); i++ {
				c := getIndexLowerCase(tt.args.key[i])
				p = p.Children[c]
				assert.NotEqual(t, nil, p)
			}
			t.Log("val:", p.Val)
			assert.Equal(t, tt.args.val, p.Val)
		})
	}
}

func TestTrieMap_getNode(t *testing.T) {
	type fields struct {
		size int
		root *TrieTreeNode
	}
	type args struct {
		root *TrieTreeNode
		key  string
		val  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TrieTreeNode
	}{
		{
			name: "case 1",
			args: args{
				key: "the",
				val: 4,
			},
			fields: fields{
				root: &TrieTreeNode{
					Children: [26]*TrieTreeNode{},
				},
			},
			want: &TrieTreeNode{
				Val: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TrieMap{
				size: tt.fields.size,
				root: tt.fields.root,
			}
			tm.putNode(tm.root, tt.args.key, tt.args.val, 0)
			got := tm.getNode(tm.root, tt.args.key)
			assert.Equal(t, tt.args.val, got.Val)
		})
	}
}
