package p162

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 01",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 2,
		}, {
			name: "case 02",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "case 03",
			args: args{
				nums: []int{1, 2, 4, 5, 4, 3, 1, 3, 5, 6, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
