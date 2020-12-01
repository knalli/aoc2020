package day01

import (
	"reflect"
	"testing"
)

func Test_findAll2EntriesMatchingSum(t *testing.T) {
	type args struct {
		numbers  []int
		matchSum int
	}
	tests := []struct {
		name string
		args args
		want []result2Type
	}{
		{
			"Sample 1",
			args{numbers: []int{1721, 979, 366, 299, 675, 1456}, matchSum: 2020},
			[]result2Type{{1721, 299}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAll2EntriesMatchingSum(tt.args.numbers, tt.args.matchSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll2EntriesMatchingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
