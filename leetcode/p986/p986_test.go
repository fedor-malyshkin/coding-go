package p986

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		firstList  [][]int
		secondList [][]int
	}
	var tests []struct {
		name string
		args args
		want [][]int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.firstList, tt.args.secondList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
