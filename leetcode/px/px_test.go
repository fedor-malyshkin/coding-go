package px

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case01",
			args: args{
				str: "abcabcbb",
			},
			want: "abc",
		},
		{
			name: "case01",
			args: args{
				str: "abcdabcdefgabcbb",
			},
			want: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.str); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
