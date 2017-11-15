package array

import (
	"reflect"
	"testing"
)

func TestIncreaseSubSequence(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Case 1",
			args{[]int{3, 10, 1, 2, 20}},
			[][]int{{1, 2, 20}},
		},
		{
			"Case 2",
			args{[]int{3, 2}},
			[][]int{{3}, {2}},
		},
		{
			"Case 3",
			args{[]int{50, 3, 10, 7, 40, 80}},
			[][]int{{7, 40, 80}},
		},
		{
			"Case 4",
			args{[]int{15, 16, 17, 50, 3, 10, 7, 40, 80}},
			[][]int{{15, 16, 17, 50}},
		},
		{
			"Case fail 1",
			args{[]int{}},
			[][]int{},
		},
		{
			"Case fail 2",
			args{[]int{1}},
			[][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncreaseSubSequence(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncreaseSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
