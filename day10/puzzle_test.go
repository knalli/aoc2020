package day10

import (
	"github.com/knalli/aoc"
	"reflect"
	"testing"
)

func readFileAsArray(path string) []string {
	if lines, err := aoc.ReadFileToArray(path); err != nil {
		panic(err)
	} else {
		return lines
	}
}

func Test_orderAdapters(t *testing.T) {
	type args struct {
		adapters []int
		maxDiff  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   map[int]int
		wantErr bool
	}{
		{
			"Sample1",
			args{adapters: aoc.ParseStringToIntArray(readFileAsArray("sample1.txt")), maxDiff: 3},
			22,
			map[int]int{1: 7, 3: 5},
			false,
		},
		{
			"Sample2",
			args{adapters: aoc.ParseStringToIntArray(readFileAsArray("sample2.txt")), maxDiff: 3},
			52,
			map[int]int{1:22, 3: 10},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := orderAdapters(tt.args.adapters, tt.args.maxDiff)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderAdapters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("orderAdapters() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("orderAdapters() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
