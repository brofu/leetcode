package list

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head  *ListNode
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				input: []int{1, 2, 2, 1},
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				input: []int{1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head = ConstructListNodeFromSlice(tt.args.input)
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
