package main

import (
	"testing"
)

func Test_checkContainsTarget(t *testing.T) {
	type args struct {
		input  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false if input == 0",
			args: args{
				input:  1,
				target: 7,
			},
			want: false,
		},
		{
			name: "should return false if input has one-digit and not equal to target",
			args: args{
				input:  1,
				target: 7,
			},
			want: false,
		},
		{
			name: "should return false if input has more than one-digit and contains no target",
			args: args{
				input:  123456898654321,
				target: 7,
			},
			want: false,
		},
		{
			name: "should return true if input has one-digit and equal to target",
			args: args{
				input:  7,
				target: 7,
			},
			want: true,
		},
		{
			name: "should return true if input has more than one-digit and contains one target",
			args: args{
				input:  71111111111,
				target: 7,
			},
			want: true,
		},
		{
			name: "should return true if input has more than one-digit and contains multiple targets",
			args: args{
				input:  7777777711111,
				target: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkContainsTarget(tt.args.input, tt.args.target); got != tt.want {
				t.Errorf("checkContainsTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containing7s(t *testing.T) {
	tests := []struct {
		name    string
		number  int
		want    int
		wantErr bool
	}{
		{
			name:    "should return error if input < 0",
			number:  -1,
			want:    0,
			wantErr: true,
		},
		{
			name:    "should return error if input == 0",
			number:  0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "should return 0 if input == 6",
			number:  6,
			want:    0,
			wantErr: false,
		},
		{
			name:    "should return 1 if input == 7",
			number:  6,
			want:    0,
			wantErr: false,
		},
		{
			name:    "should return 2 if input == 20",
			number:  6,
			want:    0,
			wantErr: false,
		},
		{
			name:    "should return 8 if input == 70",
			number:  6,
			want:    0,
			wantErr: false,
		},
		{
			name:    "should return 19 if input == 100",
			number:  6,
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := containing7s(tt.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("containing7s() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("containing7s() = %v, want %v", got, tt.want)
			}
		})
	}
}
