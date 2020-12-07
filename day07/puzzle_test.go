package day07

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

func Test_parseRules(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    []ruleType
		wantErr bool
	}{
		{
			"Sample1",
			args{lines: readFileAsArray("sample1.txt")},
			[]ruleType{
				{"light red", map[string]int{"bright white": 1, "muted yellow": 2}},
				{"dark orange", map[string]int{"bright white": 3, "muted yellow": 4}},
				{"bright white", map[string]int{"shiny gold": 1}},
				{"muted yellow", map[string]int{"shiny gold": 2, "faded blue": 9}},
				{"shiny gold", map[string]int{"dark olive": 1, "vibrant plum": 2}},
				{"dark olive", map[string]int{"faded blue": 3, "dotted black": 4}},
				{"vibrant plum", map[string]int{"faded blue": 5, "dotted black": 6}},
				{"faded blue", map[string]int{}},
				{"dotted black", map[string]int{}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRules(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRules() got = %v, want %v", got, tt.want)
			}
		})
	}
}
