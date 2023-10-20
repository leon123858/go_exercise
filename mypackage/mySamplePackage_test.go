package mySamplePackage

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		num uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Fib(0)",
			args: args{num: 0},
			want: 0,
		},
		{
			name: "Fib(1)",
			args: args{num: 1},
			want: 1,
		},
		{
			name: "Fib(5)",
			args: args{num: 5},
			want: 5,
		},
		{
			name: "Fib(10)",
			args: args{num: 10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.num); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
