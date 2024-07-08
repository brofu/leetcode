package tree

import "testing"

func Test_findSmallestRegion(t *testing.T) {
	type args struct {
		regions [][]string
		region1 string
		region2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			"lc case 1",
			args{
				[][]string{
					{"Earth", "North America", "South America"},
					{"North America", "United States", "Canada"},
					{"United States", "New York", "Boston"},
					{"Canada", "Ontario", "Quebec"},
					{"South America", "Brazil"},
				},
				"Quebec",
				"New York",
			},
			"North America",
		},
		{
			"lc case 2",
			args{
				[][]string{
					{"Earth", "North America", "South America"}, {"North America", "United States", "Canada"}, {"United States", "New York", "Boston"}, {"Canada", "Ontario", "Quebec"}, {"South America", "Brazil"},
				},
				"Canada",
				"South America",
			},
			"Earth",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestRegion(tt.args.regions, tt.args.region1, tt.args.region2); got != tt.want {
				t.Errorf("findSmallestRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
