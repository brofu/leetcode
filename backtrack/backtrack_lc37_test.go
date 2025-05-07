package backtrack

import (
	"reflect"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board  [][]string
		output [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				board: [][]string{
					{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"},
				},
				output: [][]string{{"5", "3", "4", "6", "7", "8", "9", "1", "2"},
					{"6", "7", "2", "1", "9", "5", "3", "4", "8"},
					{"1", "9", "8", "3", "4", "2", "5", "6", "7"},
					{"8", "5", "9", "7", "6", "1", "4", "2", "3"},
					{"4", "2", "6", "8", "5", "3", "7", "9", "1"},
					{"7", "1", "3", "9", "2", "4", "8", "5", "6"},
					{"9", "6", "1", "5", "3", "7", "2", "8", "4"},
					{"2", "8", "7", "4", "1", "9", "6", "3", "5"},
					{"3", "4", "5", "2", "8", "6", "1", "7", "9"}},
			},
		},
	}
	for _, tt := range tests {
		input := make([][]byte, len(tt.args.board))
		for i := range tt.args.board {
			input[i] = make([]byte, len(tt.args.board[i]))
			for j, s := range tt.args.board[i] {
				input[i][j] = []byte(s)[0]
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(input)
			for i := range input {
				for j := range input[0] {
					tt.args.board[i][j] = string(input[i][j])
				}
			}
			t.Log(reflect.DeepEqual(tt.args.board, tt.args.output))
			if !reflect.DeepEqual(tt.args.board, tt.args.output) {
				t.Log("expect:", tt.args.output)
				t.Log("actual:", tt.args.board)
			}
		})
	}
}
