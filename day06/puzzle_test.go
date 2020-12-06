package day06

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

func Test_parseAnswers(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			"Sample1",
			args{lines: readFileAsArray("sample1.txt")},
			[][]string{
				{"abc"},
				{"a", "b", "c"},
				{"ab", "ac"},
				{"a", "a", "a", "a"},
				{"b"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAnswers(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAnswers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAnswers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
