package day02

import (
	"reflect"
	"testing"
)

func Test_parsePolicy(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *policyType
		wantErr bool
	}{
		{
			"Sample 1",
			args{str: "1-3 a"},
			&policyType{str: "a", min: 1, max: 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePolicy(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePolicy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
