package day14

import (
	"reflect"
	"testing"
)

func Test_intToBits(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"11 -> 101",
			args{val: 11},
			[]int{1, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int2Bits(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int2Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitsToInt(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"101 -> 11",
			args{bits: []int{1, 0, 1, 1}},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bits2Int(tt.args.bits); got != tt.want {
				t.Errorf("bits2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
