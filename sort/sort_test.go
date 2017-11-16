package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[]int{7, 1, 3, 4, 5, 2, 6}},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"case 2",
			args{[]int{}},
			[]int{},
		},
		{
			"case 3",
			args{[]int{7}},
			[]int{7},
		},
		{
			"case 4",
			args{[]int{6, 1, 2, 1, 3, 2, 5, 6, 4}},
			[]int{1, 1, 2, 2, 3, 4, 5, 6, 6},
		},
		{
			"case 5",
			args{[]int{5, 5, 5, 5, 1, 5, 5, 5, 5}},
			[]int{1, 5, 5, 5, 5, 5, 5, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
