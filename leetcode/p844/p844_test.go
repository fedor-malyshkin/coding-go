package p844

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "a#b",
				t: "b",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				s: "bxj##tw",
				t: "bxj###tw",
			},
			want: false,
		}, {
			name: "case 5",
			args: args{
				s: "nzp#o#g",
				t: "b#nzp#o#g",
			},
			want: true,
		}, {
			name: "case 3",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				s: "bxj##tw",
				t: "bxj###tw",
			},
			want: false,
		}, {
			name: "case 5",
			args: args{
				s: "nzp#o#g",
				t: "b#nzp#o#g",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
