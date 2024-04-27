package str

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},

		{
			name: "lc case 3",
			args: args{
				s1: "abcdxabcde",
				s2: "abcdeabcdx",
			},
			want: true,
		},
		{
			name: "lc local 4",
			args: args{
				s1: "abcdxabcde",
				s2: "abcdefabcdx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInclusionV2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lc case 1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "lc case 2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},

		{
			name: "lc case 3",
			args: args{
				s1: "abcdxabcde",
				s2: "abcdeabcdx",
			},
			want: true,
		},
		{
			name: "lc local 4",
			args: args{
				s1: "abcdxabcde",
				s2: "abcdefabcdx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusionV2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
