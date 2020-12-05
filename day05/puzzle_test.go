package day05

import (
	"reflect"
	"testing"
)

func Test_decodeBoardingPass(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *passType
		wantErr bool
	}{
		{
			"Sample1",
			args{
				"FBFBBFFRLR",
			},
			&passType{44, 5, 357},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeBoardingPass(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeBoardingPass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeBoardingPass() got = %v, want %v", got, tt.want)
			}
		})
	}
}
